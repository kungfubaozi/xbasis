build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t gs-user .

run:
	docker run --net=host gs-user
