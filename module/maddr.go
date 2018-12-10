package module

import (
	"fmt"
	"net"
	"strconv"
	"webcrawler/errors"
)

type mAddr struct {
	network string
	address string
}

func (maddr *mAddr) Network() string {
	return maddr.network
}

func (maddr *mAddr) String() string {
	return maddr.address
}

// NewAddr 会根据参数创建并返回一个网络地址值。
// 如果参数不合法，那么会返回非nil的错误值。
func NewAddr(network string, ip string, port uint64) (net.Addr, error) {
	if network != "http" && network != "https" {
		errMsg := fmt.Sprintf("illegal network for module address: %s", network)
		return nil, errors.NewIllegalParameterError(errMsg)
	}
	if parsedIP := net.ParseIP(ip); parsedIP == nil {
		errMsg := fmt.Sprintf("illegal IP for module address: %s", ip)
		return nil, errors.NewIllegalParameterError(errMsg)
	}
	return &mAddr{
		network: network,
		address: ip + ":" + strconv.Itoa(int(port)),
	}, nil
}
