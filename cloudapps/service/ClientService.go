package service

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

func GetClient() *cfclient.Client {
	config := &cfclient.Config{
		ApiAddress:   "https://api.<cloud-foundry-ip>.xip.io",
		Username:     "",
		Password:     "",
		SkipSslValidation: true,
	}

	var	client *cfclient.Client
	client, _ = cfclient.NewClient(config)

	return client
}
