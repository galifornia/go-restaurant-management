PID = ./go-restaurant-management.pid
GO_FILES = $(wildcard *.go)
APP = ./go-restaurant-management

serve: build start
	@fswatch -x -o --event Created --event Updated --event Renamed -r -e '.*' -i '\.go$$'  . | xargs -n1 -I{}  make restart || make kill

kill:
	@kill `cat $(PID)` || true

build: $(GO_FILES)
	@go build -o $(APP)

$(APP): $(GO_FILES)
	@go build $? -o $@

start:
	# @sh -c "$(APP) & echo $$! > $(PID)"
	@$(APP) & echo $$! > $(PID)

restart: kill build start
