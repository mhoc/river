
package river

import (
  "fmt"
  "github.com/gorilla/websocket"
  "net/url"
  "os"
)

func RunWebsocket(host string) {
  u, err := FormatHost(host)
  if err != nil {
    fmt.Printf("Hostname provided is not formatted properly\n")
    os.Exit(1)
  }
  conn, _, err := websocket.DefaultDialer.Dial(u, nil)
  if err != nil {
    fmt.Printf("Error connecting to websocket client\n")
    fmt.Printf("%v\n", err.Error())
    os.Exit(1)
  }
  go WsWriter(conn)
  go WsReader(conn)
}

func FormatHost(hostname string) (string, error) {
  if hostname[:6] != "ws://" {
    hostname = "ws://" + hostname
  }
  u, err := url.Parse(hostname)
  return u.String(), err
}

func WsWriter(c *websocket.Conn) {
  for {
    msg := <-SendToServer
    Display <- NewConsoleMsg(SENDING, msg)
    err := c.WriteMessage(websocket.TextMessage, []byte(msg))
    if err != nil {
      Display <- NewConsoleMsg(ERROR, err.Error())
    }
  }
}

func WsReader(c *websocket.Conn) {
  for {
    _, msg, err := c.ReadMessage()
    if err != nil {
      Display <- NewConsoleMsg(ERROR, err.Error())
    } else {
      Display <- NewConsoleMsg(RECEIVING, string(msg))
    }
  }
}
