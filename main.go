package main

import (
        _ "fmt"
        "github.com/fsouza/go-dockerclient"
)

func main() {
        endpoint := "unix:///var/run/docker.sock"
        client, _ := docker.NewClient(endpoint)
	c, err := client.CreateContainer()
}
