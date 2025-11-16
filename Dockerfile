FROM golang:1.22.2-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
LABEL maintainer="Efimtsev Stanislav <evendot@yandex.ru>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["go", "run", "main.go"]