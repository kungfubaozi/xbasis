build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t gosion-permission .

run:
	docker run -P gosion-permission
