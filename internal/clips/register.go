package clips

import (
    "github.com/vulogov/TSAK/internal/log"
    "github.com/vulogov/TSAK/internal/nr"
)


func RegisterFunctions() {
  log.Trace("CLIPS functions registering")
  AddClipsFun("Now", nr.NowMillisec)
}
