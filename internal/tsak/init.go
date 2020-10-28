package tsak

import (
  "flag"
  "os"
  "github.com/sirupsen/logrus"
  "github.com/vulogov/TSAK/internal/log"
  "github.com/vulogov/TSAK/internal/conf"
  "github.com/vulogov/TSAK/internal/signal"
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
  flag.StringVar(&conf.Nrapi, "nrapi", os.Getenv("NEW_RELIC_LICENSE_KEY"), "New Relic API key")
  flag.StringVar(&conf.Logfile, "log", "", "Name of the log file")
  flag.StringVar(&conf.ID, "id", uid.String(), "Application unique identifier")
  flag.StringVar(&conf.Name, "name", uid.String(), "Name of the application")
  flag.StringVar(&conf.Account, "account", os.Getenv("NEW_RELIC_ACCOUNT"), "New Relic user account number")
  flag.StringVar(&conf.Logapi, "logapi", "https://log-api.newrelic.com/log/v1", "LOG API endpoint")
  flag.StringVar(&conf.Evtapi, "evtapi", "https://insights-collector.newrelic.com/v1/accounts/%s/events", "EVT API endpoint")
  flag.IntVar(&conf.Maxsize, "logsize", 100, "Maximum size of the log file in Mb")
  flag.IntVar(&conf.Maxage, "logage", 7, "Maximum age of the logfile in days")
  flag.Parse()
  log.InitLog()
  signal.InitSignal()
  log.Event("TsakEvent", logrus.Fields{
    "event":    "Application started",
    "evtc":     0,
  })
}
