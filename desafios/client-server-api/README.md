# Desafio Go: Webserver HTTP, Contextos, Banco de Dados e Manipulação de Arquivos

## Descrição do Desafio

Este projeto tem como objetivo aplicar conceitos de webserver HTTP, contextos, banco de dados e manipulação de arquivos utilizando a linguagem Go (Golang). O desafio consiste em implementar dois sistemas: `client.go` e `server.go`.

- **client.go**: Realiza uma requisição HTTP para o servidor (`server.go`) solicitando a cotação do dólar.
- **server.go**: Consome uma API para obter a cotação do dólar em relação ao real e retorna o valor para o cliente. Além disso, registra cada cotação recebida em um banco de dados SQLite.

### Requisitos

1. O `client.go` deve fazer uma requisição HTTP para o `server.go` no endpoint `/cotacao` para obter a cotação do dólar.
2. O `server.go` deve consumir a API [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL) para obter a cotação do dólar e retornar o valor para o `client.go` em formato JSON.
3. O `server.go` deve utilizar o pacote `context` para:
    - Limitar o tempo de chamada da API de cotação do dólar a 200ms.
    - Limitar o tempo para persistir os dados no banco de dados a 10ms.
4. O `client.go` deve usar o pacote `context` para definir um timeout máximo de 300ms ao esperar a resposta do `server.go`.
5. Caso os limites de tempo sejam ultrapassados, os erros devem ser registrados nos logs.
6. O `client.go` deve salvar o valor da cotação em um arquivo chamado `cotacao.txt` no formato: `Dólar: {valor}`.

### Endpoints

- **GET /cotacao**: Retorna a cotação atual do dólar obtida da API e registrada no banco de dados.

### Tecnologias Utilizadas

- Go (Golang)
- Pacote `context` do Go
- Banco de Dados SQLite
- HTTP Client/Server
- Manipulação de arquivos

## Passos para Execução

### Pré-requisitos

- Go instalado na máquina (versão 1.16 ou superior)
- Conexão com a internet para acessar a API de cotação
- [Driver SQLite para Go](https://github.com/mattn/go-sqlite3) instalado.

### Instruções para Executar o Projeto

1. **Clone o repositório para sua máquina local:**

    ```bash
    git clone https://github.com/gustavoggsb/go-expert.git
    ```

2. **Navegue até o diretório do projeto:**

    ```bash
    cd go-expert/desafios/client-server-api
    ```

3. **Instale as dependências:**

    ```bash
    go mod tidy
    ```

4. **Execute o servidor:**

    O servidor vai iniciar um webserver na porta 8080 e ficará ouvindo requisições HTTP.

    ```bash
    go run ./server/main.go
    ```

5. **Execute o cliente:**

    O cliente fará uma requisição HTTP ao servidor para obter a cotação do dólar e salvará o valor no arquivo `cotacao.txt`.

    ```bash
    go run ./client/main.go
    ```

### Observações

- O arquivo `cotacao.txt` será criado no mesmo diretório em que o `client.go` é executado, contendo o valor atual da cotação do dólar.
- Caso ocorram erros relacionados a timeouts, eles serão registrados nos logs.
- Certifique-se de que o servidor esteja em execução antes de rodar o cliente, caso contrário, o `client.go` não conseguirá fazer a requisição.

### Estrutura do Projeto

```txt
client-server-api/                # Diretório raiz do projeto
├── client/                       # Diretório do cliente
│   ├── cotacao.txt               # Arquivo de saída com a cotação (gerado pelo cliente)
│   └── main.go                   # Código fonte do cliente
├── server/                       # Diretório do servidor
│   ├── cotacoes.db               # Banco de dados SQLite para armazenar as cotações (gerado pelo servidor)
│   └── main.go                   # Código fonte do servidor
├── .gitignore                    # Arquivo para ignorar arquivos desnecessários no git
├── go.mod                        # Arquivo de definição de módulos Go
├── go.sum                        # Arquivo de checagem de dependências do Go
└── README.md                     # Instruções e documentação do projeto
```

## Contato

Se você tiver dúvidas ou sugestões sobre este projeto, sinta-se à vontade para abrir uma *issue* ou enviar um *pull request*. 

**Autor:** Gustavo Baptista<br>
**Email:** gustavogsb@outlook.com  
**GitHub:** [gustavoggsb](https://github.com/gustavoggsb)

