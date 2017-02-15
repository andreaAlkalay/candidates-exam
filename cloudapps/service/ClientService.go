package service

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

func GetClient() *cfclient.Client {
	config := &cfclient.Config{
		ApiAddress:   "https://api.51.141.3.112.xip.io",
		Username:     "andreaa",
		Password:     "andreaa",
		SkipSslValidation: true,
	}

	var	client *cfclient.Client
	client, _ = cfclient.NewClient(config)

	return client
}