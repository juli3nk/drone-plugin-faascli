package main

import (
	"os/exec"
)

const dockerExe = "/usr/local/bin/docker"
const dockerdExe = "/usr/local/bin/dockerd"
const dockerHome = "/root/.docker/"

func commandDockerDaemon(daemon *DockerDaemon, registry *Registry) *exec.Cmd {
	args := []string{}

	if registry.Host != "" && registry.Insecure {
		args = append(args, "--insecure-registry", registry.Host)
	}
	for _, dns := range daemon.DNS {
		args = append(args, "--dns", dns)
	}

	return exec.Command(dockerdExe, args...)
}

func commandDockerVersion() *exec.Cmd {
	return exec.Command(dockerExe, "version")
}

func commandDockerInfo() *exec.Cmd {
	return exec.Command(dockerExe, "info")
}

func commandDockerLogin(reg *Registry) *exec.Cmd {
	return exec.Command(
		dockerExe, "login",
		"-u", reg.Username,
		"-p", reg.Password,
		reg.Host,
	)
}
