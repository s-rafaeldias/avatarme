test:
	go test ./... -cover

test_ci:
	go test ./... -cover -coverprofile=/tmp/artifacts/c.txt -covermode=atomic 

go_get:
	go get ./...
