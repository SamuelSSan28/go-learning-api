# Study Go - API REST

Uma API REST simples em Go para estudos e aprendizado.

## 📁 Estrutura do Projeto

```
study-go/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── handlers/
│   │   └── handlers.go          # Manipuladores de requisições HTTP
│   └── models/
│       └── user.go              # Modelos de dados
├── pkg/
│   └── middleware/
│       └── middleware.go        # Middlewares (logging, CORS)
├── go.mod                       # Gerenciamento de dependências
├── go.sum                       # Checksums das dependências
├── README.md                    # Este arquivo
└── .gitignore                   # Arquivos ignorados pelo Git
```

## 📋 Explicação da Estrutura

### `cmd/` - Comandos da Aplicação
- **`cmd/api/main.go`**: Ponto de entrada principal da aplicação
  - Configura o servidor HTTP
  - Inicializa middlewares
  - Configura rotas
  - Define a porta do servidor

### `internal/` - Código Interno da Aplicação
- **`internal/handlers/handlers.go`**: Manipuladores de requisições
  - `GetUsers()` - Lista todos os usuários
  - `CreateUser()` - Cria um novo usuário
  - `GetUser()` - Busca um usuário específico
  - `UpdateUser()` - Atualiza um usuário
  - `DeleteUser()` - Remove um usuário
  - `HealthCheck()` - Verifica se a API está funcionando

- **`internal/models/user.go`**: Modelos de dados
  - `User` - Estrutura principal do usuário
  - `CreateUserRequest` - Dados para criar usuário
  - `UpdateUserRequest` - Dados para atualizar usuário

### `pkg/` - Pacotes Reutilizáveis
- **`pkg/middleware/middleware.go`**: Middlewares
  - `LoggingMiddleware` - Registra informações das requisições
  - `CORSMiddleware` - Permite requisições cross-origin

## 🛠️ Como Construir o Projeto

### Pré-requisitos
- Go 1.22 ou superior
- Git

### Passo a Passo

1. **Clone o repositório:**
   ```bash
   git clone <url-do-repositorio>
   cd study-go
   ```

2. **Instale as dependências:**
   ```bash
   go mod tidy
   ```

3. **Execute a aplicação:**
   ```bash
   go run cmd/api/main.go
   ```

4. **Acesse a API:**
   - URL base: `http://localhost:8080`
   - Health check: `http://localhost:8080/health`

## 📦 Dependências do Projeto

### Principais Dependências
- **`github.com/gorilla/mux`**: Roteador HTTP para criar APIs REST
  - Usado para definir rotas e métodos HTTP
  - Permite parâmetros de URL dinâmicos

### Dependências Padrão do Go
- **`net/http`**: Servidor HTTP padrão do Go
- **`encoding/json`**: Codificação/decodificação JSON
- **`log`**: Sistema de logging
- **`time`**: Manipulação de tempo

## 🔧 Arquivos de Configuração

### `go.mod`
```go
module study-go

go 1.22

require github.com/gorilla/mux v1.8.1
```
- Define o nome do módulo
- Especifica a versão do Go
- Lista as dependências

### `go.sum`
- Contém checksums das dependências
- Garante integridade dos pacotes
- Gerado automaticamente pelo Go

### `.gitignore`
- Ignora arquivos desnecessários no Git
- Binários compilados
- Arquivos de teste
- Logs e arquivos temporários

## 🚀 Funcionalidades

- ✅ **CRUD completo** de usuários
- ✅ **Middleware de logging** para monitorar requisições
- ✅ **Middleware CORS** para requisições cross-origin
- ✅ **Health check** endpoint
- ✅ **Estrutura organizada** seguindo boas práticas do Go
- ✅ **Documentação completa** do código

## 🔗 Endpoints da API

| Método | Endpoint | Descrição | Exemplo de Resposta |
|--------|----------|-----------|-------------------|
| GET | `/health` | Verifica se a API está funcionando | `{"status":"OK","message":"API funcionando corretamente"}` |
| GET | `/api/users` | Lista todos os usuários | `[{"id":1,"name":"João Silva","email":"joao@email.com"}]` |
| POST | `/api/users` | Cria um novo usuário | `{"id":3,"name":"Pedro Costa","email":"pedro@email.com"}` |
| GET | `/api/users/{id}` | Busca um usuário específico | `{"id":1,"name":"João Silva","email":"joao@email.com"}` |
| PUT | `/api/users/{id}` | Atualiza um usuário | `{"id":1,"name":"João Silva Atualizado","email":"joao@email.com"}` |
| DELETE | `/api/users/{id}` | Remove um usuário | `204 No Content` |

## 🧪 Exemplos de Uso

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

### Atualizar usuário
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "João Silva Atualizado", "email": "joao@email.com"}'
```

### Remover usuário
```bash
curl -X DELETE http://localhost:8080/api/users/1
```

## 📚 Conceitos de Go Aprendidos

### Estrutura de Projeto
- **Layout padrão do Go**: `cmd/`, `internal/`, `pkg/`
- **Separação de responsabilidades**: handlers, models, middleware
- **Módulos Go**: `go.mod` e `go.sum`

### HTTP e APIs
- **Servidor HTTP**: `net/http` package
- **Roteamento**: Gorilla Mux para APIs REST
- **JSON**: Codificação/decodificação de dados
- **Middleware**: Interceptação de requisições

### Boas Práticas
- **Tratamento de erros**: Verificação de erros em cada operação
- **Logging**: Registro de informações importantes
- **CORS**: Configuração para requisições cross-origin
- **Documentação**: Comentários explicativos no código

## 🔮 Próximos Passos

- [ ] **Banco de Dados**: Adicionar PostgreSQL ou MySQL
- [ ] **Autenticação**: Implementar JWT
- [ ] **Validação**: Adicionar validação de dados
- [ ] **Testes**: Implementar testes unitários
- [ ] **Documentação**: Adicionar Swagger/OpenAPI
- [ ] **Docker**: Containerização da aplicação
- [ ] **CI/CD**: Pipeline de integração contínua
- [ ] **Monitoramento**: Logs estruturados e métricas

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -m 'feat: adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes. 