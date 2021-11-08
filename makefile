BINARY_PATH=./bin
BINARY_NAME=$(BINARY_PATH)/calendar-api
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
	env GOOS=$(os) GOARCH=$(arch) go build -trimpath -ldflags="-s -w" -o $(BINARY_NAME) .

clean:
	go clean
	rm -f $(BINARY_NAME)

restart:
	sudo /bin/systemctl restart $(SERVICE_NAME)

.PHONY: build $(PLATFORMS)
.DEFAULT_GOAL := build
