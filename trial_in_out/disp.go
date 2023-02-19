package main

import (
	"fmt"
	"sort"
)

func Disp(arrslice [26][]float64, key string) {
  fmt.Println("===== Result =====")

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

  fmt.Println("==================")
}

