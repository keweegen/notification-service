<h1 align="center">Notification service</h1>
<p align="center">
This is <strong>not the final example</strong> of a notification service implementation. At the moment, sending messages to Telegram, E-mail is supported. There is support for message templates for different sending channels.
</p>

----

## ✨ Setup
Requirements:
- [Docker](https://docs.docker.com/desktop/) + [Docker Compose plugin](https://docs.docker.com/compose/)
- [Golang](https://go.dev/)
- Make (optional)

```
make setup
```

## 📄 API Documentation
[OpenAPI 3.0 - YAML schema](https://github.com/keweegen/notification-service/docs/api.yaml)

## 🪄 Console commands
Run Docker services (Redis + PostgreSQL)
```
make d-up
```
---
Run Go service
```
make http-server
```
---
Generate mocks (gomock) / models (sqlboiler)
```
make generate
```
---
Run tests (go test -v ./...)
```
make test
```
---
Run tests with coverage (output: ./coverage.html)
```
make coverage
```

## ⭐ TODO
- [ ] Add validation requests
- [ ] Add [Jaeger](https://www.jaegertracing.io/)
