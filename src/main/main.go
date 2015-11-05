
package main

import (
  "fmt"
  "os"
  "river"
  "runtime"
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  if len(os.Args) != 2 {
    fmt.Printf("Must provide a hostname\n")
    os.Exit(1)
  }
  hostname := os.Args[1]
  river.RunWebsocket(hostname)
  river.CommandHandler()
  river.CommandExecuter()
  river.BuildUI()
}
