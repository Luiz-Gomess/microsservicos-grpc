#!/bin/bash

# 1. Configurar variáveis
GITHUB_USERNAME=Luiz-Gomess
GITHUB_EMAIL=fernando.albuquerque@academico.ifpb.edu.br
SERVICE_NAME=order
RELEASE_VERSION=v1.2.3

# 2. Instalar dependências do sistema (Ubuntu)
echo "--- Verificando e instalando dependências do sistema ---"
sudo apt-get update
sudo apt-get install -y protobuf-compiler golang-go git

# 3. Instalar plugins do Go para Protobuf e gRPC
echo "--- Instalando plugins Go (protoc-gen-go e protoc-gen-go-grpc) ---"
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 4. Configurar o PATH para a sessão atual do script
# Isso garante que o 'protoc' consiga encontrar os plugins que acabamos de instalar
export PATH="$PATH:$(go env GOPATH)/bin"

# 5. Gerar código
echo "--- Gerando código fonte Go ---"
# Cria o diretório de destino se não existir
mkdir -p golang

# Nota: Certifique-se de que o arquivo .proto existe em ./payment/*.proto
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

# 6. Inicializar módulo Go
# O script original tenta entrar numa pasta específica. 
# Se o protoc não gerou a pasta com o nome do serviço (depende do go_package no .proto),
# ajustaremos para criar o mod na raiz do output ou na pasta específica.

TARGET_DIR="golang/${SERVICE_NAME}"

#