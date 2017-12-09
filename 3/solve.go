package main

import "math"

func getVal(value int) int {
    num := 2
    curr := 1
    for num < value {
        num = num + curr*8
        curr = curr + 1
    }

    curr = curr -1
    num = num - curr*8

    left := value - num
    side := (math.Mod(float64(left+1),float64(curr*2)))
    return int(math.Abs(float64(curr)-side))+curr
}

func main() {
    print("The distance between the 1 and our input is: ")
    println(getVal(265149))
}
