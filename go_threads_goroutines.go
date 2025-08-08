package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func main() {
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
    fmt.Println("Initial Goroutines:", runtime.NumGoroutine())

    var wg sync.WaitGroup
    total := 100000
    wg.Add(total)

    for i := 0; i < total; i++ {
        go func(i int) {
            defer wg.Done()
            time.Sleep(10 * time.Millisecond)
            if i == 0 {
                fmt.Println("Just one goroutine says hi!")
            }
        }(i)
    }

    // Let the scheduler catch up
    time.Sleep(100 * time.Millisecond)
    fmt.Println("After spawn Goroutines:", runtime.NumGoroutine())

    wg.Wait()
    fmt.Println("Done.")
}
