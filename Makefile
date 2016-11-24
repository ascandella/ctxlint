PROJECT_ROOT := github.com/sectioneight/ctxlint

LINT_LOG := lint.log
COV_REPORT := overalls.coverprofile

PKG_FILES = *.go ctxlint

.PHONY: lint
lint:
	gofmt -d -s $(PKG_FILES) | tee -a $(LINT_LOG)
	$(foreach dir,$(PKG_FILES),go tool vet $(VET_RULES) $(dir) 2>&1 | tee -a $(LINT_LOG);)
	@[ ! -s $(LINT_LOG) ]

.PHONY: test
test: $(COV_REPORT)

$(COV_REPORT): $(PKG_FILES)
	overalls -project=$(PROJECT_ROOT) -- -race -v | \
		grep -v "No Go Test Files"

.PHONY: coveralls
coveralls: $(COV_REPORT)
	@goveralls -coverprofile=$(COV_REPORT)

.PHONY: dependencies
dependencies:
	@which overalls || go get -u github.com/go-playground/overalls
	@which goveralls || go get -u -f github.com/mattn/goveralls
	@go get github.com/stretchr/testify/assert
	@go get github.com/stretchr/testify/require

.PHONY: clean
clean:
	rm -f $(COV_REPORT) $(LINT_LOG)
