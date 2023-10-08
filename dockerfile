FROM golang:1.21.2-alpine3.18 as base

RUN apk update
WORKDIR /src/gotaco

COPY go.mod go.sum ./
COPY . .
RUN go build -o gotaco ./cmd/api

FROM alpine:3.18 as binary
COPY --from=base /src/gotaco/gotaco .
EXPOSE 3000
CMD [ "./gotaco" ]
