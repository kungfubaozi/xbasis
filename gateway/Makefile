build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t xbs-gateway:0.1 .

run:
	docker run -p 9090:9090 xbs-gateway
