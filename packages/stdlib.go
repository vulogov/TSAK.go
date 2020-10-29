package packages

import (
  "time"
  "reflect"
  "github.com/google/uuid"
  "github.com/mattn/anko/env"
  "github.com/erikdubbelboer/gspt"
  "github.com/vulogov/TSAK/internal/signal"
)

func NowMilliseconds() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func UUID() string {
  uid, _ := uuid.NewUUID()
  return uid.String()
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
  }
}