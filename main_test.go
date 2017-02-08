package main

import (
    "testing"
)

// To run test: `go test -cpu=1,4,8 -bench=.`
/*
Some benchmark result
MM-MAC-3365:go-playground czhan$ go test -cpu=1,4,8 -bench=.
testing: warning: no tests to run
BenchmarkRRNoLock       100000000           21.4 ns/op
BenchmarkRRNoLock-4     100000000           15.3 ns/op
BenchmarkRRNoLock-8     100000000           12.7 ns/op
BenchmarkRRAtomic       50000000            40.0 ns/op
BenchmarkRRAtomic-4     30000000            46.1 ns/op
BenchmarkRRAtomic-8     50000000            35.8 ns/op
BenchmarkRRLock         20000000           109 ns/op
BenchmarkRRLock-4       10000000           201 ns/op
BenchmarkRRLock-8        5000000           266 ns/op
BenchmarkRRChan          3000000           578 ns/op
BenchmarkRRChan-4        3000000           407 ns/op
BenchmarkRRChan-8        3000000           582 ns/op
PASS
ok      github.com/zhanchch/go-playground   22.667s
*/

func BenchmarkRRNoLock(b *testing.B){
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            RRNoLock()
        }
    })
}

func BenchmarkRRAtomic(b *testing.B){
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            RRAtomic()
        }
    })
}

func BenchmarkRRLock(b *testing.B){
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            RRLock()
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