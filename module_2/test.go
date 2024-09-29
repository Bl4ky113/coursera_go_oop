package main

import "fmt"

func main() {

	var a, s, v, t float64
	fmt.Printf("Enter acceleration : ")
	fmt.Scanf("%f ", &a)
	fmt.Printf("Enter initial velocity : ")
	fmt.Scanf("%f ", &v)
	fmt.Printf("Enter initial displacement : ")
	fmt.Scanf("%f ", &s)
	fmt.Printf("Enter time : ")
	fmt.Scanf("%f ", &t)

	displacement := GenDisplaceFn(a, v, s)
	fmt.Println("Final Displacement : ", displacement(t))
}

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v*t + s
	}
}
