package time

import (
	"github.com/zapotek6/go-util/util"
	"runtime"
	"time"
)

type ExecutionHandler = func()

type ScheduledExecution struct {
	name       string
	run        bool
	killChan   chan string
	notifyChan chan string
	event      string
	startTime  time.Time
	duration   time.Duration
	handler    ExecutionHandler
	gohlp      *util.GoroutineHelper
}

func newTimer(name string, duration time.Duration, notifyChan chan string, event string, timerhandler ExecutionHandler) *ScheduledExecution {
	timer := ScheduledExecution{
		name:       name,
		startTime:  time.Now(),
		duration:   duration,
		killChan:   make(chan string),
		notifyChan: notifyChan,
		event:      event,
		handler:    timerhandler,
		gohlp:      util.NewGoroutineHelper(),
	}

	return &timer
}

func NewTimerWFunc(name string, duration time.Duration, timerhandler ExecutionHandler) *ScheduledExecution {
	return newTimer(name, duration, nil, "", timerhandler)
}

func NewTimerWChan(name string, duration time.Duration, notifyChan chan string, event string) *ScheduledExecution {
	return newTimer(name, duration, notifyChan, event, nil)
}

func (t *ScheduledExecution) Run() *ScheduledExecution {
	go func() {
		t.run = true
		t.gohlp.Add()
		timer := time.NewTimer(t.duration)
		for {
			select {
			case <-timer.C:
				//fmt.Println("time expired send notification", t.name, t.duration)
				if nil != t.notifyChan {
					t.notifyChan <- t.event
				} else if nil != t.handler {
					t.handler()
				}
				runtime.Goexit()
			case <-t.gohlp.GetQuitChan():
				if nil != timer {
					timer.Stop()
				}
				t.gohlp.Done()
				//fmt.Println("kill chan", t.name)
				runtime.Goexit()
			}
		}
	}()

	return t
}

func (t *ScheduledExecution) Kill() {
	t.gohlp.Quit()
}
