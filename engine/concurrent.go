/*
 * @Description:调度器复杂版本实现
 * @Author: neozhang
 * @Date: 2022-01-08 21:46:18
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 21:59:42
 */
package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler //Scheduler来源
	WorkerCount      int       //worker数量
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult) //创建输出通道
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler) //创建worker
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out //接收out通道发来的数据
		for _, item := range result.Items {
			go func(i Item) {
				e.ItemChan <- i
			}(item)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParseResult,
	ready ReadyNotifier) {
	go func() {
		for {
			//tell scheduler i am ready
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result //将结果发给out通道
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
