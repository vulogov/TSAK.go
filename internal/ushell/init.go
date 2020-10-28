package ushell

import (
  "flag"
  "os"
  "github.com/vulogov/Ushell/internal/log"
  "github.com/vulogov/Ushell/internal/conf"
  "github.com/vulogov/Ushell/internal/signal"
  "github.com/google/uuid"
)

func Init() {
  uid, _ := uuid.NewUUID()
  flag.BoolVar(&conf.Nocolor, "nocolor", false, "Disable colors in terminal output")
  flag.BoolVar(&conf.Debug, "debug", false, "Enable debug output")
  flag.BoolVar(&conf.Error, "error", false, "Enable ERROR outpout")
  flag.BoolVar(&conf.Warning, "warning", false, "Enable WARNING outpout")
  flag.BoolVar(&conf.Info, "info", false, "Enable INFO outpout")
  flag.BoolVar(&conf.Stdout, "stdout", false, "Send log entries to /dev/stdout as well")
  flag.BoolVar(&conf.Production, "production", false, "Running Ushell in production mode")
  flag.BoolVar(&conf.Nragent, "nragent", false, "NR infrastructure agent is present on the host")
  flag.StringVar(&conf.Nrapi, "nrapi", os.Getenv("NEW_RELIC_LICENSE_KEY"), "New Relic API key")
  flag.StringVar(&conf.Nriapi, "nriapi", os.Getenv("NEW_RELIC_INSERT_LICENSE_KEY"), "New Relic Insert API key")
  flag.StringVar(&conf.Logfile, "log", "", "Name of the log file")
  flag.StringVar(&conf.ID, "id", uid.String(), "Application unique identifier")
  flag.StringVar(&conf.Name, "name", uid.String(), "Name of the application")
  flag.StringVar(&conf.Logapi, "logapi", "https://log-api.newrelic.com/log/v1", "LOG API endpoint")
  flag.IntVar(&conf.Maxsize, "logsize", 100, "Maximum size of the log file in Mb")
  flag.IntVar(&conf.Maxage, "logage", 7, "Maximum age of the logfile in days")
  flag.Parse()
  log.InitLog()
  signal.InitSignal()
}
