FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /you-together

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o /app/main cmd/main.go
ADD /configs /app/configs


FROM scratch

MAINTAINER matvey-sizov@mail.ru
ENV CGO_ENABLED 0
ENV PROD 1

ENV POSTGRES_DB "you-together"

ENV POSTGRES_PASSWORD "postgres"

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/configs /app/configs

CMD ["./main"]