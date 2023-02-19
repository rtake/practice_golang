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

func startGame(length, num int) (res []Result) {
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


func disp(arrslice [26][]float64, key string) {
  for i:= 0;i<26;i++ {
    if len(arrslice[i]) == 0 {
      arrslice[i] = append(arrslice[i], 0)
    }

    fmt.Printf("%s", string(rune(i+'a')))

    sort.Float64s(arrslice[i])


    if key == "max" {
      fmt.Printf(" %.3f", arrslice[i][0])
      fmt.Printf("\n")
    } else {
      for _, v := range arrslice[i] {
        fmt.Printf(",%.3f", v)
      }

      fmt.Printf("\n")
    }

  }
}

func CalcAverage(arrslice [26][]float64) (avearr [26]float64) {

  for i:=0;i<26;i++ {
    sum := float64(0)

    for j:=0;j<len(arrslice[i]);j++{
      sum = sum + arrslice[i][j]
    }

    if len(arrslice[i]) == 0 {
      avearr[i] = 0
    } else {
      avearr[i] = sum/float64(len(arrslice[i]))
    }
  }

  return 
}

func main() {
  length := 4
  num := 3
  res := startGame(length, num)

  var arrslice [26][]float64

  for _, r := range res {
    for _, c := range r.word {
      idx := int(c-'a')
      arrslice[idx] = append(arrslice[idx], float64(r.t))
    }
  }

  disp(arrslice, "all")

  avearr := CalcAverage(arrslice)
  for i:=0;i<26;i++ {
    fmt.Printf("%s %fs\n", string('a'+i), avearr[i])
  }

}
