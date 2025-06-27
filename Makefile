.PHONY: watch
watch:
	go tool air -c .air.toml

.PHONY: build
build: ui
	go build -o ./tmp/main .

.PHONY: ui
ui:
	npm run build --prefix ./ui

.PHONY: test
test: build
	go tool gotestsum