package servicemonitors

import (
	"github.com/grafana/agent/component"
	"github.com/grafana/agent/component/prometheus/operator"
	"github.com/grafana/agent/component/prometheus/operator/common"
	"github.com/grafana/agent/service/cluster"
)

func init() {
	component.Register(component.Registration{
		Name:          "prometheus.operator.servicemonitors",
		Args:          operator.Arguments{},
		NeedsServices: []string{cluster.ServiceName},

		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			return common.New(opts, args, common.KindServiceMonitor)
		},
	})
}
