// results-------------------------------------
// @file      : manage.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2024/9/17 20:59
// -------------------------------------------

package results

import (
	"time"
)

const (
	batchSize     = 50
	flushInterval = 30 * time.Second
)

type ResultQueue struct {
	Queue   chan interface{}
	CloseCh chan struct{}
}

var ResultQueues = make(map[string]*ResultQueue)

func InitializeResultQueue() {
	// 模块列表
	modules := []string{
		"SubdomainScan", "SubdomainSecurity",
		"AssetMapping", "PortScan", "AssetResultHandl", "URLScan",
		"URLSecurity", "WebCrawler", "VulnerabilityScan",
	}
	// 初始化模块队列和 Goroutine
	for _, module := range modules {
		ResultQueues[module] = &ResultQueue{
			Queue: make(chan interface{}, batchSize),
		}
		go processQueue(module, ResultQueues[module])
	}

	// 初始化去重模块
	InitializeDuplicate()
	// 初始化结果处理模块
	InitializeHandler()
}

func processQueue(module string, mq *ResultQueue) {
	ticker := time.NewTicker(flushInterval)
	defer ticker.Stop()

	var buffer []interface{}

	for {
		select {
		case batch := <-mq.Queue:
			buffer = append(buffer, batch)
			if len(buffer) >= batchSize {
				flushBuffer(module, &buffer)
			}
		case <-ticker.C:
			if len(buffer) > 0 {
				flushBuffer(module, &buffer)
			}
		case <-mq.CloseCh:
			// 处理关闭信号
			if len(buffer) > 0 {
				flushBuffer(module, &buffer)
			}
			return
		}
	}
}

func flushBuffer(module string, buffer *[]interface{}) {
	if len(*buffer) == 0 {
		return
	}
	switch module {
	case "SubdomainScan":
		SubdoaminResultHandler(buffer)
	}

	*buffer = nil
}

func Close() {
	for _, mq := range ResultQueues {
		close(mq.Queue)   // 关闭队列
		close(mq.CloseCh) // 发送关闭信号
	}
}