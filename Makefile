linux:
	CGO_ENABLED=0 GOOS=linux go build
	docker build -t behemphi/stress .

osx:
	go build

