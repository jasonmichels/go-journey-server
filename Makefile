build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-journey-server .
test:
	go test -v ./... -bench . -cover