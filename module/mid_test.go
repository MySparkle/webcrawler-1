package module

import (
	"net"
	"reflect"
	"testing"
)

func TestGenMID(t *testing.T) {
	type args struct {
		mtype Type
		sn    uint64
		maddr net.Addr
	}
	tests := []struct {
		name    string
		args    args
		want    MID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenMID(tt.args.mtype, tt.args.sn, tt.args.maddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenMID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenMID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitMID(t *testing.T) {
	type args struct {
		mid MID
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitMID(tt.args.mid)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitMID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitMID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_legalSN(t *testing.T) {
	type args struct {
		snStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := legalSN(tt.args.snStr); got != tt.want {
				t.Errorf("legalSN() = %v, want %v", got, tt.want)
			}
		})
	}
}
