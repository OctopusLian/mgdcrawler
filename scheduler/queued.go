/*
 * @Description:Scheduler实现版本3
 * @Author: neozhang
 * @Date: 2022-01-08 21:46:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 21:51:21
 */
package scheduler

import "mgdcrawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	//生成channel
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		//make两个队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select { //监听多个通道
			case r := <-s.requestChan:
				//收到一个request就排队
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				//收到一个work就排队
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
