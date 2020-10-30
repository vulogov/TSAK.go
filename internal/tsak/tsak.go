package tsak

import (
  "github.com/vulogov/TSAK/internal/signal"
)

func TSAK() {
  Init()
  Run()
  signal.Loop()
  Fin()
}
