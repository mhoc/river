
package river

type ConsoleMsgType int

const (
  SENDING ConsoleMsgType = iota
  RECEIVING ConsoleMsgType = iota
  ERROR ConsoleMsgType = iota
)

type ConsoleMsg struct {
  Type ConsoleMsgType
  Message string
}
