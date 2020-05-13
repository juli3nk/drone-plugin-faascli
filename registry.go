package main

import (
	"encoding/json"
)

func generateDockerConfigJson(reg *Registry) ([]byte, error) {
	type Auth struct {
		Auth string `json: "auth"`
	}
	type Config struct {
		Auths map[string]Auth `json: "auths"`
	}

	host := "https://index.docker.io/v1/"
	if reg.Host != "" {
		host = reg.Host
	}

	a := Auth{Auth: reg.Auth}
	auths := make(map[string]Auth)
	auths[host] = a

	c := Config{Auths: auths}

	d, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return d, nil
}
