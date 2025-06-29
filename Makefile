.PHONY: watch
watch:
	go tool air -c .air.toml

.PHONY: build
build: ui sqlc
	go build -o ./tmp/main .

.PHONY: bin
bin:
	./tmp/main \
		-smtp-host 127.0.0.1 \
		-smtp-port 1025 \
		-smtp-user "" \
		-smtp-pass "" \
		-mail-from "Uptime Monitor <no-reply@uptimemonitor.dev>"

.PHONY: ui
ui:
	npm run build --prefix ./ui

.PHONY: sqlc
sqlc:
	go tool sqlc generate

.PHONY: test
test: build
	go tool gotestsum