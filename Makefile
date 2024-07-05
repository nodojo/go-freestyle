# this is for programs, not packages or libraries
# also, if you put the @ sign before the command it won't be output to the console
build:
	@go build -o go-freestyle

run: build
	@./go-freestyle
