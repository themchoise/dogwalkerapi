COVERAGE_FILE=coverage.out

coverage:
	go test -coverprofile=$(COVERAGE_FILE) ./...

clean:
	rm -f $(COVERAGE_FILE)


open-coverage: coverage
	code $(COVERAGE_FILE)
