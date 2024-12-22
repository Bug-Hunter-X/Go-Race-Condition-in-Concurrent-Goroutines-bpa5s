```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        fmt.Printf("Goroutine %d is running\n", i)
                }(i) // Pass i to the closure
        }

        wg.Wait()
}
```