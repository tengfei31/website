base_dir=.
#main
main_file=main.go
profile_name_main=website
#cron
cron_file=cron.go
profile_name_cron=website_cron
.PHONY: build main cron tool lint clean help

all: build

build: main cron

main: go build -v -o $(profile_name_main) $(main_file)

cron: go build -v -o $(profile_name_cron) $(cron_file)

tool:
	go vet $(base_dir)/...; true
	gofmt -w $(base_dir)

lint: golint ./...

clean:
	rm -rf $(profile_name_main) $(profile_name_cron)
	go clean -i $(base_dir)

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"




