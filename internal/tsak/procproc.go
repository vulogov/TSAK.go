package tsak

import (
  "sync"
  // "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  // "github.com/vulogov/TSAK/internal/conf"
  "github.com/vulogov/TSAK/internal/signal"
  "github.com/vulogov/TSAK/internal/nr"
)

func ProcProc() {
  var start = nr.NowMillisec()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    procproc()
    log.Trace("Proc thread exiting")
    nr.RecordDuration("Proc() duration", start)
  }(signal.WG())
}

func procproc() {

}
