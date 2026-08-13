# Sistema de Gestão de Ordens de Serviço

Sistema de gerenciamento de Ordens de Serviço (OS) desenvolvido em **Go**, com persistência dos dados em arquivo JSON, aplicação de regras de negócio através de uma camada Service, interface de Repository, CLI interativa e sistema de auditoria utilizando goroutines.

## 🔗 Repositório

**GitHub:** [LucasDuarteV/sistema-de-gestao-de-ordens-de-servico](https://github.com/LucasDuarteV/sistema-de-gest-o-de-ordens-de-servi-o.git?utm_source=chatgpt.com)

## 📋 Sobre o projeto

O projeto foi desenvolvido para praticar conceitos importantes da linguagem Go e de organização de aplicações, implementando um CRUD completo de Ordens de Serviço.

O sistema permite:

* Criar Ordens de Serviço
* Listar Ordens de Serviço
* Buscar uma OS por ID
* Atualizar uma OS
* Deletar uma OS
* Validar regras de negócio
* Calcular dias restantes até a entrega
* Registrar ações de criação, atualização e exclusão
* Utilizar goroutines para realizar a auditoria de forma assíncrona

## 🛠️ Tecnologias utilizadas

* **Go**
* JSON
* Goroutines
* Interfaces
* CLI
* `sync`
* Pacotes da biblioteca padrão do Go

## 📁 Estrutura do projeto

```text
sistema-os/
│
├── auditoria/
│   └── auditoria.go
│
├── cli/
│   └── menu.go
│
├── models/
│   └── ordem_servico.go
│
├── repository/
│   ├── repository.go
│   └── json_repository.go
│
├── service/
│   └── ordem_servico_service.go
│
├── dados.json
├── auditoria.log
├── go.mod
└── main.go
```

## 🏗️ Arquitetura

O projeto utiliza uma separação de responsabilidades entre as camadas:

```text
CLI
 │
 ▼
Service
 │
 ▼
Repository
 │
 ▼
JSON
```

A auditoria funciona de forma independente:

```text
Service
   │
   ▼
Auditoria
   │
   ▼
Goroutine
   │
   ▼
auditoria.log
```

### CLI

Responsável pela interação com o usuário através do terminal.

### Service

Responsável pelas regras de negócio da aplicação.

### Repository

Responsável pela persistência dos dados em JSON.

### Models

Contém as estruturas de dados utilizadas pelo sistema.

### Auditoria

Registra as operações de criação, atualização e exclusão utilizando goroutines.

## ⚙️ Funcionalidades

### Criar OS

Permite cadastrar uma nova Ordem de Serviço informando:

* ID
* Cliente
* Descrição
* Valor estimado
* Data de entrega

### Listar OS

Exibe todas as Ordens de Serviço cadastradas.

### Buscar OS

Permite localizar uma Ordem de Serviço através do seu ID.

### Atualizar OS

Permite alterar informações de uma OS existente.

### Deletar OS

Permite excluir uma OS após confirmação do usuário.

### Auditoria

As operações são registradas automaticamente no arquivo `auditoria.log`.

## 📌 Regras de negócio

* O ID deve ser maior que zero.
* Não é permitido cadastrar duas OS com o mesmo ID.
* O cliente é obrigatório.
* A descrição é obrigatória.
* O valor estimado não pode ser negativo.
* O valor final não pode ser negativo.
* Não é possível atualizar uma OS inexistente.
* Não é possível excluir uma OS inexistente.

## 📅 Datas

A data de entrega utiliza o formato:

```text
02/01/2006
```

Exemplo:

```text
30/08/2026
```

O sistema também calcula os dias restantes até a data de entrega.

## 🚀 Como executar

### 1. Clonar o repositório

```bash
git clone https://github.com/LucasDuarteV/sistema-de-gest-o-de-ordens-de-servi-o.git
```

### 2. Entrar na pasta

```bash
cd sistema-de-gestao-de-ordens-de-servico
```

### 3. Executar

```bash
go run .
```

## 💻 Menu

```text
===== SISTEMA DE ORDEM DE SERVIÇO =====
1 - Criar OS
2 - Listar OS
3 - Buscar OS por ID
4 - Atualizar OS
5 - Deletar OS
0 - Sair
Escolha uma opção:
```

## 📝 Persistência

Os dados das Ordens de Serviço são armazenados no arquivo:

```text
dados.json
```

## 📊 Auditoria

As operações realizadas são registradas em:

```text
auditoria.log
```

Exemplo:

```text
13/08/2026 13:30:00 - OS 10 criada
13/08/2026 13:35:00 - OS 10 atualizada
13/08/2026 13:40:00 - OS 10 deletada
```

## 🧪 Testes realizados

Foram realizados testes manuais envolvendo:

* Criação de OS
* Listagem
* Busca por ID
* Atualização
* Exclusão
* IDs duplicados
* IDs inválidos
* Cliente vazio
* Descrição vazia
* Valores negativos
* Datas inválidas
* Cancelamento de exclusão
* Auditoria de criação
* Auditoria de atualização
* Auditoria de exclusão

## 🎯 Objetivos de aprendizado

Este projeto teve como objetivo praticar:

* Go
* Structs
* Métodos
* Interfaces
* Ponteiros
* Pacotes
* Organização de projetos
* CRUD
* Persistência em JSON
* Tratamento de erros
* Regras de negócio
* CLI
* Manipulação de datas
* Goroutines
* Auditoria
* Separação de responsabilidades

## 👨‍💻 Autor

**Lucas Duarte**

Projeto desenvolvido para estudos e prática de desenvolvimento backend com Go.
