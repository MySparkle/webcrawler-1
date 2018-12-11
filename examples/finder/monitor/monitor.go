package mointor

import (
	"webcrawler/helper/log"
	sched "webcrawler/scheduler"
)

var logger = log.DLogger()

// 监控结果摘要
type summary struct {
	// Goroutine 数量
	NumGoroutine int `json:"goroutine_number"`
	//调度器的摘要信息
	SchedSummary sched.SummaryStruct `json:"sched_summary"`
	//开始监控至今的时间
	EscapedTime string `json:"escaped_time"`
}

// 已达到最大空闲计数的消息模板
var msgReachMaxIdleCount = "The scheduler has been idle for a period of time" +
	" (about %s)." + " Consider to stop it now."

// 代表停止调度器的消息模板
var msgStopScheduler = "Stop scheduler ...%s."

//日志记录函数的类型
// 参数level代表日志级别 0-普通，1-警告, 2-错误
type Record func(level uint8, content string)

// Monitor 用于监控调度器。
// 参数scheduler代表作为监控目标的调度器。
// 参数checkInterval代表检查间隔时间，单位：纳秒。
// 参数summarizeInterval代表摘要获取间隔时间，单位：纳秒。
// 参数maxIdleCount代表最大空闲计数。
// 参数autoStop被用来指示该方法是否在调度器空闲足够长的时间之后自行停止调度器。
// 参数record代表日志记录函数。
// 当监控结束之后，该方法会向作为唯一结果值的通道发送一个代表了空闲状态检查次数的数值。
