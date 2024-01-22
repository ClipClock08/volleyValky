PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin

MOQ = $(PROJECT_BIN)/moq
MOQ_VERSION = v0.3.1

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint
GOLANGCI_LINT_VERSION = v1.52.0

MIG_DIR := internal/storage/migrations


# === Mocks generator ===

.PHONY: .install-moq
.install-moq:
	@echo "Installing moq..."
	@mkdir -p $(PROJECT_BIN)
	[ -f $(MOQ) ] || GOBIN=$(PROJECT_BIN) go install github.com/matryer/moq@$(MOQ_VERSION)


# === Linter ===
.PHONY: .install-linter
.install-linter:
	### INSTALL GOLANGCI-LINT ###
	[ -f $(GOLANGCI_LINT) ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) $(GOLANCI_LINT_VERSION)

.PHONY: lint
lint: .install-linter
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml

.PHONY: lint-fast
lint-fast: .install-linter
	$(GOLANGCI_LINT) run ./... --fast --config=./.golangci.yml


# === Install environment ===
.PHONY: install-env
install-env: .install-moq .install-linter


# === Tests ===
.PHONY: test
test:
	go test ./...

# === Migrations ===
.PHONY: goose-status

goose-status:
	@cd $(MIG_DIR) && goose postgres "user=postgres dbname=volley_valky_bot password=postgres sslmode=disable" status

.PHONY: goose-up

goose-up:
	@cd $(MIG_DIR) && goose postgres "user=postgres dbname=volley_valky_bot password=postgres sslmode=disable" up

.PHONY: goose-down

goose-down:
	@cd $(MIG_DIR) && goose postgres "user=postgres dbname=volley_valky_bot password=postgres sslmode=disable" down