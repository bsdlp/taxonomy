package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/bsdlp/taxonomy/internal/clients"
	"github.com/bsdlp/taxonomy/internal/configuration"
	"github.com/bsdlp/taxonomy/rpc/taxonomy"
	"github.com/bsdlp/taxonomy/server"
)

func main() {
	// get config file path env var
	cfgFilePath := os.Getenv("TAXONOMY_CONFIG_FILE")
	if cfgFilePath == "" {
		cfgFilePath = "/etc/taxonomy/config"
	}

	// consume config file
	var cfg configuration.Object
	if configFh, err := os.Open(cfgFilePath); err == nil {
		err = json.NewDecoder(configFh).Decode(&cfg)
		if err != nil {
			log.Fatalf("error decoding config: %s", err.Error())
		}
		configFh.Close()
	} else {
		log.Fatalf("error reading config file: %s", err.Error())
	}

	err := cfg.SetSecretsFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	clients, err := clients.New(cfg)
	if err != nil {
		log.Fatalf("error initiating clients: %s", err)
	}

	srv := &server.Server{
		Clients: clients,
	}

	err = http.ListenAndServe(cfg.ServerHostPort, taxonomy.NewTaxonomyServer(srv))
	if err != nil {
		log.Fatal(err.Error())
	}
}
