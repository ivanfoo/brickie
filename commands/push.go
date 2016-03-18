package commands

import (
	"bytes"
	"log"

	"github.com/ivanfoo/brickie/utils"

	"github.com/fsouza/go-dockerclient"
)

type CmdPush struct {
	Image       string `long:"image" description:"name of the image" required:"true"`
	Registry    string `long:"registry" description:"registry to push"`
	credentials *docker.AuthConfiguration
}

func NewCmdPush() *CmdPush {
	return &CmdPush{}
}

func (c *CmdPush) Execute(args []string) error {
	client := utils.NewDockerClient()

	if c.Registry != "" {
		c.credentials = utils.GetRegistryCredentialsFromFile(c.Registry)
	} else {
		c.credentials = &docker.AuthConfiguration{}
	}

	err := client.PushImage(docker.PushImageOptions{
		Name:         c.Image,
		OutputStream: bytes.NewBuffer(nil),
	}, *c.credentials)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
