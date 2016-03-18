package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fsouza/go-dockerclient"
)

var credentialsFile string = fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".brickie.auth")

type Credentials struct {
	Auths map[string]Auth `json:"auths"`
}

type Auth struct {
	Auth  string `json:"auth"`
	Email string `json:"email"`
}

func LoadCredentialsFile() *docker.AuthConfigurations {
	content, _ := os.Open(credentialsFile)
	jsonCredentials, err := docker.NewAuthConfigurations(content)
	if err != nil {
		log.Fatal(err)
	}

	return jsonCredentials
}

func GetRegistryCredentialsFromFile(registryName string) *docker.AuthConfiguration {
	jsonCredentials := LoadCredentialsFile()

	credentials := &docker.AuthConfiguration{
		Username:      jsonCredentials.Configs[registryName].Username,
		Email:         jsonCredentials.Configs[registryName].Email,
		Password:      jsonCredentials.Configs[registryName].Password,
		ServerAddress: jsonCredentials.Configs[registryName].ServerAddress,
	}

	return credentials
}

func GenerateCredentialsFile(username string, password string, email string, registry string) error {
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	c := &Credentials{
		Auths: map[string]Auth{
			registry: {
				Auth:  encodedCredentials,
				Email: email,
			},
		},
	}

	jsonView, _ := json.Marshal(c)
	err := ioutil.WriteFile(credentialsFile, jsonView, 0600)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
