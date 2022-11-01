<h1 align="center">Notification service</h1>
<p align="center">
This is <strong>not the final example</strong> of a notification service implementation. At the moment, sending messages to Telegram, E-mail is supported. There is support for message templates for different sending channels.
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/keweegen/notification-service">
    <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B?style=flat-square">
  </a>
  <a href="https://gocover.io/github.com/keweegen/notification-service">
    <img src="https://img.shields.io/badge/%F0%9F%94%8E%20gocover-97.8%25-75C46B.svg?style=flat-square">
  </a>
  <a href="https://github.com/keweegen/notification-service/actions?query=workflow%3ASecurity">
    <img src="https://img.shields.io/github/workflow/status/keweegen/notification-service/Security?label=%F0%9F%94%91%20gosec&style=flat-square&color=75C46B">
  </a>
  <a href="https://github.com/keweegen/notification-service/actions?query=workflow%3ATest">
    <img src="https://img.shields.io/github/workflow/status/keweegen/notification-service/Test?label=%F0%9F%A7%AA%20tests&style=flat-square&color=75C46B">
  </a>
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
[OpenAPI 3.0 - YAML schema](https://raw.githubusercontent.com/keweegen/notification-service/main/docs/api.yaml)

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
