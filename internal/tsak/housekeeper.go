package tsak

import (
  "sync"
  "time"
  "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  "github.com/vulogov/TSAK/internal/signal"
  "github.com/vulogov/TSAK/internal/nr"
  "github.com/vulogov/TSAK/internal/piping"
  "github.com/vulogov/TSAK/internal/script"
  "github.com/vulogov/TSAK/internal/conf"
  "github.com/vulogov/TSAK/internal/cron"
)

var HOUSE_EVERY = (10 * time.Second)

func HouseProc() {
  var start = nr.NowMillisec()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    cron.Start()
    housekeeper()
    log.Trace("Housekeeper thread exiting")
    signal.ExitRequest()
    nr.RecordDuration("Housekeeper() duration", start)
  }(signal.WG())
}

func HouseShutdown() {
  log.Trace("Housekeeper terminating")
  cron.Stop()
}

func housekeeper() {
  if conf.Conf != "" {
    res := script.RunScript("house", conf.Conf)
    log.Event(
      "Bootstrap is loaded for Housekeeper()",
      logrus.Fields{
        "result":     res,
        "confSource": conf.Conf,
    })
  }
  for ! signal.ExitRequested() {
    time.Sleep(HOUSE_EVERY)
    housekeeperReport()
    if conf.House != "" {
      script.RunScript("house", conf.House)
    }
  }
  signal.ExitRequest()
}

func housekeeperReport() {
  nr.RecordValue("tsak.INCH.size", "Number of elements in TSAK pipes", piping.Len(piping.INCH))
  nr.RecordValue("tsak.OUTCH.size", "Number of elements in TSAK pipes", piping.Len(piping.OUTCH))
}
