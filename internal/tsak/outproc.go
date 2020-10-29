package tsak

import (
  "sync"
  "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  "github.com/vulogov/TSAK/internal/script"
  "github.com/vulogov/TSAK/internal/signal"
  "github.com/vulogov/TSAK/internal/nr"
  "github.com/vulogov/TSAK/internal/conf"
)

func OutProc() {
  var start = nr.NowMillisec()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    outproc()
    log.Trace("Out thread exiting")
    nr.RecordDuration("Out() duration", start)
  }(signal.WG())
}

func outproc() {
  if conf.Conf != "" {
    res := script.RunScript("out", conf.Conf)
    log.Event(
      "Bootstrap is loaded for Out()",
      logrus.Fields{
        "result":     res,
        "confSource": conf.Conf,
    })
  }
  script.RunScript("out", conf.Out)
}
