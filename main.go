// just a main function
package main

import (
    "fmt"
    "os"
)

func toPositive(input uint32) (pos uint32) {
  pos = input & 0x7fffffff
  return
}

func partition(input uint32) (part uint32) {
  part = input % 60
  return
}

func main() {
    if (len(os.Args) == 1) {
      fmt.Println("no args given")
    } else {
      arg := os.Args[1]
      fmt.Println(partition(toPositive(MurmurHash2([]byte(arg)))))
    }
}

