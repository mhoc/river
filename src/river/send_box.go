
package river

import (
  ui "github.com/gizak/termui"
)

var (
  Submissions = make(chan string, 10)
)

func CreateSendBox() *ui.Par {
  SendBox := ui.NewPar("")
  SendBox.Height = 3
  SendBox.Width = 50
  SendBox.TextFgColor = ui.ColorWhite
  SendBox.BorderLabel = "Send"
  SendBox.BorderFg = ui.ColorCyan
  SendBoxEvents(SendBox)
  return SendBox
}

func SendBoxEvents(SendBox *ui.Par) {

  ui.Handle("/sys/kbd/<enter>", func(e ui.Event) {
    if SendBox.Text[0] == '/' {
      Receives <- ConsoleMsg{
        Type: COMMAND,
        Message: SendBox.Text[1:],
      }
    } else {
      Receives <- ConsoleMsg{
        Type: SENDING,
        Message: SendBox.Text,
      }
      Submissions <- SendBox.Text
    }
    SendBox.Text = ""
    ui.Render(ui.Body)
  })

  ui.Handle("/sys/kbd/<space>", func(e ui.Event) {
    SendBox.Text += " "
    ui.Render(ui.Body)
  })

  ui.Handle("/sys/kbd/C-8", func(e ui.Event) {
    if len(SendBox.Text) > 0 {
      SendBox.Text = SendBox.Text[:len(SendBox.Text)-1]
      ui.Render(ui.Body)
    }
  })

  ui.Handle("/sys/kbd", func(e ui.Event) {
    char := e.Data.(ui.EvtKbd).KeyStr
    SendBox.Text += char
    ui.Render(ui.Body)
  })

}
