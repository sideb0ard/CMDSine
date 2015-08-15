package main

func fib(ticker chan int) {

	for {
		i := <-ticker
		sumNum = fib_tail(0, 1, i) % int(bpm)
	}
}

func fib_tail(a, b, n int) int {
	if n > 0 {
		return fib_tail(b, a+b, n-1)
	} else {
		return a
	}
}
