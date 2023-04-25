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

ENV PROD 1

ENV COOKIE_NAME ${COOKIE_NAME}
ENV COOKIE_PATH ${COOKIE_PATH}

ENV POSTGRES_DB ${POSTGRES_DB}
ENV POSTGRES_PASSWORD ${POSTGRES_PASSWORD}
ENV POSTGRES_USERNAME ${POSTGRES_USERNAME}

ENV EMAIL_USER ${EMAIL_USER}
ENV EMAIL_PASSWORD ${EMAIL_PASSWORD}
ENV EMAIL_HOST ${EMAIL_HOST}
ENV EMAIL_PORT ${EMAIL_PORT}

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/configs /app/configs

CMD ["./main"]