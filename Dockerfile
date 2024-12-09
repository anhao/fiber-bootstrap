FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o ./app ./main.go


FROM alpine

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/app /app/bin/app


EXPOSE 3000

CMD ["./bin/app"]