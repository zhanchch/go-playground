package main

import (
    "sync"
    "sync/atomic"
)

var mu = sync.Mutex{}
var urls = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
var healthy = []bool{true, false, true, true, false, false, true, true}
var which int = 0
var atomicWhich uint32 = 1<<32 - 5

var getRR = make(chan chan string)

func start() {
    for {
        select {
            case rr := <- getRR:
                rr <- RRNoLock()
        }
    }
}

type rrfunction func() string

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

func RRAtomic() string{
    length := len(urls)
    for i := 0; i < length; i++ {
        num := atomic.AddUint32(&atomicWhich, 1)
        //fmt.Print("Atomic")
        //fmt.Println(num)
        
        idx := int(num) % length
        if healthy[idx]{
            return urls[idx]
        }
    }
    return ""
}

func main() {
}