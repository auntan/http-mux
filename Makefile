APP=http-mux

build: clean
	go build -o bin/${APP} cmd/${APP}/main.go

clean:
	@rm -f bin/${APP}
