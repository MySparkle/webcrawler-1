package scheduler

import "webcrawler/module"

type Args interface {
	Check() error
}

type RequestArgs struct {
	// 代表可以接受的URL的主域名列表
	AcceptedDomains []string `json:"accepted_primary_domains"`
	// 需要被爬的最大深度
	MaxDepth uint32 `json:max_depth`
}

func (args *RequestArgs) Check() error {
	if args.AcceptedDomains == nil {
		return genError("nil accepted primary domain list")
	}
	return nil
}

func (args *RequestArgs) Same(another *RequestArgs) bool {
	if another == nil {
		return false
	}
	if another.MaxDepth != args.MaxDepth {
		return false
	}
	anotherDomains := another.AcceptedDomains
	anotherDomainsLen := len(anotherDomains)
	if anotherDomainsLen != len(args.AcceptedDomains) {
		return false
	}
	if anotherDomainsLen > 0 {
		for i, domain := range anotherDomains {
			if domain != args.AcceptedDomains[i] {
				return false
			}
		}
	}
	return true
}

type DataArgs struct {
	// 请求缓冲器的容量
	ReqBufferCap uint32 `json:"req_buffer_cap"`
	// 最大数量
	ReqMaxBufferNumber uint32 `json:"req_max_buffer_number"`
	// 响应缓冲器的容量
	RespBufferCap uint32 `json:"resp_buffer_cap"`
	// 响应缓冲器的最大数量
	RespMaxBufferNumber uint32 `json:"resp_max_buffer_number"`
	// 条目缓冲器的容量
	ItemBufferCap uint32 `json:"item_max_buffer_cap"`
	// 条目缓冲器最大数
	ItemMaxBufferNumber uint32 `json:"item_buffer_cap"`
	//错误缓冲器的容量
	ErrorBufferCap uint32 `json:"error_buffer_cap"`
	//错误缓冲器的最大数量
	ErrorMaxBufferNumber uint32 `json:"error_max_buffer_number"`
}

func (args *DataArgs) Check() error {
	if args.ReqBufferCap == 0 {
		return genError("zero request buffer capacity")
	}
	if args.ReqMaxBufferNumber == 0 {
		return genError("zero max request buffer number")
	}
	if args.RespBufferCap == 0 {
		return genError("zero response buffer capacity")
	}
	if args.RespMaxBufferNumber == 0 {
		return genError("zero max response buffer number")
	}
	if args.ItemBufferCap == 0 {
		return genError("zero item buffer capacity")
	}
	if args.ItemMaxBufferNumber == 0 {
		return genError("zero  max item buffer number")
	}
	if args.ErrorBufferCap == 0 {
		return genError("zero error buffer capacity")
	}
	if args.ErrorMaxBufferNumber == 0 {
		return genError("zero max error buffer number")
	}
	return nil
}

type ModuleArgsSummary struct {
	DownloaderListSize int `json:"downloader_list_size"`
	AnalyzerListSize   int `json:"analyzer_list_size"`
	PipelineListSize   int `json:"pipeline_list_size"`
}

type ModuleArgs struct {
	Downloaders []module.Downloader
	Analyzers   []module.Analyzer
	Pipelines   []module.Pipeline
}

func (args *ModuleArgs) Check() error {
	if len(args.Downloaders) == 0 {
		return genError("empty downloader list")
	}
	if len(args.Analyzers) == 0 {
		return genError("empty analyzer list")
	}
	if len(args.Pipelines) == 0 {
		return genError("empty pipeline list")
	}
	return nil
}

func (args *ModuleArgs) Summary() ModuleArgsSummary {
	return ModuleArgsSummary{
		DownloaderListSize: len(args.Downloaders),
		AnalyzerListSize:   len(args.Analyzers),
		PipelineListSize:   len(args.Pipelines),
	}
}
