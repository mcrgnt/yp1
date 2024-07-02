GOLANGCI_LINT_CACHE?=/tmp/praktikum-golangci-lint-cache

.PHONY: fieldalignment
fieldalignment:
	-@find . -type d -not -name ".*"|while read path; do if [ `find $$path -maxdepth 1 -name *.go|wc -l` -gt 0 ]; then fieldalignment -fix `find $$path -maxdepth 1 -name *.go`; fi; done

.PHONY: formattag
formattag:
	-@find . -name "*.go" -type f -exec formattag -file {} \;


.PHONY: golangci-lint-run
golangci-lint-run: _golangci-lint-rm-unformatted-report

.PHONY: _golangci-lint-reports-mkdir
_golangci-lint-reports-mkdir:
	mkdir -p ./golangci-lint

.PHONY: _golangci-lint-run
_golangci-lint-run: _golangci-lint-reports-mkdir
		-docker run --rm \
    -v $(shell pwd):/app \
    -v $(GOLANGCI_LINT_CACHE):/root/.cache \
    -w /app \
    golangci/golangci-lint:v1.57.2 \
        golangci-lint run \
            -c .golangci.yml \
	> ./golangci-lint/report-unformatted.json

.PHONY: _golangci-lint-format-report
_golangci-lint-format-report: _golangci-lint-run
	cat ./golangci-lint/report-unformatted.json | jq > ./golangci-lint/report.json

.PHONY: _golangci-lint-rm-unformatted-report
_golangci-lint-rm-unformatted-report: _golangci-lint-format-report
	rm ./golangci-lint/report-unformatted.json

.PHONY: golangci-lint-clean
golangci-lint-clean:
	sudo rm -rf ./golangci-lint

.PHONY: server
server:
	go generate ./...; cd cmd/server; DATABASE_DSN="postgres://yp1:yp1@localhost:5432/yp1" go run .

.PHONY: agent
agent:
	go generate ./...; cd cmd/agent; go run .

.PHONY: postgres
postgres:
	docker run -d -p5432:5432 -e POSTGRES_PASSWORD=yp1 -e POSTGRES_PASSWORD=yp1 -e POSTGRES_USER=yp1 -e POSTGRES_DB=yp1 postgres