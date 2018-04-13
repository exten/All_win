/*
 *  flower9
 *  14/12/2017
 *
 */
package common

import (
	"sync"
	"time"
	"log"
)

func init() {
	log.Println("init timer sync...")
}

type BatchAsync struct {
	syncChann chan int
	count     int
}

var (
	lock  = &sync.Mutex{}
	Async = &BatchAsync{make(chan int), 0}
)


//增加同步标记，用户必须自己将程序阻塞
func (sync *BatchAsync) AddSync() chan int {
	lock.Lock()
	defer lock.Unlock()

	sync.count ++;
	return sync.syncChann
}

func (sync *BatchAsync) restSync() (int, chan int) {
	lock.Lock()
	defer lock.Unlock()

	count := sync.count
	sync.count = 0;
	return count, sync.syncChann
}

func TimerBackRunOpenLotteryTicket(duration time.Duration) {
	time.AfterFunc(duration, func() {
		count, chl := Async.restSync()

		log.Println("定时发放开奖号码信息 -- 开始")
		for i := 0; i < count; i++ {
			chl <- i
		}
		log.Println("定时发放开奖号码信息 -- 完成")
		TimerBackRunOpenLotteryTicket(duration)
	})
}
