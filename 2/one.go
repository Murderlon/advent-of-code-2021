package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
  contents, err := ioutil.ReadFile("2/input.txt")

  if err != nil {
    log.Fatal(err)
  }

  array := strings.Split(string(contents), "\n")
  horizontal := 0
  depth := 0

  for _, value := range array {
    if len(value) == 0 {
      continue
    }

    values := strings.Fields(value)
    units, _ := strconv.Atoi(values[1])

    switch values[0] {
      case "forward":
        horizontal += units
      case "down":
        depth += units
      case "up":
        depth -= units
    }
  }

  fmt.Println(horizontal, "*", depth, "=", horizontal * depth)
}
