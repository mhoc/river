
package river

import (
  "fmt"
  "github.com/gorilla/websocket"
  "net/url"
  "os"
)

func RunWebsocket(host string) {
  u, err := url.Parse(host)
  if err != nil {
    fmt.Printf("Hostname provided is not formatted properly\n")
    os.Exit(1)
  }
  conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
  if err != nil {
    fmt.Printf("Error connecting to websocket client\n")
    fmt.Printf("%v\n", err.Error())
    os.Exit(1)
  }
  go WsWriter(conn)
  go WsReader(conn)
}

func WsWriter(c *websocket.Conn) {
  for {
    msg := <-Submissions
    err := c.WriteMessage(websocket.TextMessage, []byte(msg))
    if err != nil {
      Receives <- ConsoleMsg{
        Type: ERROR,
        Message: err.Error(),
      }
    }
  }
}

func WsReader(c *websocket.Conn) {
  for {
    _, msg, err := c.ReadMessage()
    if err != nil {
      Receives <- ConsoleMsg{
        Type: ERROR,
        Message: err.Error(),
      }
    } else {
      Receives <- ConsoleMsg{
        Type: RECEIVING,
        Message: string(msg),
      }
    }
  }
}
