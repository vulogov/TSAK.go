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

func RunProc() {
  var start = nr.NowMillisec()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    runproc()
    log.Trace("Run thread exiting")
    nr.RecordDuration("Run() duration", start)
  }(signal.WG())
}

func runproc() {
  if conf.Conf != "" {
    res := script.RunScript("proc", conf.Conf)
    log.Event(
      "Bootstrap is loaded for Run()",
      logrus.Fields{
        "result":     res,
        "confSource": conf.Conf,
    })
  }
  script.RunScript("proc", conf.Run)
}
