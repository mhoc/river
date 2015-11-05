package river

func CommandExecuter() {
  go ExecuteCommands()
}

func ExecuteCommands() {
  var err error
  for {
    msg := <-FilteredCommands
    if EnableEJSONParsing {
      msg, err = ParseEasyJSONString(msg)
      if err != nil {
        Display <- NewConsoleMsg(ERROR, "Error parsing EJSON input")
        continue
      }
    }
    SendToServer <- msg
  }
}
