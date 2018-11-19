package stub

import "webcrawler/module"

type ModuleInternal interface {
	module.Module
	//调用计数器加1
	IncrCalledCount()
	//接受计数器加1
	IncrAcceptedCount()
	//成功完成计数器加1
	IncrCompletedCount()
	//实时处理数加1
	IncrHandlingNumber()
	//实时处理数减1
	DecrHandlingNumber()
	//清空计数
	Clear()
}
