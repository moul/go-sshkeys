.PHONY: test
test:
	go test -v ./...

.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	goconvey -port 7082 -cover -workDir=$(realpath .) -depth=0
