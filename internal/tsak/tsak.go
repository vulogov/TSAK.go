package tsak

import (
  "github.com/vulogov/TSAK/internal/signal"
)

func TSAK() {
  Init()
  signal.Loop()
}
