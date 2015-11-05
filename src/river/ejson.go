package river

import (
  "encoding/json"
  "fmt"
  "strconv"
  "strings"
)

// Parses an easyjson string into a real json string
// easyjson is a minified json variant modeled after httpie's json shortcuts to make typing a bit faster
// for now it can only support objects that look like this
//    id:5 name:michael
// spaces in strings are not allowed. nested objects are not allowed. arrays are not allowed.
// I'd like to expand this later
func ParseEasyJSONString(st string) (string, error) {
  data := make(map[string]interface{})
  sp := strings.Split(st, " ")
  for _, kv := range sp {
    kvsp := strings.Split(kv, ":")
    if len(kvsp) < 2 {
      return "", fmt.Errorf("River couldn't parse your input as EJSON")
    }
    data[kvsp[0]] = ParseValueType(kvsp[1])
  }
  b, err := json.Marshal(data)
  return string(b), err
}

func ParseValueType(value string) interface{} {
  if value == "true" {
    return true
  }
  if value == "false" {
    return false
  }
  i, err := strconv.Atoi(value)
  if err != nil {
    return value
  } else {
    return i
  }
}
