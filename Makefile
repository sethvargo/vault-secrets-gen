VETTERS = "asmdecl,assign,atomic,bools,buildtag,cgocall,composites,copylocks,errorsas,httpresponse,loopclosure,lostcancel,nilfunc,printf,shift,stdmethods,structtag,tests,unmarshal,unreachable,unsafeptr,unusedresult"
GOFMT_FILES = $(shell go list -f '{{.Dir}}' ./...)

GIT_COMMIT = $(shell git rev-parse --short HEAD)

# List of ldflags
LD_FLAGS = \
	-s \
	-w \
	-X github.com/sethvargo/vault-secrets-gen/version.Name=vault-secrets-gen \
	-X github.com/sethvargo/vault-secrets-gen/version.GitCommit=${GIT_COMMIT}

benchmarks:
	@(cd benchmarks/ && go test -bench=. -benchmem -benchtime=1s ./...)
.PHONY: benchmarks

dev:
	@env \
		CGO_ENABLED="0" \
		go install \
			-ldflags "${LD_FLAGS}" \
			-tags "${GOTAGS}"
.PHONY: dev

test:
	@go test \
		-count=1 \
		-short \
		-timeout=5m \
		-vet="${VETTERS}" \
		./...
.PHONY: test

test-acc:
	@go test \
		-count=1 \
		-race \
		-timeout=10m \
		-vet="${VETTERS}" \
		./...
.PHONY: test-acc
