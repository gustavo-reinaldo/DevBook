# DevBook

Uma aplicação que simula o funcionamento de uma rede social.

## Pré-requisitos

Antes de começar, você precisará ter instalado as seguintes ferramentas:

- [Go](https://golang.org/) (versão 1.19 ou superior recomendada)
- [Git](https://git-scm.com/)
- [MySQL](https://dev.mysql.com/downloads/)

Certifique-se também de que o banco de dados MySQL esteja rodando.

## Configuração do Banco de Dados

1. Inicie o servidor MySQL.
2. Crie um banco de dados chamado `devbook` (ou outro nome de sua preferência).
3. Atualize as credenciais de conexão ao banco de dados no arquivo `config.json` (ou no local onde elas estão configuradas). Certifique-se de incluir:
   - Host
   - Usuário
   - Senha
   - Nome do banco de dados

   Exemplo de configuração:
   ```json
   {
     "host": "localhost",
     "user": "root",
     "password": "sua-senha",
     "database": "devbook"
   }

## Como rodar o projeto
git clone https://github.com/gustavo-reinaldo/DevBook.git
cd devbook

## Rodando a API
1. Navegue até o diretório da API:
cd api
2. Execute a API 
go run main.go

## Rodando o front-end do projeto:
1. Navegue até o diretório do frontend:
cd webapp
2. Execute o front 
go run main.go

## Acessando a aplicação
1. No seu browser, pesquise por: http://localhost:3000/
