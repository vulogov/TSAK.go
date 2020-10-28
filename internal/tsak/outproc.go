package tsak

import (
  "sync"
  // "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  // "github.com/vulogov/TSAK/internal/conf"
  "github.com/vulogov/TSAK/internal/signal"
  "github.com/vulogov/TSAK/internal/nr"
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

}
