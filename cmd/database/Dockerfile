FROM golang:1.13 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/database/main.go main.go
COPY internal/pkg internal/pkg

RUN go mod tidy
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build --mod=vendor -a -o database main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /workspace/database .
USER nonroot:nonroot

ENTRYPOINT ["/database"]