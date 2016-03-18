package commands

import (
	"log"

	"github.com/ivanfoo/brickie/utils"
)

type CmdLogin struct {
	Username string `long:"username" description:"registry username" required:"true"`
	Password string `long:"password" description:"registry password" required:"true"`
	Email    string `long:"email" <F6><F5>description:"registry email" default:"."`
	Registry string `long:"registry" description:"docker registry" required:"true"`
}

func NewCmdLogin() *CmdLogin {
	return &CmdLogin{}
}

func (c *CmdLogin) Execute(args []string) error {

	err := utils.GenerateCredentialsFile(c.Username, c.Password, c.Email, c.Registry)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
