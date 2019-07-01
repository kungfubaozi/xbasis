build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t xbs-permission .

run:
	docker run --net=host xbs-permission
