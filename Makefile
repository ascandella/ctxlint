PROJECT_ROOT := github.com/sectioneight/ctxlint

LINT_LOG := lint.log
COV_REPORT := overalls.coverprofile

PKG_FILES = . ctxlint

SOURCE_DEPS = *.go $(wildcard ctxlint/*.go)

.PHONY: lint
lint:
	gofmt -d -s $(PKG_FILES) | tee -a $(LINT_LOG)
	$(foreach dir,$(PKG_FILES),go tool vet $(VET_RULES) $(dir) 2>&1 | tee -a $(LINT_LOG);)
	$(foreach dir,$(PKG_FILES),golint $(dir) 2>&1 | tee -a $(LINT_LOG);)
	@[ ! -s $(LINT_LOG) ]

.PHONY: test
test: $(COV_REPORT)

$(COV_REPORT): $(SOURCE_DEPS)
	overalls -project=$(PROJECT_ROOT) -- -race -v | \
		grep -v "No Go Test Files"
	gocov convert $(COV_REPORT) | gocov report

.PHONY: coveralls
coveralls: $(COV_REPORT)
	@goveralls -coverprofile=$(COV_REPORT)

.PHONY: dependencies
dependencies:
	@which overalls >/dev/null || go get -u github.com/go-playground/overalls
	@which goveralls >/dev/null || go get -u -f github.com/mattn/goveralls
	@which gocov >/dev/null || go get github.com/axw/gocov/gocov
	@which golint >/dev/null || go get github.com/golang/lint/golint
	@go get github.com/stretchr/testify/assert
	@go get github.com/stretchr/testify/require

.PHONY: clean
clean:
	rm -f $(COV_REPORT) $(LINT_LOG)

coverage.html: $(SOURCE_DEPS) $(COV_REPORT)
	@which gocov-html >/dev/null || go get github.com/awx/gocov-html
	gocov convert $(COV_REPORT) | gocov-html > coverage.html

.DEFAULT_GOAL: test
