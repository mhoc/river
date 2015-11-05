package river

var (
  EnableJSONParsing = false
)

func DispatchCommand(command string) {
  if command == "json" {
    EnableJSONParsing = !EnableJSONParsing
  }
}
