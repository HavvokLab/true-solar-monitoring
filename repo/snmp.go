package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/gosnmp/gosnmp"
)

type SnmpRepo interface {
	SendAlarmTrap(deviceName, alertName, description, severity, lastedUpdateTime string) error
	Close()
}

type snmpRepo struct {
	client     *gosnmp.GoSNMP
	agentHost  string
	targetHost string
	targetPort uint16
}

func NewSnmpRepo(client *gosnmp.GoSNMP) SnmpRepo {
	conf := config.GetConfig().Snmp
	return &snmpRepo{
		client:     client,
		agentHost:  conf.AgentHost,
		targetHost: conf.TargetHost,
		targetPort: conf.TargetPort,
	}
}

func (r *snmpRepo) SendAlarmTrap(deviceName, alertName, description, severity, lastedUpdateTime string) error {
	pduClass := gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.30378.2.1",
		Type:  gosnmp.OctetString,
		Value: "HPOVComponent",
	}
	pduName := gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.30378.2.2",
		Type:  gosnmp.OctetString,
		Value: deviceName,
	}
	pduAlert := gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.30378.2.3",
		Type:  gosnmp.OctetString,
		Value: alertName,
	}
	pduDesc := gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.30378.2.4",
		Type:  gosnmp.OctetString,
		Value: description,
	}
	pduSeverity := gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.30378.2.5",
		Type:  gosnmp.OctetString,
		Value: severity,
	}
	pduLastedUpdateTime := gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.30378.2.6",
		Type:  gosnmp.OctetString,
		Value: lastedUpdateTime,
	}
	trap := gosnmp.SnmpTrap{
		Enterprise:   "1.3.6.1.4.1.30378.1.1",
		AgentAddress: r.agentHost,
		GenericTrap:  6,
		SpecificTrap: 1,
		Variables:    []gosnmp.SnmpPDU{pduClass, pduName, pduAlert, pduDesc, pduSeverity, pduLastedUpdateTime},
	}

	_, err := r.client.SendTrap(trap)
	if err != nil {
		return err
	}

	return nil
}

func (r *snmpRepo) Close() {
	if r.client != nil {
		r.client.Conn.Close()
	}
}
