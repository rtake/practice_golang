package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

type Result struct {
  word string
  t float64
}

// func startGame(length, num int) (res []Result) {
func startGame(length, num int, arrslice [26][]float64) ([26][]float64) {
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

  return arrslice
}

type Word struct {
  r rune
  f float64
}

/*
type Words []Word

func (ww Words) Len() int {
  return len(ww)
}

func (ww Words) Swap(i, j int) {
  ww[i], ww[j] = ww[j], ww[i]
}

func (ww Words) Less (i,j int) bool {
  return ww[i].f < ww[j].f
}
*/

func main() {
  arrslice := readfile()

  ///*
  length := 4
  num := 3
  arrslice = startGame(length, num, arrslice)

  Disp(arrslice, "all")
  //*/

  arr := CalcAverage(arrslice)

  var ww [26]Word
  for i:=0;i<26;i++{
    ww[i].r = rune(i+'a')
    ww[i].f = arr[i]
  }

  sort.Slice(ww[:], func(i, j int) bool {
    // return ww[i].f < ww[j].f
    return ww[i].f > ww[j].f
  })

  for i:=0;i<26;i++{
    fmt.Printf("%s,%.3fs\n", string(ww[i].r), ww[i].f)
  }

  writefile(arrslice)

}
