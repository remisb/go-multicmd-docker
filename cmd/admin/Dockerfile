FROM golang:1.22-alpine AS builder

WORKDIR /app

# fetch dependencies
COPY ./go.*  ./
RUN GOPROXY=https://proxy.golang.org \
  --mount=type=cache,target=/go/pkg/mod/ \
  go mod download -x

# build
COPY . .
WORKDIR /app/cmd/admin
RUN CGO_ENABLED=0 go build -o admin .

FROM alpine AS admin_runtimer
COPY --from=builder /app/cmd/admin /app
EXPOSE 8081
ENTRYPOINT [ "/app/admin" ]
