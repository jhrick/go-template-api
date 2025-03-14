FROM golang:1.23.0

WORKDIR /app

ADD . .

COPY go.mod go.sum ./

RUN go install github.com/air-verse/air@latest

ENV PATH="$PATH:/go/bin"

RUN go mod download

COPY . .

RUN go build -o /app/cmd/main /app/cmd/main.go

CMD ["air", "-c", ".air.toml"]