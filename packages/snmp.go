package packages

import (
  "net"
  "time"
  "github.com/deejross/go-snmplib"
  "github.com/vulogov/TSAK/internal/snmp"
  "reflect"
  "github.com/mattn/anko/env"
)

func UnmarshalTrap(sock *net.IPConn, b []byte, n int) snmplib.Trap {
  snmp := snmplib.NewSNMPOnConn("", "", snmplib.SNMPv2c, 2*time.Second, 5, sock)
  msg := b[:n]
  varbinds, _ := snmp.ParseTrap(msg)
  return varbinds
}

func init() {
  env.Packages["snmp"] = map[string]reflect.Value{
    "ParseTrap":  reflect.ValueOf(UnmarshalTrap),
    "InitMib":    reflect.ValueOf(snmp.InitMib),
    "LoadModule": reflect.ValueOf(snmp.LoadModule),
    "OID":        reflect.ValueOf(snmp.OID),
    "SYMBOL":     reflect.ValueOf(snmp.SYMBOL),
    "Client":     reflect.ValueOf(snmplib.NewSNMP),
    "SNMPv1":     reflect.ValueOf(snmplib.SNMPv1),
    "SNMPv2c":    reflect.ValueOf(snmplib.SNMPv2c),
    "SNMPv3":     reflect.ValueOf(snmplib.SNMPv3),
    "ParseOID":   reflect.ValueOf(snmplib.MustParseOid),
  }
  env.PackageTypes["snmp"] = map[string]reflect.Type{
    "SNMP":          reflect.TypeOf(snmplib.SNMP{}),
  }
}
