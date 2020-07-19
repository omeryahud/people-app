FROM golang:1.13 as builder

ENV DIR $HOME/go/src/github.com/omeryahud/people-app
ENV BINARY database

WORKDIR ${DIR}

COPY ./ ./

#RUN go mod download
RUN go mod tidy
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build --mod=vendor -a -o bin/${BINARY} cmd/${BINARY}/main.go

FROM gcr.io/distroless/static:nonroot

ENV DIR $HOME/go/src/github.com/omeryahud/people-app/bin
ENV BINARY database

WORKDIR /

COPY --from=builder ${DIR}/${BINARY} .

USER nonroot:nonroot


ENTRYPOINT ["/database"]