build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-journey-server .
build-docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-journey-server .
	docker build -t jasonmichels/go-journey-server:develop .
test:
	go test -v ./... -bench . -cover