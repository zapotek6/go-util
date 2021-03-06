package util

import (
	"sync"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
	Quit() bool
	Add()
	AddMany(delta int)
	Wait()
	Remove()
	IsQuitting() bool
	GetQuitChan() chan struct{}
}

type GoroutineHelper struct {
	quitting bool
	quit     chan struct{}
	wg       sync.WaitGroup
}

func NewGoroutineHelper() *GoroutineHelper {
	g := new(GoroutineHelper)
	g.quit = make(chan struct{})

	return g
}

func (g *GoroutineHelper) Deadline() (deadline time.Time, ok bool) {
	return
}

func (g *GoroutineHelper) Done() <-chan struct{} {
	return g.quit
}

func (g *GoroutineHelper) Err() error {
	return nil
}

func (g *GoroutineHelper) Value(key interface{}) interface{} {
	return nil
}

func (g *GoroutineHelper) GetQuitChan() chan struct{} {
	return g.quit
}

func (g *GoroutineHelper) Quit() bool {
	if nil != g.quit && !g.quitting {
		close(g.quit)
		g.quitting = true
	}
	return g.quitting
}

func (g *GoroutineHelper) Add() {
	g.AddMany(1)
}

func (g *GoroutineHelper) AddMany(delta int) {
	g.wg.Add(delta)
}

func (g *GoroutineHelper) Wait() {
	g.wg.Wait()
}

func (g *GoroutineHelper) Remove() {
	g.wg.Done()
}

func (g *GoroutineHelper) IsQuitting() bool {
	return g.quitting
}
