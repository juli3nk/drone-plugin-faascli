package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}

func getArgFromEnv(key string) string {
	value := os.Getenv(key)

	return fmt.Sprintf("%s=%s", key, value)
}
