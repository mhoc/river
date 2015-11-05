
package river

import (
  ui "github.com/gizak/termui"
)

func CreateSendBox() *ui.Par {
  p := ui.NewPar("")
  p.Height = 3
  p.Width = 50
  p.TextFgColor = ui.ColorWhite
  p.BorderLabel = "Send"
  p.BorderFg = ui.ColorCyan
  SendBoxEvents(p)
  return p
}

func SendBoxEvents(p *ui.Par) {

  ui.Handle("/sys/kbd/<enter>", func(e ui.Event) {
    Input <- p.Text
    p.Text = ""
    ui.Render(ui.Body)
  })

  ui.Handle("/sys/kbd/<space>", func(e ui.Event) {
    p.Text += " "
    ui.Render(ui.Body)
  })

  ui.Handle("/sys/kbd/C-8", func(e ui.Event) {
    if len(p.Text) > 0 {
      p.Text = p.Text[:len(p.Text)-1]
      ui.Render(ui.Body)
    }
  })

  ui.Handle("/sys/kbd", func(e ui.Event) {
    char := e.Data.(ui.EvtKbd).KeyStr
    p.Text += char
    ui.Render(ui.Body)
  })

}
