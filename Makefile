os = $(shell go env GOOS)
arch = $(shell go env GOARCH)
rev = $(shell git rev-parse --short HEAD)
output=build/cowsay-slackapp-$(os)-$(arch)-$(rev)

clean:
	rm -rf build/

$(output):
	go build -o $(output) github.com/nmonterroso/cowsay-slackapp/cmd/cowsay-slackapp-server

all: $(output)
