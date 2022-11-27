# syntax=docker/dockerfile:1

# **************************************************
# ----- Build stage
# **************************************************
FROM golang:1.18-alpine AS BuildStage

WORKDIR /workspace

ADD go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build ./cmd/main.go

# **************************************************
# ----- Build Container image
# **************************************************
FROM alpine:3.14

RUN apk add --no-cache  && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=BuildStage /workspace/main ./main

RUN chmod 775 ./main

EXPOSE 8080

CMD [ "./main" ]
