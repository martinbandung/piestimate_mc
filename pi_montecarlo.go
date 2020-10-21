package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
)
//menghitung nilai Pi dengan monte carlo pada kuadran 1 (1/4 lingkaran)
var x=0.0
var y=0.0
//ganti n dengan lebih besar supaya mungkin lebih akurat

var ndalam=0.0
var nluar=0.0

func main() {
    fmt.Println("Simple Pi Estimation Using Monte Carlo")
    var n int
    fmt.Print("N Points\t:  ")
    fmt.Scanf("%d", &n)
    rand.Seed(time.Now().UTC().UnixNano())
    for i := 0; i < int(n); i++ {
        x=rand.Float64()
        //fmt.Println(rand.Float64())
        y=rand.Float64()
        if y <= math.Sqrt(1-x*x){
            ndalam=ndalam+1
        }
     }
    
    fmt.Println("Pi\t\t: ",math.Pi)
    fmt.Println("Pi MC\t\t: ",4*ndalam/float64(n))
    fmt.Println("Error(%)\t: ",math.Abs(math.Pi-(4*ndalam/float64(n)))/math.Pi*100)
}
