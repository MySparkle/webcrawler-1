package scheduler

import (
	"context"
	"net/http"
	"sync"
	"webcrawler/helper/log"
	"webcrawler/module"
	"webcrawler/toolkit/buffer"
)

var logger = log.DLogger()

type Scheduler interface {
	// 初始化调度器
	Init(requestArgs RequestArgs,
		dataArgs DataArgs,
		moduleArgs ModuleArgs) (err error)
	//启动调度器并执行爬取流程
	Start(firstHHTPReq *http.Request) (err error)
	//停止调度器的运行
	Stop() (err error)
	//获取调度器的状态
	Status() Status
	//获得错误通道
	ErrorChan() <-chan error
	//判断所有处理模块是否都处于空闲状态
	Idle() bool
	//获取摘要实例
	Summary() SchedSummary
}

func NewScheduler() Scheduler {
	return &myScheduler{}
}

//调度器实现类型
type myScheduler struct {
	//爬取的最大深度 首次为0
	maxDepth uint32
	//可接受的URL主域名的字典
	acceptedDomainMap cmap.ConcurrentMap
	//组件注册器
	registrar module.Registrar
	//请求的缓冲池
	reqBufferPool buffer.Pool
	//响应的缓冲池
	respBufferPool buffer.Pool
	//条目缓冲池
	itemBufferPool buffer.Pool
	//错误缓冲池
	errorBufferPool buffer.Pool
	//已处理的URL字典
	urlMap cmap.ConcurrentMap
	//上下文 用于感知调度器的停止
	ctx context.Context
	//取消函数 用于停止调度器
	cancelFunc context.CancelFunc
	//代表状态
	status Status
	//专用于状态的读写锁
	statusLock sync.RWMutex
	//代表摘要信息
	summary SchedSummary
}

func (sched *myScheduler) Init(
	requestArgs RequestArgs,
	dataArgs DataArgs,
	moduleArgs ModuleArgs) (err error) {
	// 检查状态
	logger.Info("Check status for inititalization...")
	var oldStatus Status
	oldStatus, err = sched.checkAndSetStatus(SCHED_STATUS_INITIALIZING)
	if err != nil {
		return
	}
	defer func() {
		sched.statusLock.Lock()
		if err != nil {
			sched.status = oldStatus
		} else {
			sched.status = SCHED_STATUS_INITIALIZED
		}
		sched.statusLock.Unlock()
	}()
	//检查参数
	logger.Info("Check request arguments...")
	if err = requestArgs.Check(); err != nil {
		return err
	}
	logger.Info("Check data arguments...")
	if err = dataArgs.Check(); err != nil {
		return err
	}
	logger.Info("Data arguments are valid.")
	logger.Info("Check module arguments...")
	if err = moduleArgs.Check(); err != nil {
		return err
	}
	logger.Info("Module arguments are valid.")
	//初始化内部字段
	logger.Info("Initialize scheduler's fields...")
	if sched.registrar == nil {
		sched.registrar = module.NewRegistrar()
	} else {
		sched.registrar.Clear()
	}
	sched.maxDepth = requestArgs.MaxDepth
	logger.Infof("-- Max depth: %d", sched.maxDepth)
	sched.acceptedDomainMap, _ = cmap.NewCocurrentMap(1, nil)
	for _, domain := range requestArgs.AcceptedDomains {
		sched.acceptedDomainMap.Put(domain, struct{}{})
	}
	logger.Infof("-- Accepted primary domains: %v", requestArgs.AcceptedDomains)
	sched.urlMap, _ = cmap.NewConcurrentMap(16, nil)
	logger.Infof("-- URL map: length: %d, concurrency: %d",
		sched.urlMap.Len(), sched.urlMap.Concurrency())
	sched.initBufferPool(dataArgs)
	sched.resetContext()
	sched.summary = newSchedSummary(requestArgs, dataArgs, moduleArgs, sched)
	//注册组件
	logger.Info("Register modules...")
	if err = sched.registrarModules(moduleArgs); err != nil {
		return
	}
}
