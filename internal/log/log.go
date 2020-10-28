package log

import (
  "fmt"
  "os"
  "io"
  "github.com/sirupsen/logrus"
  "gopkg.in/natefinch/lumberjack.v2"
  "github.com/vulogov/TSAK/internal/conf"
  "github.com/vulogov/TSAK/internal/nr"
)

type Fields logrus.Fields

var log = logrus.New()
var ljack *lumberjack.Logger
var writer io.Writer
var Event = nr.Event

func InitLog() {
  if len(conf.Logfile) > 0 {
    ljack = &lumberjack.Logger{
      Filename:   conf.Logfile,
      MaxSize:    conf.Maxsize,
      MaxAge:     conf.Maxage,
      Compress:   false,
    }

    if conf.Stdout {
        writer = io.MultiWriter(os.Stdout, ljack)
    } else {
       writer = io.MultiWriter(ljack)
    }
  } else {
    if conf.Stdout {
      writer = io.MultiWriter(os.Stdout)
    } else {
      writer = io.MultiWriter()
    }
  }
  if conf.Production {
    log.SetFormatter(&logrus.JSONFormatter{})
    log.Level = logrus.TraceLevel
    log.SetOutput(writer)
  } else {
    log.SetFormatter(&logrus.TextFormatter{})
    if conf.Nocolor {
      log.Formatter.(*logrus.TextFormatter).DisableColors = true
    } else {
      log.Formatter.(*logrus.TextFormatter).DisableColors = false
    }
    log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
    log.Formatter.(*logrus.TextFormatter).FullTimestamp = true
    log.Level = logrus.TraceLevel
    log.SetOutput(writer)
  }
  if conf.Debug {
    conf.Info = true
    conf.Warning = true
    conf.Error = true
  }
  if conf.Info {
    conf.Info = true
    conf.Warning = true
    conf.Error = true
  }
  if conf.Warning {
    conf.Warning = true
    conf.Error = true
  }
  Trace(fmt.Sprintf("Production level: %v", conf.Production))
  Trace(fmt.Sprintf("Maximum size of log file (Mb): %v", conf.Maxsize))
  Trace(fmt.Sprintf("Maximum age of log file (days): %v", conf.Maxage))
  Trace(fmt.Sprintf("Application UUID %v", conf.ID))
  if conf.Nrapi != "" {
    Trace(fmt.Sprintf("NRAPI Enabled"))
  }
  Trace("Log subsystem initialized")
}


func Trace(msg string, ctx ...logrus.Fields) {
  var c logrus.Fields
  if conf.Debug {
    if len(ctx) > 0 {
      c = ctx[0]
    } else {
      c = logrus.Fields{}
    }
    Log().WithFields(c).Trace(msg)
    if conf.Nrapi != "" {
      if conf.Production {
        nr.Log(msg, "trace", c)
      }
    }
  }
}

func Info(msg string) {
  if conf.Info {
    params := map[string]interface{}{
              "source":     "tsak",
              "appname":    conf.Name,
              "appID":      conf.ID,
    }
    Log().WithFields(params).Info(msg)
  }
}

func Warning(msg string) {
  if conf.Warning {
    log.Warning(msg)
  }
}

func Error(msg string) {
  if conf.Error {
    log.Error(msg)
  }
}

func Log() *logrus.Logger {
    return log
}

func Shutdown() {
  Trace("Log subsystem shutodown")
}
