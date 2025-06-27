.PHONY: watch
watch:
	go tool air -c .air.toml

.PHONY: build
build: ui sqlc
	go build -o ./tmp/main .

.PHONY: ui
ui:
	npm run build --prefix ./ui

.PHONY: sqlc
sqlc:
	go tool sqlc generate

.PHONY: test
test: build
	go tool gotestsum