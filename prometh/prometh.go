package prometh

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"log"
)

func Push(pushAddr string, jobName string, res int64) {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{Name: jobName})
	gauge.Set(float64(res))
	err := push.New(pushAddr, jobName).Grouping("module", "aleoscan").Collector(gauge).Push()
	if err != nil {
		log.Printf("push prometheus %s failed:%s", pushAddr, err)
	}

}
