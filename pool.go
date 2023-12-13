package main

import (
	"errors"
	"time"
)

type Pool[t any] interface {
	Take() (*t, error)
	Put(obj *t)
	Size() int
	Close()
}

type pool[t any] struct {
	ch chan *t
}

func newPool[t any](size int) *pool[t] {
	ch := make(chan *t, size)
	for i := 0; i < size; i++ {
		ch <- new(t)
	}
	return &pool[t]{ch: ch}
}

func (p *pool[t]) Take() (*t, error) {
	for {
		select {
		case obj, open := <-p.ch:
			if !open {
				return nil, errors.New("pool already closed")
			}
			return obj, nil
		case <-time.After(50 * time.Millisecond):
			// could also create a new obj
			continue
		}
	}
}

func (p *pool[t]) Put(obj *t) {
	// put back to the free pool
	p.ch <- obj
}

func (p *pool[t]) Size() int {
	// get the pool size
	return len(p.ch)
}

func (p *pool[t]) Close() {
	select {
	case <-p.ch:
	default:
		close(p.ch)
	}
	for item := range p.ch {
		_ = item
	}
}
