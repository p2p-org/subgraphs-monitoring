# Subgraph monitoring exporter

Prometheus exporter which provides metrics for monitoring multi subgraphs and rpc nodes by graphql request to graph-node endpoint like [here](https://thegraph.com/docs/hostedservice/deploy-subgraph-hosted#checking-subgraph-health)

Script sends request to `8020(by default)` graph-node port.

To set config name use env variable `CONFIG`

Example config is located [here](pkg/config.yml)

List of available metrics:

```
subgraphs_monitoring_rpc_chain_head        - Current rpc node highest block
subgraphs_monitoring_subgraph_chain_head   - Current subgraph's highest block
subgraphs_monitoring_subgraph_error_count  - Error count when getting metrics for subgraph
subgraphs_monitoring_subgraphs_total_count - Subgraphs total count
subgraphs_monitoring_subgraph_health       - Subgraph health: 0 for failed, 1 for unhealthy, 2 for healhty
subgraphs_monitoring_subgraph_synced       - Subgraph synced: 0 for unsynced, 1 for synced
```

Prometheus alerts example is located [here](./alerts-example.yml)


# Guide for deploying connext subgrphs

## Before you start 
If you are using hetzner, gcp, aws you need to use server profile for you ipfs node. More info [here](https://github.com/ipfs/go-ipfs/issues/4343)

Also you can save a lot of disk space by disable block_hash for graph-node:

```
GRAPH_ETHEREUM_CLEANUP_BLOCKS: 'true'
```

It can lead to increasing number of total rpc requests.

## How to deploy connext subgraph

- Install all needed packages:
  ```
  apt-get update
  apt-get install -y docker.io jq docker-compose npm
  npm install --global yarn
  ```

- Copy graph repo:

  ```
  git clone https://github.com/connext/nxtp.git
  ```

- Change provider url in docker-compose :

  ```
  cd ~/graph-node/docker
  -      ethereum: 'mainnet:http://host.docker.internal:8545'
  +      ethereum: 'network(for example-matic):your_rpc_node_url'
  ```

- Run `./setup.sh`
- Run docker-compose

  ```
  docker-compose up -d
  ```

- Clone connext repo

  ```
  git clone https://github.com/connext/nxtp.git
  cd nxtp/packages/subgraph
  ```

Be sure that you are using graph version 0.21.1 (Because current subgraphs use api version 0.0.4)

- Compile subgraph(config are stored in ./configs directory)

  ```
  yarn deploy v1-runtime v1-runtime matic
  ```

  This will compile subgraph for matic(this shoud be in config ./configs/mainnet.json and in graph docker-compose you should have provider for that network)

- Create subgraph and deploy it(make sure that graph docker-compose is up and running)

  ```
  graph create --node http://localhost:8020/ connext/nxtp-your-subgraph-name
  graph deploy --node http://localhost:8020/ --ipfs http://localhost:5001 connext/your-subgraph-name
  ```

- Check subgraph

  ```
  curl http://127.0.0.1:8000/subgraphs/name/connext/nxtp-your-subgraph-name
  ```
