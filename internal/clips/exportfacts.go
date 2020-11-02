package clips

import (
    "fmt"
    "github.com/keysight/clipsgo/pkg/clips"
    "github.com/vulogov/TSAK/internal/log"
    "github.com/vulogov/TSAK/internal/piping"
    "github.com/Jeffail/gabs"
)

func procImpliedFact(f clips.Fact) string {
  var res []interface{}
  _out := gabs.New()
  f.Extract(&res)
  c := 0
  for _, _v := range res {
    _out.Set(_v, fmt.Sprintf("value%d", c))
    c += 1
  }
  return _out.String()
}

func ExportAllFacts(ch int) bool {
  var res string
  log.Trace(fmt.Sprintf("Exporting all facts to %d", ch))
  for _, f := range env.Facts() {
    if f.Template().Implied() {
      res = procImpliedFact(f)
      piping.To(ch, []byte(res))
    } else {

    }
  }
  return true
}
