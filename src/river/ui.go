
package river

import (
  ui "github.com/gizak/termui"
)

func BuildUI() {
  ui.Init()
  defer ui.Close()

  receiveBox := CreateReceiveBox()
  sendBox := CreateSendBox()
  ui.Body.AddRows(
    ui.NewRow(ui.NewCol(12, 0, receiveBox)),
    ui.NewRow(ui.NewCol(12, 0, sendBox)),
  )

  ui.Body.Align()
  ui.Render(ui.Body)

  ui.Handle("/sys/kbd/C-x", func(e ui.Event) {
    ui.StopLoop()
  })

  ui.Handle("/timer/1s", func(e ui.Event) {
    ui.Body.Align()
    ui.Render(ui.Body)
  })

  ui.Loop()
}
