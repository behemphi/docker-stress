# Purpose

The `docker-stress` container is intended to for generating multiple workload 
containers on a single machine.  

# Overview

The [`stress`](http://people.seas.harvard.edu/~apw/stress/) tool is a simple 
workload generator. This project dockerizes the tool while also allowing the 
user to run _n_ number of containers from a single `docker run` command.  

# Usage

If you would like to see the help text for the `stress` tool, you can invoke 
the container thus:

```
docker run \
	--env CONTAINER_COUNT=1 \
	behemphi/stress \
		--help
```

Gotcha: The environment variable is an artifact of the inception-like nature 
of the container. Pull requests welcome.

An example of using the tool to generate 10 seconds of stress with a single 
container:

```
docker run \
	--detach \
	--env CONTAINER_COUNT=1 \
	behemphi/stress \
		--cpu 1 --io 1 --vm 2 --vm-bytes 16M --timeout 10s
```

Notice that the CLI of `stress` is preserved exactly.  

An example of using the tool to generate 100 seconds of stress 
with 6 containers:

```
docker run \
	--detach \
	--env CONTAINER_COUNT=6 \
	--privileged \
	--volume /usr/local/bin/docker:/docker \
	--volume /var/run/docker.sock:/var/run/docker.sock \
	behemphi/stress \
		--cpu 1 --io 1 --vm 2 --vm-bytes 8M --timeout 100s
```

In this case each of the 6 containers will run for 100 seconds with the 
settings passed in.

# Warnings

There is a reason this container is called `stress`. Know what you are using 
by reading up on the purpose of the tool.  It is not recommended to use this
in production or other high-value environments.  

# Use Cases

The motivation for this was an evaluation of Google's 
[`cAdvisor`](https://github.com/google/cadvisor). The goal was to generate 
sustained, varied  activity on a docker node and see what cAdvisor had to 
say about it.

Other potential use cases are:
* Testing scheduling algorithms
* Stressing software applications in test environments by providing noisy
  neighbors

# License 

MIT 