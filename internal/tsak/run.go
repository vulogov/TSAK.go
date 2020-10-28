package tsak

import (
  "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  // "github.com/vulogov/TSAK/internal/conf"
  // "github.com/vulogov/TSAK/internal/signal"
)

func Run() {
  log.Event("TsakEvent", logrus.Fields{
    "message":    "Run() loop reached",
    "evtc":       2,
  })
}
