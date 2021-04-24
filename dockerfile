FROM node:lts-alpine as js-builder

WORKDIR /app/web-ui

COPY web-ui/package*.json ./

RUN npm ci 

COPY . /app

RUN npm run build

#####

FROM golang:1.16-alpine as go-builder

WORKDIR /app

ENV CGO_ENABLED 0 

COPY go.* ./

RUN go mod download

COPY . .

COPY --from=js-builder /app/web-ui/dist pkg/assets/web-ui

RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.6.1

RUN go generate

RUN go build -o sharepass *.go

#####

FROM alpine:3.13

RUN apk --update add ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

VOLUME [ "/app/data" ]

EXPOSE 5000

COPY --from=go-builder /app/sharepass ./

CMD ["/app/sharepass"]