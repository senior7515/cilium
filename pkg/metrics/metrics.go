package metrics

// This file holds all our metrics. They are exported to the rest of the daemon
//
// **Adding a metric**
// - Add a metric object of the appropriate type that is exported
// - Register the new object in the init function

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// This is the namespace presented to prometheus and should be scoped to us
	Namespace = "cilium"

	// Endpoints
	NumEndpoints = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "endpoints",
		Help:      "Number of endpoints managed by this agent",
		// FIXME: do we have any node IDs for the agent? or agent IDs? Do we even have to provide anything?
		// ConstLabels: prometheus.Labels{"node": "a node ID, from k8s?",
	})

	NumEndpointsRegenerating = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "endpoints_regenerating",
		Help:      "Number of endpoints currently regenerating",
	})

	CountEndpointsRegenerations = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "endpoints_regenerations",
		Help:      "Count of all endpoint regenerations that have completed, tagged by outcome",
	},
		[]string{"outcome"})

	// Policies
	NumPolicies = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "policies",
		Help:      "Number of policies currently loaded",
	})

	PolicyRevision = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "policies_max_revision",
		Help:      "Highest policy revision number in the agent",
	})

	PolicyImportErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "policies_import_errors",
		Help:      "Number of times a policy import has failed",
	})
)

func init() {
	prometheus.MustRegister(NumEndpoints)
	prometheus.MustRegister(NumEndpointsRegenerating)
	prometheus.MustRegister(CountEndpointsRegenerations)

	prometheus.MustRegister(NumPolicies)
	prometheus.MustRegister(PolicyRevision)
	prometheus.MustRegister(PolicyImportErrors)
}
