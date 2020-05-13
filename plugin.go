package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/docker/docker/client"
)

func (p *Plugin) startDockerDaemon() {
	cmd := commandDockerDaemon(p.Docker, p.Registry)
	if p.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = ioutil.Discard
		cmd.Stderr = ioutil.Discard
	}

	go func() {
		trace(cmd)
		cmd.Run()
	}()

	for i := 0; i < 15; i++ {
		cli, err := client.NewEnvClient()
		if err == nil {
			cli.Close()
			break
		}

		time.Sleep(time.Second * 1)
	}
}

func (p *Plugin) registryLogin() error {
	method := "none"

	// login to the Docker registry
	if p.Registry.Username != "" && p.Registry.Password != "" {
		cmd := commandDockerLogin(p.Registry)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Error authenticating: %s", err)
		}

		method = "credentials"
	}

	// Create Auth Config File
	if p.Registry.Auth != "" {
		os.MkdirAll(dockerHome, 0600)
		path := filepath.Join(dockerHome, "config.json")

		data, err := generateDockerConfigJson(p.Registry)
		if err != nil {
			return fmt.Errorf("Error generating config.json: %s", err)
		}

		if err := ioutil.WriteFile(path, data, 0600); err != nil {
			return fmt.Errorf("Error writing config.json: %s", err)
		}

		method = "auth"
	}

	switch method {
	case "credentials":
		fmt.Println("Detected registry credentials")
	case "auth":
		fmt.Println("Detected registry credentials file")
	default:
		fmt.Println("Registry credentials or Docker config not provided. Guest mode enabled.")
	}

	return nil
}

func (p Plugin) exec() error {
	var cmds []*exec.Cmd
	cmds = append(cmds, commandFaasCliVersion())

	// Build stack
	if p.FaasCLI.Build.Enable {
		p.startDockerDaemon()

		cmds = append(cmds, commandDockerVersion())
		cmds = append(cmds, commandDockerInfo())

		for _, tpl := range p.FaasCLI.Build.Templates {
			cmd, err := commandFaasCliTemplatePull(tpl)
			if err != nil {
				return err
			}
			cmds = append(cmds, cmd)
		}

		cmds = append(cmds, commandFaasCliBuild(p.FaasCLI.General, p.FaasCLI.Build))
	}

	// Push stack
	if p.FaasCLI.Build.Enable && p.FaasCLI.Push.Enable {
		if err := p.registryLogin(); err != nil {
			return err
		}

		cmds = append(cmds, commandFaasCliPush(p.FaasCLI.General, p.FaasCLI.Push))
	}

	// Execute all commands in batch mode.
	for _, cmd := range cmds {
		trace(cmd)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
