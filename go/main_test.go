package main

import (
	"testing"
	"time"
)

func BenchmarkMain(b *testing.B) {
	startTime := time.Now()
	main()
	b.Log("time", time.Now().Sub(startTime).Seconds())
}

/**
	синхрон
2м 7с
--- BENCH: BenchmarkMain-4
    main_test.go:9: start Friday, 29-Mar-19 13:49:59 MSK
    main_test.go:11: stop Friday, 29-Mar-19 13:52:06 MSK
*/

/**
асинх
BenchmarkMain-4   	       1	136822769111 ns/op
--- BENCH: BenchmarkMain-4
    main_test.go:9: start Friday, 29-Mar-19 14:49:14 MSK
    main_test.go:11: stop Friday, 29-Mar-19 14:51:31 MSK
PASS
*/
