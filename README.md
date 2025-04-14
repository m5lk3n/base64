# base64

`base64` implemented as a web service in Golang.

Build:

```bash
make docker-build
```

Run:

```bash
docker run -d \
  --restart unless-stopped \
  -u "$(id -u):$(id -g)" \
  -p 8080:8080 \
  --name base64 \
  lttl.dev/base64:0.1.0
```

Use:

```bash
curl -sX OPTIONS http://localhost:8080 | jq
{
  "routes": [
    "GET /spenc/:string - Base64-encodes the string with a newline appended",
    "GET /encode/:string - Encodes the string in base64",
    "GET /decode/:string - Decodes a base64 string",
    "GET /health - Provides basic health info"
  ],
  "status": 200
}
```

Stop:

```bash
docker stop base64
docker rm -f base64
```