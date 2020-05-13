package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

func main() {
	var s Settings

	if err := envconfig.Process("plugin", &s); err != nil {
		log.Fatal(err.Error())
	}

	dd := DockerDaemon{
		DNS: s.DockerDns,
	}

	registry := Registry{
		Host:     s.Registry,
		Username: s.RegistryUsername,
		Password: s.RegistryPassword,
		Auth:     s.RegistryAuth,
		Insecure: s.RegistryInsecure,
	}

	general := General{
		YamlFile: s.YamlFile,
	}
	fBuild := Build{
		Enable:    s.Build,
		Templates: s.Templates,
		Args:      s.BuildArgs,
	}
	for _, arg := range s.BuildArgsFromEnv {
		a := getArgFromEnv(arg)
		if a == "" {
			log.Fatalf("Arg \"%s\" from env does not exist", arg)
		}
		fBuild.Args = append(fBuild.Args, a)
	}
	fPush := Push{
		Enable: s.Push,
	}
	faascli := FaasCLI{
		General: general,
		Build:   fBuild,
		Push:    fPush,
	}

	plugin := &Plugin{
		Docker:   &dd,
		Registry: &registry,
		FaasCLI:  &faascli,
		Debug:    s.Debug,
	}

	if err := plugin.exec(); err != nil {
		log.Fatal(err)
	}
}
