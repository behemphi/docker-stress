linux:
	CGO_ENABLED=0 GOOS=linux go build -o buildoutput/docker-stress
	docker build -t behemphi/stress .

osx:
	go build -o buildoutput/docker-stress

