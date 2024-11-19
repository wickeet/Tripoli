#variables
PROJECT_DIR := $(CURDIR)

# make test
test:
	@echo "Realizando tests..."
	@go test -v $(PROJECT_DIR)/...

# make check
check:
	@echo "Comprobando sintaxis..."
	@if gofmt -e $(PROJECT_DIR)/internal/ > /dev/null; then \
		echo "No hay errores de sintaxis"; \
	else \
		gofmt -e $(PROJECT_DIR)/internal/ > /dev/null; \
	fi

