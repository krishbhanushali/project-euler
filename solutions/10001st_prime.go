package main

import (
	"fmt"
	"math"
)

func main (){
	count := 5
	pointer := 10
	for count <= 10001{
		if pointer%2!=0 {
			upperLimit := int(math.Sqrt(float64(pointer)))
			var nonPrimeFound bool = false
			for i := 2; i <= upperLimit; i++ {
				if pointer%i == 0 {
					nonPrimeFound = true
					break
				}
			}
			if !nonPrimeFound{
				count ++
			}
		}
		pointer ++;
	}
	fmt.Println(pointer-1)
}
