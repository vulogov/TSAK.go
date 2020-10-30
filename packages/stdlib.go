package packages

import (
  "time"
  "reflect"
  "github.com/google/uuid"
  "github.com/mattn/anko/env"
  "github.com/erikdubbelboer/gspt"
  "github.com/vulogov/TSAK/internal/signal"
  "github.com/vulogov/TSAK/internal/piping"
)

func NowMilliseconds() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func UUID() string {
  uid, _ := uuid.NewUUID()
  return uid.String()
}

func String(src []byte) string {
  return string(src)
}

func init() {
  env.Packages["stdlib"] = map[string]reflect.Value{
    "Answer":         reflect.ValueOf(42),
    "SetProcTitle":   reflect.ValueOf(gspt.SetProcTitle),
    "ExitRequest":    reflect.ValueOf(signal.ExitRequest),
    "ExitRequested":  reflect.ValueOf(signal.ExitRequested),
    "Release":        reflect.ValueOf(signal.Release),
    "NowMilliseconds":reflect.ValueOf(NowMilliseconds),
    "UUID":           reflect.ValueOf(UUID),
    "From":           reflect.ValueOf(piping.From),
    "To":             reflect.ValueOf(piping.To),
    "Len":            reflect.ValueOf(piping.Len),
    "INCH":           reflect.ValueOf(piping.INCH),
    "OUTCH":          reflect.ValueOf(piping.OUTCH),
    "CLIPS":          reflect.ValueOf(piping.CLIPS),
    "String":         reflect.ValueOf(String),
  }
}
