package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	url           *url.URL
	alive         bool
	mutex         sync.RWMutex
	reverse_proxy *httputil.ReverseProxy
}

type ServerPool struct {
	backends []*Backend
	current  uint64
}

// impl nextIndex for ServerPool
// impl getNextBackend for ServerPool

func (s *ServerPool) nextIndex() uint64 {
	// s.current.Add(1)
	return atomic.AddUint64(&s.current, 1)
}

func (b *Backend) isAlive() bool {
	b.mutex.Lock()
	alive := b.isAlive()
	b.mutex.Unlock()
	return alive
}

func (b *Backend) setAlive(alive bool) {
	b.mutex.Lock()
	b.alive = alive
	b.mutex.Unlock()
}

func main() {

	// u, _ := url.Parse("http://localhost:8080")
	// rp := httputil.NewSingleHostReverseProxy()

	// http.HandlerFunc(rp.ServeHTTP)

}
