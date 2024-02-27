t="/tmp/go-cover.test.tmp"

test:
	go test ./...

coverage:
	go test -coverprofile=$t ./... && go tool cover -html=$t && unlink $t

