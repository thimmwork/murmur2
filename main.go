// just a main function
package main

import (
    "fmt"
    "os"
    "strconv"
)

func toPositive(input uint32) (pos uint32) {
  pos = input & 0x7fffffff
  return
}

func partition(input uint32, noOfPartitions uint32) (part uint32) {
  part = input % noOfPartitions
  return
}

func keyAndPartitions(args []string) (key string, noOfPartitions uint32) {
  noOfPartitions = 60
  if (len(os.Args) == 4) {
    if (os.Args[1] == "-n") {
      noOfPartitions64, _ := strconv.ParseUint(os.Args[2], 0, 32)
      noOfPartitions = uint32(noOfPartitions64)
      key = os.Args[3]
    } else {
      fmt.Println("use -n to give number of partitions")
    }
  } else {
    key = os.Args[1]
  }
  return
}

func main() {
  if (len(os.Args) == 1) {
    fmt.Println("no args given")
    return
  }

  key, noOfPartitions := keyAndPartitions(os.Args)
  fmt.Println(partition(toPositive(MurmurHash2([]byte(key))), noOfPartitions))
}

