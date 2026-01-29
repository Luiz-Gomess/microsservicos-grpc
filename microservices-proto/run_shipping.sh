#!/bin/bash

GITHUB_USERNAME=Luiz-Gomess
GITHUB_EMAIL=fernando.albuquerque@academico.ifpb.edu.br
SERVICE_NAME=shipping
RELEASE_VERSION=v1.2.3

echo "--- Verificando e instalando dependências do sistema ---"
sudo apt-get update
sudo apt-get install -y protobuf-compiler golang-go git

echo "--- Instalando plugins Go (protoc-gen-go e protoc-gen-go-grpc) ---"
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"

echo "--- Gerando código fonte Go ---"
mkdir -p golang

if [ -d "./${SERVICE_NAME}" ]; then
    protoc --go_out=./golang \
      --go_opt=paths=source_relative \
      --go-grpc_out=./golang \
      --go-grpc_opt=paths=source_relative \
      ./${SERVICE_NAME}/*.proto
else
    echo "Erro: Diretório ./${SERVICE_NAME} não encontrado. Verifique onde está o arquivo .proto."
    exit 1
fi

echo "--- Arquivos gerados: ---"
ls -al ./golang/


TARGET_DIR="golang/${SERVICE_NAME}"

