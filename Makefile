build:
	rm -rf target/
	mkdir target/
	GOOS=windows  GOARCH=amd64 go build -o target/quick-tpl_windows.exe main.go
	GOOS=linux  GOARCH=amd64 go build -o target/quick-tpl_linux main.go
	GOOS=darwin  GOARCH=amd64 go build -o target/quick-tpl_darwin main.go
clean:
	rm -rf target/
run:
	go run main.go
