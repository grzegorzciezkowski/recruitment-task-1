build:
	mkdir -p "target"
	go build -o "target/server" "cmd/server/main.go"
	cp "exampleconfiguration/config.yaml" "target/config.yaml"
	cp "exampleconfiguration/input.txt" "target/input.yaml"

test:
	go test ./..

clean:
	rm -rfv "target"
