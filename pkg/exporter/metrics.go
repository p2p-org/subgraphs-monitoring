package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RPCchainHead = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "subgraphs_monitoring_rpc_chain_head",
		Help: "Current rpc node highest block",
	},
		[]string{
			"subgraph",
			"hash",
			"network",
		},
	)
	SubgraphChainHead = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "subgraphs_monitoring_subgraph_chain_head",
		Help: "Current subgraph's highest block",
	},
		[]string{
			"subgraph",
			"hash",
			"network",
		},
	)
	SubgraphErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "subgraphs_monitoring_subgraph_error_count",
		Help: "Error count when getting metrics for subgraph",
	},
		[]string{
			"subgraph",
		},
	)
	SubgraphCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "subgraphs_monitoring_subgraphs_total_count",
		Help: "Subgraphs total count",
	},
		[]string{},
	)
	SubgraphHealth = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "subgraphs_monitoring_subgraph_health",
		Help: "Subgraph health: 0 for failed, 1 for unhealthy, 2 for healhty",
	},
		[]string{
			"subgraph",
			"hash",
			"network",
		},
	)
	SubgraphSynced = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "subgraphs_monitoring_subgraph_synced",
		Help: "Subgraph synced: 0 for unsynced, 1 for synced",
	},
		[]string{
			"subgraph",
			"hash",
			"network",
		},
	)
)
