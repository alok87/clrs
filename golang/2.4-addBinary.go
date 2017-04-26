package main

import (
    "log"
)

func pow(a int, b int) int {
    result := 1
    for i:=1; i<=b; i++ {
        result *= a
    }
    return result
}

func verifyResult(a []int, b[]int, c[] int) bool {
    log.Print(binaryToInt(a), binaryToInt(b), binaryToInt(c))
    return binaryToInt(a) + binaryToInt(b) == binaryToInt(c)
}

func binaryToInt(num []int) int {
    intResult := 0
    j := 0
    for i:=len(num)-1; i>=0; i-- {
        intResult += num[i] * pow(2, j)
        j += 1
    }
    return intResult
}

func add(a []int, b []int) []int {
    n := len(a)
    var c = make([]int, n + 1)
    carry := 0
    for i:=n-1; i>=0; i-- {
        j := i + 1
        c[j] = a[i] + b[i] + carry
        if c[j] > 1 {
            c[j] -= 2
            carry = 1
        } else {
            carry = 0
        }
    }
    c[0] = carry
    return c
}

func main() {
    // A[n]   = [1, 0, 0, 1, 0, 1, 0]  ~ 90
    // B[n]   = [1, 0, 0, 0, 0, 1, 0]  ~ 10
    // C[n+1] = A + B                  ~ 100
    //Binary numbers addition - https://www.youtube.com/watch?v=jB_sRh5yoZk

    a := []int{1, 0, 0, 1, 0, 1, 0}
    b := []int{1, 0, 0, 0, 0, 1, 0}
    log.Print("Adding: ", a, b)
    result := add(a, b)
    log.Print("Result: ", result)
    log.Print("Test Result: ", verifyResult(a, b, result))
}
