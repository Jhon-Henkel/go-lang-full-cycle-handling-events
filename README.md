# go-lang-full-cycle-handling-events
Repositório para armazenar o código do módulo de manipulação de eventos do curso de Go Lang.
Esse repositório vai ser uma biblioteca para ler e escrever mensagens no RabbitMQ.

# Rodando testes 
Rodar na raiz do projeto o comando:
```bash
go test ./...
```

# Rodando o projeto
Para dar oi start no projeto, basta rodar o comando:
```bash
docker-compose up -d
```
Para acessar o Rabbit, acesse: http://localhost:15672/ com o usuário e senha: guest

Fazer um bind no exchange amq.direct com a fila my_queue no rabbit.