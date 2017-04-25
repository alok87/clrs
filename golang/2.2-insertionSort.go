package main

import "log"
import "math/rand"


func main() {
    arr := rand.Perm(100) 
    log.Print("Insertion Sort", arr)
    for j := 1; j < len(arr); j++ {
        for i := 0; i < j; i ++ {
            if arr[i] < arr[j] {
                arr[i], arr[j] = arr[j], arr[i]   
            }
        }
    }
    log.Print("Sorted array", arr)
}
