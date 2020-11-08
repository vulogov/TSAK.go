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

var HOUSE_EVERY = (1 * time.Second)
var REPORT_EVERY = 15

func HouseProc() {
  var start = nr.NowMillisec()
  cron.Start()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
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
  c := 0
  for ! signal.ExitRequested() {
    time.Sleep(HOUSE_EVERY)
    if c > REPORT_EVERY {
      log.Trace("Running housekeeper")
      housekeeperReport()
      if conf.House != "" {
        script.RunScript("house", conf.House)
      }
      c = 0
    } else {
      c += 1
    }
  }
  signal.ExitRequest()
}

func housekeeperReport() {
  nr.RecordValue("tsak.INCH.size", "Number of elements in TSAK pipes", piping.Len(piping.INCH))
  nr.RecordValue("tsak.OUTCH.size", "Number of elements in TSAK pipes", piping.Len(piping.OUTCH))
}
