package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
  "log"
)

func main() {
  contents, err := ioutil.ReadFile("1/input.txt")

  if err != nil {
    log.Fatal(err)
  }

  array := toNumbersArray(strings.Fields(string(contents)))
  count := getSlidingWindowIncreasedCount(array, 3)

  fmt.Println(count)
}

func getSlidingWindowIncreasedCount(arr []int, k int) int {
  first := 0
  count := 0

  for index, value := range arr {
    first += value

    if index >= k - 1 && index + 1 < len(arr) {
      second := first - arr[index - (k - 1)] + arr[index + 1]

      if second > first {
        count++
      }
    }
  }

  return count
}

func toNumbersArray(arr []string) []int {
  result := make([]int, len(arr))

  for index, value := range arr {
    num, _ := strconv.Atoi(value)
    result[index] = num
  }

  return result
}
