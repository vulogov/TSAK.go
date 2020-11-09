package packages

import (
  "fmt"
  "github.com/k-sone/snmpgo"
  "reflect"
  "github.com/mattn/anko/env"
)

func SNMPv1Get(addr string, community string, _oid string, retry uint) interface{} {
  snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V1,
		Address:   addr,
		Retries:   retry,
		Community: community,
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
  oids, err := snmpgo.NewOids([]string{
    _oid,
  })
  if err != nil {
		// Failed to parse Oids
		fmt.Println(err)
		return false
	}

	if err = snmp.Open(); err != nil {
		// Failed to open connection
		fmt.Println(err)
		return false
	}
	defer snmp.Close()

	pdu, err := snmp.GetRequest(oids)
	if err != nil {
		// Failed to request
		fmt.Println(err)
		return false
	}
	if pdu.ErrorStatus() != snmpgo.NoError {
		// Received an error from the agent
		fmt.Println(pdu.ErrorStatus(), pdu.ErrorIndex())
	}

	res := pdu.VarBinds().MatchOid(oids[0])
  return res.Variable
}

func SNMPv2cGet(addr string, community string, _oid string, retry uint) interface{} {
  snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V2c,
		Address:   addr,
		Retries:   retry,
		Community: community,
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
  oids, err := snmpgo.NewOids([]string{
    _oid,
  })
  if err != nil {
		// Failed to parse Oids
		fmt.Println(err)
		return false
	}

	if err = snmp.Open(); err != nil {
		// Failed to open connection
		fmt.Println(err)
		return false
	}
	defer snmp.Close()

	pdu, err := snmp.GetRequest(oids)
	if err != nil {
		// Failed to request
		fmt.Println(err)
		return false
	}
	if pdu.ErrorStatus() != snmpgo.NoError {
		// Received an error from the agent
		fmt.Println(pdu.ErrorStatus(), pdu.ErrorIndex())
	}

	res := pdu.VarBinds().MatchOid(oids[0])
  return res.Variable
}

func init() {
  env.Packages["snmp/client"] = map[string]reflect.Value{
    "Getv1":    reflect.ValueOf(SNMPv1Get),
    "Getv2c":    reflect.ValueOf(SNMPv2cGet),
  }
  env.PackageTypes["snmp/client"] = map[string]reflect.Type{

  }
}
