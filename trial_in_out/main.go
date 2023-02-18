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

/*
type Score struct {
  c rune
  t float64
}

type Scores []Score

func (ss Scores) Len() int {
  return len(ss)
}

func (ss Scores) Swap(i, j int) {
  ss[i], ss[j] = ss[j], ss[i]
}

func (ss Scores) Less(i, j int) bool {
  // return ss[i].t < ss[j].t
  return ss[i].t > ss[j].t
}
*/

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

  /*
  m := make(map[rune]float64, 10)

  for _, v := range res {
    for _, c := range v.word {
      val, isThere := m[c]

      if isThere {
        m[c] = math.Min(val, v.t)
      } else {
	m[c] = v.t
      }
    }

    fmt.Printf("%s %.3fs\n", v.word, v.t)
  }


  var ss Scores

  for k, v := range m {
    ss = append(ss, Score{k, v})
  }

  sort.Sort(ss)

  for _, r := range ss {
    fmt.Printf("%s %.3fs\n", string(r.c), r.t)
  }
  */
}
