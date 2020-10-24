package ushell

import (
  "github.com/vulogov/Ushell/internal/signal"
)

func Ushell() {
  Init()
  signal.Loop()
}
