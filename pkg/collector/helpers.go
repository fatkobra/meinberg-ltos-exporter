package collector

import "github.com/prometheus/client_golang/prometheus"

type typedDesc struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
}

func (td typedDesc) mustNewConstMetric(value float64, labels ...string) prometheus.Metric {
	return prometheus.MustNewConstMetric(td.desc, td.valueType, value, labels...)
}

func boolToFloat64(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}
