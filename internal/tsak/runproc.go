package tsak

import (
  "sync"
  // "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  // "github.com/vulogov/TSAK/internal/conf"
  "github.com/vulogov/TSAK/internal/signal"
  "github.com/vulogov/TSAK/internal/nr"
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

}
