FROM golang:1.26 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
      go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o ./main .


FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /build/main ./main

ENTRYPOINT ["./main"]