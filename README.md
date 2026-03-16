# Notifications API

API de notificações assíncronas com WorkerPool em Go.

## Como funciona

```
POST /v1/notifications  -->  Handler  -->  UseCase  -->  Enqueue(channel)
                                                              |
                                                         WorkerPool
                                                         (5 goroutines)
                                                              |
                                                         dispatch (email/SMS)
```

1. O handler recebe o JSON e valida
2. O use case cria a notificação e enfileira no channel
3. A API responde 202 imediatamente (não espera o envio)
4. Workers consomem o channel em background e processam (simulação de envio)

## WorkerPool

- **5 workers** (goroutines) consumindo do mesmo channel
- **Buffer de 100** mensagens
- Cada worker lê do channel, processa e loga (simula envio real)
- Shutdown gracioso ao encerrar a aplicação

## Quick start

```bash
make run
```

API em `http://localhost:8000`. Swagger em `http://localhost:8000/swagger/index.html`.

## Comandos

| Comando   | Descrição                    |
|-----------|------------------------------|
| `make run`   | Sobe a API (gera swagger antes) |
| `make build` | Compila para `bin/senai-lab365` |
| `make swagger` | Regenera documentação        |
| `make clean` | Remove bin/ e docs/          |

## Exemplo de requisição

```bash
curl -X POST http://localhost:8000/v1/notifications \
  -H "Content-Type: application/json" \
  -d '{"user_id":"user-123","message":"Sua fatura vence amanhã","priority":"high"}'
```

**Body:** `user_id`, `message`, `priority` (low | medium | high)

## Estrutura

```
cmd/api/          # Entry point
internal/
  domain/         # Entidades e interfaces
  application/    # Use cases
  infrastructure/ # NotificationDispatcher (WorkerPool)
  interfaces/     # Handlers HTTP e DTOs
```

## PR Agent

Review automático de PRs via [PR Agent](https://github.com/qodo-ai/pr-agent). Configure o secret `OPENAI_KEY` em Settings > Secrets and variables > Actions.
