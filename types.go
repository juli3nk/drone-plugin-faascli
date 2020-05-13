package main

type Settings struct {
	YamlFile         string `split_words:"true"`
	Templates        []string
	Build            bool
	BuildArgs        []string `split_words:"true"`
	BuildArgsFromEnv []string `split_words:"true"`
	Push             bool
	DockerDns        []string `split_words:"true"`
	Registry         string
	RegistryUsername string `split_words:"true"`
	RegistryPassword string `split_words:"true"`
	RegistryAuth     string `split_words:"true"`
	RegistryInsecure bool   `split_words:"true"`
	Debug            bool
}

type Plugin struct {
	Docker   *DockerDaemon
	Registry *Registry
	FaasCLI  *FaasCLI
	Debug    bool
}

type DockerDaemon struct {
	DNS []string
}

type Registry struct {
	Host     string
	Username string
	Password string
	Auth     string
	Insecure bool
}

type FaasCLI struct {
	General General
	Build   Build
	Push    Push
}

type General struct {
	YamlFile string
}

type Build struct {
	Enable    bool
	Templates []string
	Args      []string
}

type Push struct {
	Enable bool
}
