package main

import (
    "testing"
)

// To run test: `go test -cpu=4 -bench=. ./rr_test.go`

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