package exporter

type ResponseSubgraphStruct struct {
	IndexingStatusForCurrentVersion struct {
		Chains []struct {
			ChainHeadBlock struct {
				Number string `json:"number"`
			} `json:"chainHeadBlock"`
			LatestBlock struct {
				Number string `json:"number"`
			} `json:"latestBlock"`
			Network string `json:"network"`
		} `json:"chains"`
		Health   string `json:"health"`
		Subgraph string `json:"subgraph"`
		Synced   bool   `json:"synced"`
	} `json:"indexingStatusForCurrentVersion"`
}
