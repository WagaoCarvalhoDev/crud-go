# Use a imagem oficial do Golang
FROM golang:1.21.1 AS BUILDER

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o código-fonte do aplicativo para o diretório de trabalho
COPY . .
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
 GOOS=linux go build -o crudgo .

FROM golang:1.21.1 as runner

RUN adduser user

COPY --from=BUILDER /app/crudgo /app/crudgo
COPY .env /app/.env 

RUN chown -R user:user /app
RUN chmod +x /app/crudgo

EXPOSE 8080

USER user

# Configure a entrada do contêiner
CMD ["/app/crudgo"]
