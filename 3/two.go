package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
  contents, err := ioutil.ReadFile("3/input.txt")

  if err != nil {
    log.Fatal(err)
  }

  array := strings.Fields(string(contents))
  length := len(array[0]) 
  oxygenArray := array
  co2Array := array
  idx := 0

  for idx < length {
    if len(oxygenArray) > 1 {
      oxygenArray = filterBitCriteria(oxygenArray, hasMoreZeros(oxygenArray, idx), idx)
    }
    if len(co2Array) > 1 {
      co2Array = filterBitCriteria(co2Array, !hasMoreZeros(co2Array, idx), idx)
    }
    idx++
  }

  oxygenGenerator, _ := strconv.ParseInt(oxygenArray[0], 2, 64)
  co2Scrubber, _ := strconv.ParseInt(co2Array[0], 2, 64)

  fmt.Println(oxygenGenerator * co2Scrubber)
}

func hasMoreZeros(array []string, index int) bool {
  zero := 0
  one := 0

  for _, value := range array {
    if value[index] == '0' {
      zero++
    } else {
      one++
    }
  }

  return zero > one
}

func filterBitCriteria (array []string, boolean bool, index int) []string {
  return filter(array, func(str string) bool {
    if (boolean) {
      return str[index] == '0'
    }
    return str[index] == '1'
  })
}

func filter(array []string, test func(string) bool) (ret []string) {
  for _, str := range array {
    if test(str) {
      ret = append(ret, str)
    }
  }
  return
}
