package main

import (
	"fmt"
	"math"
)

//implementing Sieve of Eratosthenes
func main() {
	limit := 2000000
	sieve := make([]bool, limit)
	crosslimit := int(math.Floor(math.Sqrt(float64(limit))))

	for i:=3;i <limit;i+=2{
		sieve[i] = true
	}

	for i:=3; i <=crosslimit; i++{
		if sieve[i-1] == false{
			for m:= i*i; m <limit; m=m+i*2{
				sieve[m-1] = true
			}
		}
	}
	sum := 0
	for i:=1;i<limit;i++{
		if sieve[i] == false{
			sum+=i+1
		}
	}
	fmt.Println(sum)
}
