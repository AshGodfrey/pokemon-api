.PHONY: test mock-test e2e-test run-main run-cli

# Default command when "make" is run without arguments
test: mock-test e2e-test

# Run mocked tests
mock-test:
	@echo "Running mocked tests..."
	@go test ./...

# Run end to end tests
e2e-test:
	@echo "Running end-to-end tests..."
	@go test -tags=e2e ./...

# Run main.go 
	@echo "Running main.go..."
	@go run main.go


# Pattern rule for running the CLI tool with a Pokémon name.
# Usage: make get-pokemon-<pokemon>
get-pokemon-%:
	@echo "Running CLI tool for $*..."
	@./pokecli pokemon $*

# Pattern rule for running the CLI tool with a Pokémon name & location.
# Usage: make get-location--<pokemon>
get-location-%:
	@echo "Running CLI tool for $*..."
	@./pokecli location $*

