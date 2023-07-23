FROM golang:1.20.3-alpine as base
WORKDIR /builder
# install upx and make
RUN apk add upx make git
ENV GO111MODULE=on CGO_ENABLED=0
COPY go.mod go.sum /builder/
RUN go mod tidy
COPY . .
RUN make build

FROM alpine:3.8
WORKDIR /app
COPY --from=base /builder/bin/analytics-service /app/bin/analytics-service
COPY --from=base /builder/db /app/db

CMD ["/app/bin/analytics-service", "server"]