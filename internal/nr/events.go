package nr

import (
  "github.com/sirupsen/logrus"
)


func RecordDuration(msg string, start int64) {
  fn, line, fun := Trace()
  Event("TsakEvent", logrus.Fields{
    "message":    msg,
    "evtc":       4,
    "duration":   NowMillisec() - start,
    "source":     fn,
    "line":       line,
    "function":   fun,
  })
}

func RecordEvidence(msg string) {
  fn, line, fun := Trace()
  Event("TsakEvent", logrus.Fields{
    "message":    msg,
    "evtc":       2,
    "source":     fn,
    "line":       line,
    "function":   fun,
  })
}
