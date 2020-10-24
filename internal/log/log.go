package log

import (
  "fmt"
  "os"
  "io"
  "github.com/sirupsen/logrus"
  "gopkg.in/natefinch/lumberjack.v2"
  "github.com/vulogov/Ushell/internal/conf"
  "github.com/newrelic/go-agent/v3/integrations/nrlogrus"
  newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

var log = logrus.New()
var app *newrelic.Application
var ljack *lumberjack.Logger
var writer io.Writer

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
    if conf.Nrapi != "" {
      app, _ = newrelic.NewApplication(
        newrelic.ConfigAppName(conf.Name),
        newrelic.ConfigLicense(conf.Nrapi),
        nrlogrus.ConfigLogger(log),
      )
    }
  } else {
    log.Formatter = new(logrus.TextFormatter)
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
    Trace(fmt.Sprintf("NRAPI %v", conf.Nrapi))
  }
  Trace("Log subsystem initialized")
}


func Trace(msg string) {
  if conf.Debug {
    params := map[string]interface{}{
              "source":     "tsak",
              "appname":    conf.Name,
              "appID":      conf.ID,
    }
    Log().WithFields(params).Trace(msg)
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
