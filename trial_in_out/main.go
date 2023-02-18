package main

import (
	"fmt"
	// "math"
	"sort"
	"time"
	"math/rand"
)

type Result struct {
  word string
  t float64
}

func startGame(length, num int) ([]Result) {
  var res []Result
  var str string

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

  return res
}

func main() {
  length := 5
  num := 3
  res := startGame(length, num)

  var arrslice [26][]float64

  for _, r := range res {
    for _, c := range r.word {
      idx := int(c-'a')
      arrslice[idx] = append(arrslice[idx], float64(r.t))
    }
  }

  for i:= 0;i<26;i++ {
    if len(arrslice[i]) == 0 {
      arrslice[i] = append(arrslice[i], 0)
    }

    fmt.Printf("%s", string(rune(i+'a')))

    sort.Float64s(arrslice[i])

    // All elements
    for _, v := range arrslice[i] {
      fmt.Printf(",%.3f", v)
    }

    // Max
    /*
    fmt.Printf(" %.3f", arrslice[i][0])
    */

    fmt.Printf("\n")
  }

}
