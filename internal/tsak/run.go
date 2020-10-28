package tsak

import (
  "github.com/vulogov/TSAK/internal/nr"
  // "github.com/vulogov/TSAK/internal/log"
  "github.com/vulogov/TSAK/internal/conf"
  "github.com/vulogov/TSAK/internal/signal"
)

func Run() {
  nr.RecordEvidence("Run() checkpoint is reached")
  if conf.Run != "" {
    signal.Reserve(1)
    nr.RecordEvidence("Exclsive Run() checkpoint reached")
    RunProc()
    nr.RecordEvidence("End of exclsive Run() checkpoint reached")
    return
  }
  if conf.In != "" {
    nr.RecordEvidence("Running In() code")
    signal.Reserve(1)
    InProc()
  }
  if conf.Proc != "" {
    nr.RecordEvidence("Running Proc() code")
    signal.Reserve(1)
    ProcProc()
  }
  if conf.Out != "" {
    nr.RecordEvidence("Running Proc() code")
    signal.Reserve(1)
    OutProc()
  }
  signal.Loop()
}
