FROM golang:1.23.3-alpine as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# -----------------------------------------------

FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

RUN chmod +x ./app

CMD ["./app"]

# -----------------------------------------------

# FROM golang:1.22.3-alpine as dev
# WORKDIR /app
# RUN go install github.com/air-verse/air

# CMD ["air"]