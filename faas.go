package main

import (
	"net/url"
	"os/exec"
)

const faasCliExe = "/usr/local/bin/faas-cli"

func commandFaasCliVersion() *exec.Cmd {
	return exec.Command(faasCliExe, "version")
}

func commandFaasCliTemplatePull(template string) (*exec.Cmd, error) {
	args := []string{"template"}

	u, err := url.Parse(template)
	if err != nil {
		return nil, err
	}

	if u.Scheme == "" && u.Host == "" {
		args = append(args, "store")
	}

	args = append(args, "pull", template)

	return exec.Command(faasCliExe, args...), nil
}

func commandFaasCliBuild(general General, build Build) *exec.Cmd {
	args := []string{"build"}

	args = append(args, "-f", general.YamlFile)

	for _, arg := range build.Args {
		args = append(args, "--build-arg", arg)
	}

	return exec.Command(faasCliExe, args...)
}

func commandFaasCliPush(general General, push Push) *exec.Cmd {
	args := []string{"push"}

	args = append(args, "-f", general.YamlFile)

	return exec.Command(faasCliExe, args...)
}
