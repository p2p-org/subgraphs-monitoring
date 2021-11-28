package exporter

import (
	"context"
	"errors"
	"fmt"
	"github.com/machinebox/graphql"
	"strconv"
)

func RecordError(subgraphName string) {
	SubgraphErrors.WithLabelValues(
		subgraphName,
	).Inc()
}

func RecordMetricsTotalSubgraphsNumber(subgraphsNumber int) {
	SubgraphCount.WithLabelValues().Set(float64(subgraphsNumber))
}

func RecordMetricsSubgraph(subgraphName string, url string) error {
	client := graphql.NewClient(url)
	//https://github.com/graphprotocol/graph-node/blob/master/server/index-node/src/schema.graphql
	req := graphql.NewRequest(`
query MyQuery {
  indexingStatusForCurrentVersion(subgraphName: "connext/nxtp-matic") {
    chains {
      network
      chainHeadBlock {
        number
      }
      latestBlock {
        number
      }
    }
    health
    subgraph
    synced
  }
}
`)
	ctx := context.Background()

	var respData ResponseSubgraphStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		return err
	}

	rpcChainHead, err := strconv.ParseFloat(
		respData.IndexingStatusForCurrentVersion.Chains[0].ChainHeadBlock.Number,
		64,
	)
	if err != nil {
		return err
	}

	subgraphChainHead, err := strconv.ParseFloat(
		respData.IndexingStatusForCurrentVersion.Chains[0].LatestBlock.Number,
		64,
	)
	if err != nil {
		return err
	}

	RPCchainHead.WithLabelValues(
		subgraphName,
		respData.IndexingStatusForCurrentVersion.Subgraph,
		respData.IndexingStatusForCurrentVersion.Chains[0].Network,
	).Set(rpcChainHead)

	SubgraphChainHead.WithLabelValues(
		subgraphName,
		respData.IndexingStatusForCurrentVersion.Subgraph,
		respData.IndexingStatusForCurrentVersion.Chains[0].Network,
	).Set(subgraphChainHead)

	if respData.IndexingStatusForCurrentVersion.Synced {
		SubgraphSynced.WithLabelValues(
			subgraphName,
			respData.IndexingStatusForCurrentVersion.Subgraph,
			respData.IndexingStatusForCurrentVersion.Chains[0].Network,
		).Set(1)
	} else {
		SubgraphSynced.WithLabelValues(
			subgraphName,
			respData.IndexingStatusForCurrentVersion.Subgraph,
			respData.IndexingStatusForCurrentVersion.Chains[0].Network,
		).Set(0)
	}

	switch respData.IndexingStatusForCurrentVersion.Health {
	case "healthy":
		SubgraphHealth.WithLabelValues(
			subgraphName,
			respData.IndexingStatusForCurrentVersion.Subgraph,
			respData.IndexingStatusForCurrentVersion.Chains[0].Network,
		).Set(2)
	case "unhealthy":
		SubgraphHealth.WithLabelValues(
			subgraphName,
			respData.IndexingStatusForCurrentVersion.Subgraph,
			respData.IndexingStatusForCurrentVersion.Chains[0].Network,
		).Set(1)
	case "failed":
		SubgraphHealth.WithLabelValues(
			subgraphName,
			respData.IndexingStatusForCurrentVersion.Subgraph,
			respData.IndexingStatusForCurrentVersion.Chains[0].Network,
		).Set(2)
	default:
		errorMessage := fmt.Sprintf("Failed to get health status for subgraph: %s", subgraphName)
		return errors.New(errorMessage)
	}

	return nil
}
