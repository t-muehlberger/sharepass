# SharePass - Secure password sharing utility



**Features:**

- End to end encryption
- Link expires after 7 days
- Link expires showing 3 times

**Demo:**

https://sharepass.muehlberger.dev

## Building

**Prerequisites:**

- Go 1.16+
- Node.js 14+
- NPM

**Steps:**
  
    cd web-ui
    npm install
    npm run build

    cd ..
    go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.6.1
    go generate
    go build

**SSL Dev Cert:**

- Set up [mkcert](https://github.com/FiloSottile/mkcert)

    cd web-ui
    mkcert localhost