build-local:
	docker image build --target production --build-arg ENV=local -f ./cmd/api/Dockerfile -t upsider/api-local .
build-dev:
	docker image build --target production --build-arg ENV=dev -f ./cmd/api/Dockerfile -t upsider/api-dev .

compose:
	docker compose -f ./build/docker-compose.yaml up
compose-build:
	docker compose -f ./build/docker-compose.yaml up --build

golangci:
	go list -f '{{.Dir}}/...' -m | xargs golangci-lint run -c configs/.golangci.yaml
test:
	go test -v -count=1 -cover ./cmd/... ./internal/...