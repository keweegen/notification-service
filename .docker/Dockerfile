FROM golang:1.19-alpine3.16 as build
WORKDIR /app
ENV CGOENABLED=0
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY *.go ./
RUN go build -o /notification-service

FROM alpine3.16
COPY --from=build /notification-service /notification-service
ENTRYPOINT ["/notification-service"]
