package clips

import (
    "time"
    "sync"
    "fmt"
    "github.com/keysight/clipsgo/pkg/clips"
    "github.com/vulogov/TSAK/internal/log"
    "github.com/vulogov/TSAK/internal/signal"
    "github.com/vulogov/TSAK/internal/nr"
    "github.com/vulogov/TSAK/internal/piping"
)

var env = clips.CreateEnvironment()

func ReinitCLIPS() {
  env.Reset()
  env.Clear()
}

func InitClips() {
  log.Trace("CLIPS module is Initializing")
  ReinitCLIPS()
  signal.Reserve(1)
  go func(wg *sync.WaitGroup) {
    var start = nr.NowMillisec()
    defer wg.Done()
    clipsproc()
    log.Trace("CLIPS thread exiting")
    nr.RecordDuration("CLIPS() duration", start)
  }(signal.WG())
}

func Shutdown() {
  log.Trace("CLIPS module is shutting down")
  env.Delete()
}

func clipsproc() {
  for ! signal.ExitRequested() && piping.Len(piping.CLIPS) == 0 {
    log.Trace("CLIPS server is cooling down")
    time.Sleep(1*time.Second)
    for piping.Len(piping.CLIPS) > 0 {
      cmd := string(piping.From(piping.CLIPS))
      env.SendCommand(cmd)
      log.Trace(fmt.Sprintf("CLIPS: %v", cmd))
    }
  }
}
