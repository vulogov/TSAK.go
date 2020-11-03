package tsak

import (
  "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  "github.com/vulogov/TSAK/internal/piping"
  "github.com/vulogov/TSAK/internal/clips"
)

func Fin() {
  piping.Shutdown()
  clips.Shutdown()
  log.Shutdown()
  HouseShutdown()
  log.Event("TsakEvent", logrus.Fields{
    "message":    "Application exited",
    "evtc":       1,
  })
}
