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

  count := 0
  arr := strings.Fields(string(contents))

  for index, value := range arr {
    if index > 0 {
      current, _ := strconv.Atoi(value)
      previous, _ := strconv.Atoi(arr[index - 1])

      if current > previous {
        count++
      }
    }
  }

  fmt.Println(count)
}
