package main

import (
	"fmt"
	"math"
)

func main(){
	var num int = 600851475143
	sqrtOfNum := int(math.Sqrt(float64(num)))
	if !isPrime(num){
		temp := num
		highestPrimeFactor := 0
		for i:=2; i < sqrtOfNum ;{
			if(isPrime(i)){
				if temp%i==0{
					temp /= i
					if i > highestPrimeFactor{
						highestPrimeFactor = i
					}
				}else{
					i++
				}
			}
		}

	}else{
		fmt.Println(num)
	}
}


func isPrime(num int) bool{
	foundDivisor := false
	if num == 2{
		return true
	}else{
		for i:=2 ; i<= int(math.Sqrt(float64(num)));i++{
			if num%i == 0{
				foundDivisor = true
				break
			}
		}
		if foundDivisor == true{
			return false
		}
		return true
	}
}