# Rate Limiter 

Para rodar o projeto: 

## Setup

    - Docker e docker-compose instalado em sua máquina. 

## Execução da Aplicação

Você pode executar a aplicação facilmente utilizando Docker. Basta seguir o exemplo abaixo:


```bash
docker compose up 
``` 

Para realizar chamadas na aplicação, você pode utilizar as formas abaixo: 
    - Postman: 
        - localhost:8080 
        - API_KEY 123abc 

    - CURL 
        - curl -vvv http://localhost:8080
        - curl -H 'API_KEY: 123abc' -vvv http://localhost:8080

# Armazenamento
O armazenamento é estabelecido por meio de uma interface LimiterStrategyInterface, que contém o método Verificar para acessar e definir valores no armazenamento. 
Atualmente, estamos utilizando o Redis como banco de dados (conforme solicitado no desafio), porém possível criar e adicionar novas implementações para outros tipos (banco de dados, cache, memórias). A lógica de limitação continuará a mesma, somente sendo necessário implementar na instância de RateLimiter por meio do controlador de dependências, injectamos através dele.
