Spiral APP

Requirements

Docker (tested in 20.10.6)

Run locally

sh local-start

http://localhost:8080

Backend

Language: Golang.

The backend was built using hexagonal architecture.

cmd
-> httpserver
-> main.go
internal
-> core
-> handler
-> repository
pkg

Frontend
