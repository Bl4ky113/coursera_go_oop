package main

import (
    "fmt"
)

func main () {
    var acceleration, velocityIni, displacementIni, time float64
    
    fmt.Printf("Enter the acceleration:\t")
    fmt.Scanf("%f", &acceleration)
    fmt.Printf("Enter the initial velocity:\t")
    fmt.Scanf("%f", &velocityIni)
    fmt.Printf("Enter the initial displacement:\t")
    fmt.Scanf("%f", &displacementIni)
    fmt.Printf("Enter the time:\t")
    fmt.Scanf("%f", &time)

    displacementFunction := GenDisplaceFn(
        acceleration,
        velocityIni,
        displacementIni,
    )

    fmt.Printf("The Displacement after %.2f secconds is:\t%.3f", time, displacementFunction(time))

    return
}

func GenDisplaceFn (acc, velIni, dispIni float64) (func (float64) (float64)) {
    return func (time float64) float64 {
        return ((acc/2) * (time * time)) + (velIni * time) + (dispIni)
    }
}
