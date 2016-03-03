package utils

import (
	"log"

	"github.com/fsouza/go-dockerclient"
)

func GetRegistryCredentialsFromDockerCfg(registryName string) *docker.AuthConfiguration {
	jsonCredentials, err := docker.NewAuthConfigurationsFromDockerCfg()
	if err != nil {
		log.Fatal(err)
	}

	credentials := &docker.AuthConfiguration{
		Username:      jsonCredentials.Configs[registryName].Username,
		Email:         jsonCredentials.Configs[registryName].Email,
		Password:      jsonCredentials.Configs[registryName].Password,
		ServerAddress: jsonCredentials.Configs[registryName].ServerAddress,
	}

	return credentials
}
