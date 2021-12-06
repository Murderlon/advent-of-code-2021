package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
  "strconv"
)

func main() {
  contents, err := ioutil.ReadFile("3/input.txt")

  if err != nil {
    log.Fatal(err)
  }

  array := strings.Fields(string(contents))
  length := len(array[0]) 
  idx := 0
  gamma := ""

  for idx < length {
    gamma += iterate(array, idx)
    idx++
  }

  replacer := strings.NewReplacer("0", "1", "1", "0")
  epsilon, _ := strconv.ParseInt(replacer.Replace(gamma), 2, 64)
  gammaNumber, _ := strconv.ParseInt(gamma, 2, 64)

  fmt.Println(epsilon * gammaNumber)
}

func iterate(matrix []string, index int) string {
  zero := 0
  one := 0

  for _, value := range matrix {
    if value[index] == '0' {
      zero++
    } else {
      one++
    }
  }

  if zero > one {
    return "0"
  } else {
    return "1"
  }
}
