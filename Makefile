base_dir=.
main_file=main.go
profile_name_main=website
compiler=go build -v -o $(profile_name_main) $(main_file)

.PHONY: build main cron tool lint clean help

all: build

build:
	$(compiler)
install:
	go install 
tool:
	go vet $(base_dir)/...; true
	gofmt -w $(base_dir)

lint: 
	golint ./...

clean:
	rm -rf $(profile_name_main)
	go clean -i $(base_dir)

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"




