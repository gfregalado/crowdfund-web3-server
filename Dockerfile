FROM golang:1.19-alpine

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
COPY .app.env .

RUN go build -o ./out/dist ./cmd/crowdfund-api
CMD ./out/dist