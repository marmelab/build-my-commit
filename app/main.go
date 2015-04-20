package main

import "os"
import "fmt"

func main() {
  fmt.Println("FOO:", os.Getenv("FOO"))
}
