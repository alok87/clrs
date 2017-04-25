package main

import "math/rand"
import "time"
import "log"

func random(min, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return rand.Intn(max-min) + min
}


func main() {
    magic := 10
    n := rand.Perm(magic)
    log.Print("Arr ==> ", n)
    searchingFor := random(1, 30)
    var found bool
    log.Print("Searching for ==> ", searchingFor)
    for i:=0; i<len(n); i++ {
        if searchingFor == n[i] {
            log.Print("Found at ", i)
            found = true
            break
        }
    }
    if ! found {
        log.Print("Not Found")
    }
}
    
