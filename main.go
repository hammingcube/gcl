package main

import (
	"log"
        _ "fmt"
        "github.com/fsouza/go-dockerclient"
)

func test(client *docker.Client) {
	log.Println("Removing container")
	ContainerId := "abc"
	if err := client.RemoveContainer(docker.RemoveContainerOptions{}); err != nil {
		log.Printf("Couldn't remove container %s (%v)\n", ContainerId, err)
	}
}

func main() {
        endpoint := "unix:///var/run/docker.sock"
        client, _ := docker.NewClient(endpoint)
	test(client)
}
