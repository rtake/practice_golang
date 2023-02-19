package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"strconv"
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

func writefile(arrslice [26][]float64) {
  f, _ := os.Create("result.txt")

  var str string

  for i:=0;i<26;i++{
    for _, v := range arrslice[i]{
      str = str + fmt.Sprintf("%s,%f\n", string('a'+i), v)
    }
  }

  data := []byte(str)
  _, _ = f.Write(data)

  return 
}

func readfile() (arrslice [26][]float64) {
  f, _ := os.Open("result.txt")

  data := make([]byte, 1024)
  cnt, err := f.Read(data)
  if err != nil {
    fmt.Println(err)
    fmt.Println("fail to read file")
  }

  data = data[:cnt]
  str := string(data)

  arr := strings.Split(str, "\n")

  for _, v := range arr{
    s := strings.Split(v, ",")

    if len(s) < 2 {
      break
    }

    // fmt.Println(s[0],s[1])

    r := []rune(s[0])
    idx := int(r[0]-'a')
    val, _ := strconv.ParseFloat(s[1], 64)
    arrslice[idx] = append(arrslice[idx], val)
  }

  return
}

func main() {
  arrslice := readfile()

  length := 4
  num := 3
  arrslice = startGame(length, num, arrslice)

  Disp(arrslice, "all")
  // CalcAverage(arrslice)

  writefile(arrslice)

}
