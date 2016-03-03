package commands

import (
	"bytes"
	"log"

	"github.com/ivanfoo/brickie/utils"

	"github.com/fsouza/go-dockerclient"
)

type CmdBuild struct {
	Name       string `long:"name" description:"image name" required:"true"`
	Context    string `long:"context" description:"context (folder) to use" default:"." required:"false"`
	Dockerfile string `long:"dockerfile" description:"dockerfile path" default:"Dockerfile" required:"false"`
}

func NewCmdBuild() *CmdBuild {
	return &CmdBuild{}
}

func (c *CmdBuild) Execute(args []string) error {
	client := utils.NewDockerClient()

	jsonCredentials, err := docker.NewAuthConfigurationsFromDockerCfg()
	if err != nil {
		log.Fatal(err)
	}

	err = client.BuildImage(docker.BuildImageOptions{
		Name:         c.Name,
		Dockerfile:   c.Dockerfile,
		OutputStream: bytes.NewBuffer(nil),
		ContextDir:   c.Context,
		AuthConfigs:  *jsonCredentials,
	})
	if err != nil {
		log.Fatal(err)
	}

	return err
}
