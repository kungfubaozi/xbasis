build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t xbs-safety .

run:
	docker run --net=host xbs-safety
