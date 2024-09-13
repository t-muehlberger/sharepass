FROM node:lts-alpine as js-builder

WORKDIR /app/web-ui

COPY web-ui/package*.json ./

RUN npm ci 

COPY . /app

RUN npm run build

#####

FROM golang:1.23-alpine as go-builder

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=js-builder /app/web-ui/dist pkg/assets/web-ui

RUN --mount=type=cache,target=/root/.cache/go-build \
    go generate && \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -o sharepass -ldflags "-w -s" *.go

#####

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:latest

RUN apk --update add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

VOLUME [ "/app/data" ]

EXPOSE 5000

COPY --from=go-builder /app/sharepass ./

CMD ["/app/sharepass"]