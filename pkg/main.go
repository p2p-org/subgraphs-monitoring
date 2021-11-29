package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"subgraphmon/config"
	"subgraphmon/exporter"
	"time"
)

func main() {
	var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
	var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
	var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

	viper.AutomaticEnv()
	viper.SetDefault("CONFIG", "./config.yml")

	Config := viper.GetString("CONFIG")
	var configuration config.Configuration

	viper.SetConfigFile(Config)

	if err := viper.ReadInConfig(); err != nil {
		Error.Printf("Unable to read config file, %s", err)
		os.Exit(1)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		Error.Printf("Unable to decode into struct, %v", err)
		os.Exit(1)
	}

	Info.Printf("Ready to go!")
	exporter.RecordMetricsTotalSubgraphsNumber(len(configuration.Subgraphs))

	// Only one gorutine for all subgraph entities
	go func() {
		for {
			for _, subgraph := range configuration.Subgraphs {
				err := exporter.RecordMetricsSubgraph(subgraph.Name, subgraph.URL)
				if err != nil {
					Warning.Printf("Unable to get subgraph metrics, %v", err)
					exporter.RecordError(subgraph.Name)
				} else {
					Info.Printf("Update metrics for subgraph")
				}
			}
			time.Sleep(time.Duration(configuration.Interval) * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

}
