package main

import (
)

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

  /*
  for i:=0;i<26;i++ {
    fmt.Printf("%s %.3f s\n", string('a'+i), avearr[i])
  }
  */

  return 
}

