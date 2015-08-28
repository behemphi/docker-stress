package main

import (
	//"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// Rather than reimplementing the `stress` CLI set the number of stress
// containers to run as an environment variable
//
// `stress` parameters will be passed to the container run time config as
// parameters of the image. The goal being to make the container run look,
// more or less, like any CLI call of `stress` itself
func main() {

	// The number of stress containers that will be run
	container_count := 4

	// The expected path of the docker daemon binary in the container
	docker_bin_path := "/docker"

	// The expect path of the docker socket file in the container
	docker_socket_path := "/var/run/docker.sock"

	// Set the container count to the user defined value if it exists
	if len(os.Getenv("CONTAINER_COUNT")) != 0 {
		i, err := strconv.Atoi(os.Getenv("CONTAINER_COUNT"))
		if err != nil {
			log.Fatal(err)
		}
		container_count = i
	}

	log.Printf("Begin launching %d `stress` containers", container_count)

	// Set the user defined location of the Docker binary in the container
	if len(os.Getenv("DOCKER_BIN_PATH")) != 0 {
		docker_bin_path = os.Getenv("DOCKER_BIN_PATH")
	}

	log.Printf("Local path to Docker binary is: %s\n", docker_bin_path)

	// Set the user defined location of the Docker daemon socket file
	if len(os.Getenv("DOCKER_BIN_PATH")) != 0 {
		docker_socket_path = os.Getenv("DOCKER_SOCKET_PATH")
	}

	log.Printf("Local path to Docker socket is: %s\n", docker_socket_path)

	args := os.Args
	args = append(args[:0], args[1:]...)

	// If the container count is just one, then we assume this is the only
	// container to be started and we call stress directly.
	// Otherwise, we keep iterating spawning containers until we reached
	// the specified number.
	// This means we can use a single image.
	if container_count == 1 {

		cmd := exec.Command("stress", args...)
		_, err := cmd.Output()
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf("Stress command to be executed executed: %s \n", cmd)

	} else {
		for i := 0; i < container_count+1; i++ {
			cmd := []string{"run", "--detach", "--env", "CONTAINER_COUNT=1",
				"behemphi/stress"}
			args = append(cmd, args...)

			// Because the run time config of the contianer will mount the
			// docker binary and socket on the host, we are talking to
			// the host docker container.
			docker_cmd := exec.Command("/docker", args...)

			log.Printf("Docker command to spown containes is: %s \n", docker_cmd)

			out, err := docker_cmd.Output()
			if err != nil {
				log.Printf(err.Error())
			}

			log.Printf("Output from docker run command: %s", out)
		}
	}
}
