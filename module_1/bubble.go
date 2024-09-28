package main

import (
    "fmt"
)

const INPUT_NUM = 10

func main () {
    inputList := make([]int, INPUT_NUM)

    for i := 0; i < INPUT_NUM; i++ {
        var usrInput int
        fmt.Printf("Enter a number: \t")
        fmt.Scanf("%d", &usrInput)

        inputList[i] = usrInput
    }

    fmt.Println(inputList)
    BubbleSort(&inputList)
    fmt.Println(inputList)

    return
}

func BubbleSort (numList *[]int) {
    for {
        swapped := false

        for i := 0; i < INPUT_NUM - 1; i++ {
            if (*numList)[i] > (*numList)[i + 1] {
                Swap(numList, i)
                swapped = true
            }
        }

        if !swapped {
            break
        }
    }

    return
}

func Swap (numList *[]int, indexSwap int) {
    var tmp int = (*numList)[indexSwap + 1]

    (*numList)[indexSwap + 1] = (*numList)[indexSwap]
    (*numList)[indexSwap] = tmp

    return
}
