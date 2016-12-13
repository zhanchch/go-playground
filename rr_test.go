package playgound

import (
    
    "sync"
    "testing"
)

// To run test: `go test -cpu=4 -bench=. ./rr_test.go`

var mu = sync.Mutex{}
var urls = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
var healthy = []bool{true, false, true, true, false, false, true, true}
var which int

var getRR = make(chan chan string)

func start() {
    for {
        select {
            case rr := <- getRR:
                rr <- RRNoLock()
        }
    }
}

func RRChan() string{
    rr := make(chan string)
    
    getRR <- rr
    return <- rr
}

func RRLock() string {
    mu.Lock()
    defer mu.Unlock()
    length := len(urls)
    for i := 0; i < length; i++ {
        which = (which + 1) % length
        if healthy[which]{
            return urls[which]
        }
    }
    return ""
}

func RRNoLock() string {
    length := len(urls)
    for i := 0; i < length; i++ {
        which = (which + 1) % length
        if healthy[which]{
            return urls[which]
        }
    }
    return ""
}

func BenchmarkRRLock(b *testing.B){
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            RRLock()
        }
    })
}

func BenchmarkRRNoLock(b *testing.B){
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            RRNoLock()
        }
    })
}

func BenchmarkRRChan(b *testing.B){
    go start()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            RRChan()
        }
    })
}