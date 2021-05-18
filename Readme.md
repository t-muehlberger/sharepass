# SharePass - Secure password sharing utility

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