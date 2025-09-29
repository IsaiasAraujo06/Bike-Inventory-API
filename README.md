# Bike Inventory API

API REST para gerenciamento de inventário da Bike Sports Fitness.

## 🚀 Tecnologias

- Go 1.21+
- SQLite
- Standard Library (net/http, database/sql)

## 📋 Pré-requisitos

- Go 1.21 ou superior
- SQLite3

## 🔧 Instalação
```bash
# Clone o repositório
git clone https://github.com/IsaiasAraujo06/bike-inventory-api.git

# Entre no diretório
cd bike-inventory-api

# Copie o arquivo de configuração
cp .env.example .env

# Baixe as dependências
go mod download

# Execute a aplicação
go run cmd/api/main.go
```

## 📚 Endpoints

### Produtos

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/products` | Criar novo produto |
| GET | `/products` | Listar todos os produtos |
| GET | `/products/:id` | Buscar produto por ID |
| PUT | `/products/:id` | Atualizar produto |
| DELETE | `/products/:id` | Deletar produto |
| GET | `/products/low-stock` | Produtos com estoque baixo |

## 🏗️ Estrutura do Projeto
```
bike-inventory-api/
├── cmd/api/          # Entry point
├── internal/         # Código da aplicação
│   ├── database/     # Camada de banco de dados
│   ├── models/       # Modelos de dados
│   └── handlers/     # HTTP handlers
└── migrations/       # Migrações do banco
```

## 📝 Roadmap

- [x] Setup inicial do projeto
- [ ] Configuração do banco de dados
- [ ] Implementar CRUD básico
- [ ] Adicionar validações
- [ ] Implementar logging estruturado
- [ ] Adicionar testes
- [ ] Deploy em produção

## 👨‍💻 Autor

Isaías Araújo - [@IsaiasAraujo06](https://github.com/IsaiasAraujo06)

## 📄 Licença

Este projeto está sob a licença MIT.