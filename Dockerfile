FROM golang:1.22-alpine3.19

RUN go install github.com/air-verse/air@v1.52.3
RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml" , "-d"]
