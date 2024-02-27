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
