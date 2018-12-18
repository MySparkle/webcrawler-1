package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	lib "webcrawler/examples/finder/internal"
	"webcrawler/examples/finder/monitor"
	"webcrawler/helper/log"
	sched "webcrawler/scheduler"
)

var (
	firstURL string
	domains  string
	depth    uint
	dirPath  string
)

var logger = log.DLogger()

func init() {
	flag.StringVar(&firstURL, "first", "http://zhihu.sogou.com/zhihu?query=golang+logo", "The first URL which you want to access.")
	flag.StringVar(&domains, "domains", "zhihu.com",
		"The primary domains which you accepted. "+"Please using comma-separated multiple domains.")
	flag.UintVar(&depth, "depth", 3, "The depth for crawling.")
	flag.StringVar(&dirPath, "dir", "./pictures", "The path which you want to save the image files.")
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tfinder [flags] \n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	fmt.Printf("firstURL -first=%s domains -domains=%s depth -depth=%d dirPath -dir=%s", firstURL, domains, depth, dirPath)
	//创建调度器
	scheduler := sched.NewScheduler()
	//准备调度器的初始参数
	domainParts := strings.Split(domains, ",")
	acceptedDomains := []string{}
	for _, domain := range domainParts {
		domain = strings.TrimSpace(domain)
		if domain != "" {
			acceptedDomains = append(acceptedDomains, domain)
		}
	}
	requestArgs := sched.RequestArgs{
		AcceptedDomains: acceptedDomains,
		// 需要被爬的最大深度
		MaxDepth: uint32(depth),
	}
	fmt.Printf("\n requestArgs %v, scheduler %v", requestArgs, scheduler)
	dataArgs := sched.DataArgs{
		// 请求缓冲器的容量
		ReqBufferCap: 50,
		// 最大数量
		ReqMaxBufferNumber: 1000,
		// 响应缓冲器的容量
		RespBufferCap: 50,
		// 响应缓冲器的最大数量
		RespMaxBufferNumber: 10,
		// 条目缓冲器的容量
		ItemBufferCap: 50,
		// 条目缓冲器最大数
		ItemMaxBufferNumber: 100,
		//错误缓冲器的容量
		ErrorBufferCap: 50,
		//错误缓冲器的最大数量
		ErrorMaxBufferNumber: 1,
	}
	downloaders, err := lib.GetDownloaders(1)
	if err != nil {
		logger.Fatalf("An error occurs when creating downloaders: %s", err)
	}
	analyzers, err := lib.GetAnalyzers(1)
	if err != nil {
		logger.Fatalf("An error occurs when creating analyzers:%s", err)
	}
	pipelines, err := lib.GetPipelines(1, dirPath)
	if err != nil {
		logger.Fatalf("An error occurs when creating pipelines:%s", err)
	}
	moduleArgs := sched.ModuleArgs{
		Downloaders: downloaders,
		Analyzers:   analyzers,
		Pipelines:   pipelines,
	}
	// 初始化调度器
	err = scheduler.Init(
		requestArgs,
		dataArgs,
		moduleArgs)
	if err != nil {
		logger.Fatalf("An error occurs when initializing scheduler:%s", err)
	}
	//准备监控参数
	checkInterval := time.Second
	summarizeInterval := 100 * time.Microsecond
	maxIdleCount := uint(5)
	checkCountChan := monitor.Monitor(
		scheduler,
		checkInterval,
		summarizeInterval,
		maxIdleCount,
		true,
		lib.Record)
	//准备调度器的参数启动
	firstHTTPReq, err := http.NewRequest("GET", firstURL, nil)
	if err != nil {
		logger.Fatalln(err)
		return
	}
	// 开启调度器
	err = scheduler.Start(firstHTTPReq)
	if err != nil {
		logger.Fatalf("An error occurs when starting scheduler: %s", err)
	}
	//等待监控结束
	<-checkCountChan
}
