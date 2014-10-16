package main

import (
	"log"
        _ "fmt"
        "github.com/fsouza/go-dockerclient"
)

func add(client *docker.Client) (error, string) {
	config := &docker.Config{
		Image: "ubuntu:14.04",
	}
	container, err := client.CreateContainer(docker.CreateContainerOptions{Config : config})
	if err != nil {
		return err, ""
	}
	return nil, container.ID
	
}

func test(client *docker.Client) {
	err, cid := add(client)
	if err != nil {
		log.Println("Something went wrong in creating container")
		return
	}
	
	log.Println("Created container %s", cid)
	log.Println("Removing container %s", cid)
	if err := client.RemoveContainer(docker.RemoveContainerOptions{ID:cid,}); err != nil {
		log.Printf("Couldn't remove container %s (%v)\n", cid, err)
	}
}

func main() {
        endpoint := "unix:///var/run/docker.sock"
        client, _ := docker.NewClient(endpoint)
	test(client)
}
