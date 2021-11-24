FROM golang:1.15-alpine as build
RUN apk add --no-cache git
RUN mkdir -p /app/api
COPY ./api/ /app/api
COPY go.* /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o main ./api

FROM alpine:3.9
RUN apk --update add ca-certificates
RUN mkdir /config
COPY --from=build /app/main /
ENTRYPOINT [ "/main" ]