
package main

import (
  "gopkg.in/alecthomas/kingpin.v2"
  "river"
  "runtime"
)

var (
  hostname = kingpin.Arg("hostname", "the url of the websocket you are connecting to").Required().String()
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  kingpin.Parse()
  river.RunWebsocket(*hostname)
  river.CommandHandler()
  river.CommandExecuter()
  river.BuildUI()
}
