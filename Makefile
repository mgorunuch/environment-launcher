generate_bindata:
	go-bindata -o ./bindata/bindata.go -pkg bindata ./static/...

test:
	go test ./...

run: generate_bindata
	go run .