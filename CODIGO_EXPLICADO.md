# 📖 Explicação Detalhada do Código

## 🔍 Imports (Importações)

### 📦 Imports Padrão do Go

```go
import (
    "log"        // Sistema de logging (registro de mensagens)
    "net/http"   // Servidor HTTP e funcionalidades web
    "os"         // Sistema operacional (variáveis de ambiente)
)
```

#### `"log"`
- **O que é**: Pacote para registrar mensagens no console
- **Para que serve**: Mostrar informações sobre o que está acontecendo na aplicação
- **Exemplo**: `log.Printf("Servidor iniciando na porta %s", port)`

#### `"net/http"`
- **O que é**: Pacote para criar servidores web
- **Para que serve**: Receber e responder requisições HTTP
- **Exemplo**: `http.ListenAndServe(":8080", r)` - inicia o servidor

#### `"os"`
- **O que é**: Pacote para interagir com o sistema operacional
- **Para que serve**: Acessar variáveis de ambiente, arquivos, etc.
- **Exemplo**: `os.Getenv("PORT")` - pega a porta do ambiente

### 📦 Imports do Nosso Projeto

```go
import (
    "study-go/internal/handlers"  // Nossos manipuladores de requisições
    "study-go/pkg/middleware"     // Nossos middlewares
)
```

#### `"study-go/internal/handlers"`
- **O que é**: Nossos próprios handlers (manipuladores)
- **Para que serve**: Processar requisições HTTP (GET, POST, PUT, DELETE)
- **Exemplo**: `handlers.SetupRoutes(r)` - configura as rotas

#### `"study-go/pkg/middleware"`
- **O que é**: Nossos middlewares (interceptadores)
- **Para que serve**: Processar requisições antes de chegarem aos handlers
- **Exemplo**: Logging, CORS, autenticação

### 📦 Imports Externos

```go
import (
    "github.com/gorilla/mux"  // Roteador HTTP externo
)
```

#### `"github.com/gorilla/mux"`
- **O que é**: Biblioteca externa para roteamento HTTP
- **Para que serve**: Criar rotas mais poderosas que o padrão do Go
- **Exemplo**: `mux.NewRouter()` - cria um roteador

---

## 🔧 Explicação do Código Principal

### 1. Criação do Roteador
```go
r := mux.NewRouter()
```
- **O que faz**: Cria um novo roteador HTTP
- **Por que**: Para organizar as rotas da nossa API
- **Analogia**: Como criar um mapa de ruas para direcionar o tráfego

### 2. Configuração de Middlewares
```go
r.Use(middleware.LoggingMiddleware)
r.Use(middleware.CORSMiddleware)
```
- **O que faz**: Adiciona interceptadores às requisições
- **LoggingMiddleware**: Registra cada requisição (quem acessou, quando, quanto tempo levou)
- **CORSMiddleware**: Permite que sites de outros domínios acessem nossa API
- **Analogia**: Como filtros que processam cada requisição antes de chegar ao destino

### 3. Configuração das Rotas
```go
handlers.SetupRoutes(r)
```
- **O que faz**: Define todas as rotas da API
- **Exemplo de rotas**:
  - `GET /health` → Verifica se a API está funcionando
  - `GET /api/users` → Lista todos os usuários
  - `POST /api/users` → Cria um novo usuário
- **Analogia**: Como definir os endereços e o que cada um faz

### 4. Configuração da Porta
```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```
- **O que faz**: Pega a porta do ambiente ou usa 8080 como padrão
- **Por que**: Flexibilidade para rodar em diferentes ambientes
- **Exemplo**: Se você definir `PORT=3000`, a API rodará na porta 3000

### 5. Inicialização do Servidor
```go
log.Printf("Servidor iniciando na porta %s", port)
log.Fatal(http.ListenAndServe(":"+port, r))
```
- **O que faz**: Inicia o servidor HTTP
- **`log.Printf`**: Mostra uma mensagem informando que o servidor está iniciando
- **`http.ListenAndServe`**: Inicia o servidor e fica esperando requisições
- **`log.Fatal`**: Se der erro, para a aplicação e mostra o erro

---

## 🔄 Fluxo de uma Requisição

```
1. Cliente faz requisição → http://localhost:8080/api/users
2. CORSMiddleware → Verifica se pode aceitar a requisição
3. LoggingMiddleware → Registra a requisição
4. Roteador → Encontra a rota correta (/api/users)
5. Handler → Processa a requisição (GetUsers)
6. Resposta → Retorna os dados para o cliente
```

---

## 📝 Exemplo Prático

Quando você acessa `http://localhost:8080/api/users`:

1. **Servidor recebe**: Requisição GET para `/api/users`
2. **CORSMiddleware**: "Ok, aceito requisições de qualquer origem"
3. **LoggingMiddleware**: "Registrando: GET /api/users às 14:30:25"
4. **Roteador**: "Encontrei! Vou para o handler GetUsers"
5. **GetUsers**: "Vou retornar a lista de usuários"
6. **Resposta**: `[{"id":1,"name":"João Silva","email":"joao@email.com"}]`

---

## 🎯 Por que essa Estrutura?

### ✅ **Vantagens**
- **Organizado**: Cada parte tem sua responsabilidade
- **Escalável**: Fácil adicionar novas funcionalidades
- **Manutenível**: Código limpo e bem estruturado
- **Testável**: Cada parte pode ser testada separadamente

### 🔧 **Padrões Usados**
- **Separação de responsabilidades**: Cada arquivo tem uma função específica
- **Middleware pattern**: Interceptadores para funcionalidades comuns
- **Router pattern**: Organização de rotas
- **Handler pattern**: Processamento de requisições

---

## 🚀 Próximos Conceitos para Aprender

1. **Banco de Dados**: Como conectar e salvar dados
2. **Autenticação**: Como proteger a API
3. **Validação**: Como verificar se os dados estão corretos
4. **Testes**: Como testar o código
5. **Docker**: Como empacotar a aplicação 