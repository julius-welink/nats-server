docker buildx build --progress=plain --platform linux/amd64,linux/arm64/v8 -t juliuszaromskis/nats-server:latest -t juliuszaromskis/nats-server:2.11.8.1  -f ./docker/Dockerfile --push .
