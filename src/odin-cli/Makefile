GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_GET=$(GO_CMD) get
NAME=odin
DIR=./cmd/odin

all: build
build:
	$(GO_BUILD) -o $(NAME) -v $(DIR)
clean:
	$(GO_CLEAN) -v ./...
	rm -f $(NAME)
run:
	./$(NAME)
