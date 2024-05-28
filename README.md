# Sistema de Stress test

Objetivo: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.


O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.


Execução do Teste:

    Realizar requests HTTP para a URL especificada.
    Distribuir os requests de acordo com o nível de concorrência definido.
    Garantir que o número total de requests seja cumprido.

Geração de Relatório:

    Apresentar um relatório ao final dos testes contendo:
        Tempo total gasto na execução
        Quantidade total de requests realizados.
        Quantidade de requests com status HTTP 200.
        Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

    Execução da aplicação:

    Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
        docker run <sua imagem docker> —url=http://google.com —requests=1000 —concurrency=10


---


# Desafio Go: Stress test

Os stress tests são metodologias rigorosas empregadas para avaliar a resistência e adaptabilidade de sistemas e aplicações sob condições extremas e adversas.
Seu objetivo principal é simular um ambiente hostil, saturando o sistema com uma avalanche de solicitações e cargas de trabalho, para entender suas limitações, identificar pontos de falha e assegurar que ele possa operar eficientemente em cenários de alta pressão.

## Pré-requisitos

Antes de começar, certifique-se de ter o seguinte instalado em sua máquina:

- Docker
- Makefile
- Go
- Curl


## Instalação e Configuração

### Clonando o repositório

```bash
git clone https://github.com/ivofulco/goexpert-stress-test.git
cd goexpert-stress-test
```

## Compilação e Execução


1. **Buildnando o código-fonte:**

   ```bash
   go build -o loadtester ./cmd/loadtester
   ```

2. **Executar o teste de carga:**

   Substitua os valores de `--url`, `--requests` e `--concurrency` de acordo com suas necessidades.

   ```bash
   ./loadtester --url=http://google.com --requests=100 --concurrency=10
   ```

## Executando via Docker

1. **Construir a imagem Docker:**

   ```bash
   docker build -t goexpert-stress-test .
   ```

2. **Executar o teste de carga via Docker:**

   Substitua os valores de `--url`, `--requests` e `--concurrency` de acordo com suas necessidades.

   ```bash
   docker run goexpert-stress-test --url=http://google.com --requests=100 --concurrency=10
   ```

## Resultado do Teste

Após a execução, o sistema apresentará um relatório detalhado que inclui:

- Tempo total gasto na execução
- Total de requests realizados
- Quantidade de requests com status HTTP 200
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
a