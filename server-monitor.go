package main

import (
  "os"
  "fmt"
  "github.com/jjoos/process-status-collector"
)

func main() {
	os.Stdout.WriteString("Hello World!\n")
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", process_status_collector.Sqrt(2))
}
