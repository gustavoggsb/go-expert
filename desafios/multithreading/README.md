# Desafio Go: Multithreading e APIs - Buscando a Resposta Mais Rápida

## Descrição do Desafio

Este projeto tem como objetivo aplicar conceitos de multithreading, concorrência e chamadas de APIs utilizando a linguagem Go (Golang). O desafio consiste em implementar um sistema que realize requisições simultâneas a duas APIs distintas para buscar informações sobre um CEP, aceitando a resposta mais rápida e descartando a mais lenta.

### Requisitos

1. Realizar requisições simultâneas para as seguintes APIs utilizando o mesmo CEP:
   - [BrasilAPI](https://brasilapi.com.br/api/cep/v1/01153000)
   - [ViaCEP](http://viacep.com.br/ws/01153000/json/)
   
2. Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

3. Exibir no terminal os dados do endereço retornados pela API mais rápida, incluindo qual API forneceu a resposta.

4. Limitar o tempo de resposta a 1 segundo. Caso o tempo seja ultrapassado, um erro de timeout deve ser exibido.

### Tecnologias Utilizadas

- Go (Golang)
- Pacote `context` do Go
- HTTP Client para requisições simultâneas
- Manipulação de JSON

## Passos para Execução

### Pré-requisitos

- Go instalado na máquina (versão 1.16 ou superior)
- Conexão com a internet para acessar as APIs de consulta de CEP

### Instruções para Executar o Projeto

1. **Clone o repositório para sua máquina local:**

    ```bash
    git clone https://github.com/gustavoggsb/go-expert.git
    ```

2. **Navegue até o diretório do projeto:**

    ```bash
    cd go-expert/desafios/multithreading
    ```

3. **Execute o programa:**

    O programa irá realizar requisições simultâneas às duas APIs utilizando o mesmo CEP e exibirá a resposta mais rápida no terminal.

    ```bash
    go run main.go
    ```

### Observações

- As respostas das APIs são diferentes em termos de estrutura, então o programa utiliza structs específicas para tratar cada resposta adequadamente.
- Em caso de timeout (mais de 1 segundo), o programa exibirá uma mensagem de erro indicando que o tempo limite foi atingido.
- Se ambas as APIs responderem dentro do tempo, apenas a resposta mais rápida será exibida.

### Estrutura do Projeto

```txt
multithreading-api/               # Diretório raiz do projeto
├── main.go                       # Código fonte principal do projeto
└── README.md                     # Instruções e documentação do projeto
```

## Contato

Se você tiver dúvidas ou sugestões sobre este projeto, sinta-se à vontade para abrir uma *issue* ou enviar um *pull request*. 

**Autor:** Gustavo Baptista<br>
**Email:** gustavogsb@outlook.com  
**GitHub:** [gustavoggsb](https://github.com/gustavoggsb)

