build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t gs-permission .

run:
	docker run --net=host gs-permission
