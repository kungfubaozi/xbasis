build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t gosion-connection .

run:
	docker run -P -p 9000:9000 gosion-connection
