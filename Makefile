runPub:
	go run ./cmd/pub/main.go

runService:
	go run ./cmd/service/main.go

runNatsStreaming:
	docker run -ti -p 4222:4222 -p 8223:8223 nats-streaming -p 4222 -m 8223

hello:
	echo "hello"

test:
	go test -race -timeout 30s ./...

.DEFAULT_GOAL := hello