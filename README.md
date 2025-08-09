# Study Go - API REST

Uma API REST simples em Go para estudos e aprendizado.

## Estrutura do Projeto

```
study-go/
├── cmd/api/main.go          # Ponto de entrada da aplicação
├── internal/
│   ├── handlers/handlers.go # Manipuladores de requisições
│   └── models/user.go       # Modelos de dados
├── pkg/middleware/          # Middlewares
└── README.md               # Este arquivo
```

## Funcionalidades

- ✅ CRUD completo de usuários
- ✅ Middleware de logging
- ✅ Middleware CORS
- ✅ Health check endpoint
- ✅ Estrutura organizada e escalável

## Endpoints da API

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/health` | Verifica se a API está funcionando |
| GET | `/api/users` | Lista todos os usuários |
| POST | `/api/users` | Cria um novo usuário |
| GET | `/api/users/{id}` | Busca um usuário específico |
| PUT | `/api/users/{id}` | Atualiza um usuário |
| DELETE | `/api/users/{id}` | Remove um usuário |

## Como Executar

1. **Instalar dependências:**
   ```bash
   go mod tidy
   ```

2. **Executar a API:**
   ```bash
   go run cmd/api/main.go
   ```

3. **Acessar a API:**
   - URL base: `http://localhost:8080`
   - Health check: `http://localhost:8080/health`

## Exemplos de Uso

### Criar um usuário
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "João Silva", "email": "joao@email.com"}'
```

### Listar usuários
```bash
curl http://localhost:8080/api/users
```

### Buscar usuário específico
```bash
curl http://localhost:8080/api/users/1
```

## Próximos Passos

- [ ] Adicionar banco de dados (PostgreSQL/MySQL)
- [ ] Implementar autenticação JWT
- [ ] Adicionar validação de dados
- [ ] Implementar testes unitários
- [ ] Adicionar documentação Swagger
- [ ] Configurar Docker 