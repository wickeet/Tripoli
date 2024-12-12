.PHONY: check test

check:
	echo "Comprobando sintaxis..."
	gofmt -e .

test:
	echo "Realizando tests..."
	go test ./...