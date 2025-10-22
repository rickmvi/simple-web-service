# 🔎 GitHub User Finder - Web Service Simples

Este é um mini projeto de web service construído em Go que expõe uma API para buscar dados de um usuário específico do GitHub e exibe esses resultados em uma interface web simples e estilizada.

---

## ✨ Recursos

- **Serviço de Backend em Go**: Uma API RESTful minimalista (`/api/user?username=...`) para interagir com a API pública do GitHub.
- **Busca de Perfil**: Retorna informações-chave do perfil, como login, nome, bio, URL do avatar, repositórios públicos, seguidores e seguindo.
- **Frontend Simples**: Interface web responsiva e escura (dark theme) para a entrada do nome de usuário e exibição dos resultados.
- **Tratamento de Erros**: Gerenciamento de erros de requisição e notificação de usuários não encontrados.

---

## 🛠️ Tecnologias Utilizadas
- **Backend**:
  - **Go (Golang)**: Linguagem principal para o web service.
- **Frontend**:
  - **HTML5**
  - **CSS3** (com um tema _dark_ elegante)
  - **JavaScript** (para interação assíncrona com a API)

---

## 🚀 Como Executar

Estes passos presumem que você tem o Go instalado em seu sistema.

**1. Clonar o Repositório**
```bash
git clone https://github.com/rickmvi/simple-web-service.git
cd simple-web-service
```

**2. Rodar o Servidor Go**
O projeto está configurado para usar módulos Go e o servidor pode ser iniciado diretamente:
```bash
go run main.go
```
Após a execução, você verá a mensagem:
```
Listening on port 3000
```

**3. Acessar a Aplicação**

Abra seu navegador e acesse:
```
http://localhost:3000
```

A interface web será carregada, e você poderá começar a buscar usuários do GitHub.

---

## 💻 Estrutura do Projeto

O projeto segue uma estrutura modular para separar a lógica de negócios, a camada de API e o ponto de entrada principal.
```
simple-web-service/
├── internal/
│   ├── api/          # Lógica para construir URLs da API externa
│   └── domain/       # Lógica de manipulação de dados, o Handler da API
├── static/           # Arquivos estáticos (HTML, CSS, JS)
│   ├── CSS/
│   ├── Js/
│   └── index.html
└── main.go           # Ponto de entrada e configuração do servidor HTTP
```
---

## 🔗 Uso da API (Backend)
O web service expõe um único endpoint principal para a busca de usuários:

**Endpoint:** `/api/user` **Método:** `GET`

**Parâmetros de Query**

| Parâmetro     | Tipo |                        Descrição                        | Exemplo |
|:------| :------: |:-------------------------------------------------------:| :------- |
| username | string  | **Obrigatório**. O nome de usuário do GitHub a ser buscado. | octocat |

**Exemplo de Requisição**
```http
GET http://localhost:3000/api/user?username=torvalds
```
**Exemplo de Resposta (Sucesso -** `200 OK`**)**
```json
{
  "login": "torvalds",
  "name": "Linus Torvalds",
  "avatar_url": "...",
  "bio": "...",
  "public_repos": 7,
  "followers": 184567,
  "following": 0
}
```

**Exemplo de Resposta (Erro - Usuário Não Encontrado** `404 Not Found`**)**
```
GitHub returned: 404 - {"message":"Not Found","documentation_url":"..."}
```
---
## 🧩 Implementação em Go (Visão Geral)

- `internal/api/url.go`: Contém a função `GetUrl` que constrói de forma segura o _endpoint_ da API do GitHub para um dado nome de usuário.
- `internal/domain/github_user.go`:
  - Define a struct `GithubUser` para o _marshalling_ dos dados JSON.
  - Contém a função `Handler`, que é a principal responsável por:
     
    _1._ Validar o método HTTP e o parâmetro `username`.
      
    _2._ Fazer a requisição HTTP para a API do GitHub.
      
    _3._ Tratar a resposta (`200 OK`, `404`, ou outros erros).

    _4._ Desserializar a resposta JSON e enviar o resultado formatado para o cliente.
- `main.go`: Configura o servidor HTTP, serve os arquivos estáticos da pasta `/static` e registra o _handler_ da API em `/api/user`.

---

## 🎨 Interface Web (Visão Geral)

O frontend é um SPA (Single Page Application) simples que usa JavaScript nativo para interagir com o backend:

- `static/index.html`: A estrutura básica da página.
- `static/Js/script.js`: Contém a lógica de busca:

  _1._ Captura o evento de submit do formulário.
    
  _2._ Faz um `fetch` assíncrono para o endpoint `/api/user` do Go.
    
  _3._ Em caso de sucesso, preenche os campos do HTML com os dados do usuário.
    
  _4._ Em caso de falha (erro de rede ou status `!response.ok`), exibe a mensagem "**Usuário não encontrado 😢**".
    
- `static/CSS/style.css`: Define o visual do aplicativo, incluindo o tema _dark_ e a estilização dos resultados.
