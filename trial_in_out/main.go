package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Result struct {
  word string
  t float64
}

// func startGame(length, num int) (res []Result) {
func startGame(length, num int) (arrslice [26][]float64) {
  var str string
  var res []Result

  rand.Seed(time.Now().UnixNano())

  for i := 0; i < num; i++ {
    // Gen problem
    var ans string
    for j := 0; j < length; j++ {
      ans += string(rune(int('a')+rand.Intn(26)))
    }

    // 
    fmt.Println("=================")

    now := time.Now()

    for {
      fmt.Println(ans)
      // fmt.Println("=================")

      fmt.Scan(&str)

      if str == ans {
	res = append(res, Result{ans, time.Since(now).Seconds()})
	break
      }

      fmt.Println("Miss!")
    }
  }


  for _, r := range res {
    for _, c := range r.word {
      idx := int(c-'a')
      arrslice[idx] = append(arrslice[idx], float64(r.t))
    }
  }

  return
}

func main() {
  length := 4
  num := 3

  arrslice := startGame(length, num)

  Disp(arrslice, "all")

  // CalcAverage(arrslice)

}
