# base64

`base64` implemented as a web service in Golang.

Usage:

```bash
curl -sX OPTIONS http://localhost:8080 | jq
{
  "routes": [
    "GET /spenc/:string - Base64-encodes the string with a newline appended",
    "GET /encode/:string - Encodes the string in base64",
    "GET /decode/:string - Decodes a base64 string",
    "GET /health - Health check endpoint"
  ],
  "status": 200
}
```