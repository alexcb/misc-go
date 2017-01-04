package main

import "fmt"

func calcE(n int) float64 {
	e := float64(1)
	for i := 1; i < n; i++ {
		nFactorial := float64(1)
		for j := i; j > 0; j-- {
			nFactorial = nFactorial * float64(j)
		}
		e += 1.0 / float64(nFactorial)
	}
	return e
}

func main() {
	fmt.Printf("%f\n", calcE(1000))
}
