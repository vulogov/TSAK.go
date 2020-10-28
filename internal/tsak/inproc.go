package tsak

import (
  "sync"
  // "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  // "github.com/vulogov/TSAK/internal/conf"
  "github.com/vulogov/TSAK/internal/signal"
  "github.com/vulogov/TSAK/internal/nr"
)

func InProc() {
  var start = nr.NowMillisec()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    inproc()
    log.Trace("In thread exiting")
    nr.RecordDuration("In() duration", start)
  }(signal.WG())
}

func inproc() {

}
