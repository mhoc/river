
package river

type ConsoleMsgType int

const (
  SENDING ConsoleMsgType = iota
  RECEIVING ConsoleMsgType = iota
  COMMAND ConsoleMsgType = iota
  ERROR ConsoleMsgType = iota
)

type ConsoleMsg struct {
  Type ConsoleMsgType
  Message string
}

func NewConsoleMsg(typ ConsoleMsgType, message string) ConsoleMsg {
  return ConsoleMsg{
    Type: typ,
    Message: message,
  }
}
