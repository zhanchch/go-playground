package main

import (
    "sync"
    "sync/atomic"
    "time"
    "fmt"
    "math/rand"
)

var rwmtx = sync.RWMutex{}
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

func getHealth(i int) bool{
    rwmtx.Lock()
    defer rwmtx.Unlock()
    return healthy[which]
}

func RRNoLock() string {
    length := len(urls)
    for i := 0; i < length; i++ {
        which = (which + 1) % length
        //defer rwmtx.Unlock()
        if getHealth(which){
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func init() {
    rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func main() {
    abc := struct{
        A string
        B string
    }{
        "a",
        "b",
    }
    
    go func(){
        for{
            abc.A = RandStringRunes(1)
            abc.B = RandStringRunes(1)
        }    
    }()
    go func(){
        for{
            fmt.Println(abc.A, abc.B)
        }
    }()
    time.Sleep(time.Second)
}