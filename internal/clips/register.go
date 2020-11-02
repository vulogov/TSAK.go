package clips

import (
    "github.com/vulogov/TSAK/internal/log"
    "github.com/vulogov/TSAK/internal/nr"
    "github.com/keysight/clipsgo/pkg/clips"
)

func RegisterVariables() {
  GetVarBindDef(clips.Symbol("INCH"), 0)
  GetVarBindDef(clips.Symbol("OUTCH"), 1)
  GetVarBindDef(clips.Symbol("CLIPS"), 2)
  GetVarBindDef(clips.Symbol("FACTS"), 3)
  GetVarBindDef(clips.Symbol("Answer"), 42)
}

func RegisterFunctions() {
  log.Trace("CLIPS functions registering")
  AddClipsFun("Now", nr.NowMillisec)
  AddClipsFun("exportallfacts", ExportAllFacts)
  AddClipsFun("var", SetVarBind)
  AddClipsFun("getvar", GetVarBind)
  AddClipsFun("VAR", GetVarBindDef)
  AddClipsFun("enablefactpipe", EnableFactPipe)
  AddClipsFun("disablefactpipe", DisableFactPipe)
  AddClipsFun("enablecmdpipe", EnableCmdPipe)
  AddClipsFun("disablecmdpipe", DisableCmdPipe)
  RegisterVariables()
}
