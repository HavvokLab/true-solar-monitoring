package repo

import "github.com/HavvokLab/true-solar-monitoring/util"

type mockSnmpRepo struct{}

func NewMockSnmpRepo() SnmpRepo {
	return &mockSnmpRepo{}
}

func (m *mockSnmpRepo) SendAlarmTrap(deviceName, alertName, description, severity, lastedUpdateTime string) error {
	data := map[string]interface{}{
		"deviceName":       deviceName,
		"alertName":        alertName,
		"description":      description,
		"severity":         severity,
		"lastedUpdateTime": lastedUpdateTime,
	}

	util.PrintJSON(data)
	return nil
}

func (m *mockSnmpRepo) Close() {}
