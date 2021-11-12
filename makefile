VERSION := "2.0.0"

BINARY_NAME=./calendar-api
SERVICE_NAME=calendar-api.service

PLATFORMS := linux/amd64/x86_64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))
name = $(word 3, $(temp))

tidy:
	go mod tidy

test:
	go test ./...

check:
	go fmt ./
	go vet ./

build: check clean $(PLATFORMS)

$(PLATFORMS):
	env GOOS=$(os) GOARCH=$(arch) go build -trimpath -ldflags="-s -w -X 'main.Version=${VERSION}'" -o $(BINARY_NAME) .

clean:
	go clean

restart:
	sudo /bin/systemctl restart $(SERVICE_NAME)

.PHONY: build $(PLATFORMS)
.DEFAULT_GOAL := build
