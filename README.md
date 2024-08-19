# MenuFy 🍴
### Sobre o projeto
<p>O projeto tem como objetivo criar uma aplicação de cardápio digital para restaurantes e estabelecimentos de alimentação.<br>
A aplicação permitirá que os usuários criem, editem, apaguem e visualizem os produtos que farão parte do cardápio.<br>
</p>

### Por que "Menufy"?
<p>O nome "Menufy" foi escolhido para refletir o propósito e os valores do projeto. "Menufy" é uma combinação de "Menu" e "Simplificar".<br>
É uma escolha que busca combinar a essência do projeto com uma palavra que seja fácil de lembrar e pronunciar..<br>
</p>

### Requisitos funcionais
-  O usuário cadastra um produto informando o ```nome```, ```descrição```, ```preço```;
-  O usuário pode editar os dados de um produto cadastrado;
-  O usuário pode apagar um produto cadastrado;
-  O usuário pode visualizar todos os produtos cadastrados em forma de um cardápio;

### Tecnologias utilziadas
- Go
- Gin
- Docker
- PostgreSQL
- HTMX
- Tailwind CSS

# Back-end
### Configuração e criação do projeto
- [ ]  Criar projeto usando Go e Gin
- [ ]  Configurar conteinerização com Docker
- [ ]  Configurar banco de dados na aplicação (PostgreSQL)
- [ ]  Utilizar Docker para containerização da aplicação

## Infrastructure
- [ ]  Criar tabela de `Products` no banco de dados
- [ ]  Configurar conexão com o banco de dados
- [ ]  Criar migrations para alterar a estrutura do banco de dados

## Domain Model
- [ ]  Criar entidades que irá representar um `Produto`
  - Product
    - `ID`(primary key, uuid): identificado único do produto
    - `Name`(string): nome do produto
    - `Description`(string): descrição do produto
    - `Price`(decimal):preço do produto
    - `ImagePath`(string): caminho do arquivo de imagem do produto

```go
type Product struct {
    ID          uuid.UUID       `json:"id"`
    Name        string          `json:"name"`
    Description string          `json:"description"`
    Price       decimal.Decimal `json:"price"`
    ImagePath   string          `json:"image_path"`
}
```

## Database Schema

- [ ]  Schema do banco dados correspondente a entidade de `Product`

```sql
CREATE TABLE products (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    image_path VARCHAR(255)
);
```

## Repository
- [ ]  Criar repository da entidade `Produto`
  - ProductRepository
    - Métodos para acessar e manipular os dados do produto no banco de dados

## Service
- [ ]  Criar serviço que irá lidar com as operações de `Produto`
  - ProductService
    - Métodos para criar, editar, excluir e consultar produtos
    - Utilizar o repository para acessar os dados do produto
    - Utilizar o domain model para validar e lidar com as regras de negócios do produto

## Endpoints
- [ ]  Criar endpoint de cadastro de produto **POST** `/product`
  - Request Body:
```json
{
  "name": "Product Name",
  "description": "Product Description",
  "price": 19.99,
  "image_path": "https//example.com/image.jpg"
}
```
- [ ]  Criar endpoint de consulta de produto **GET** `/product/{productId}`
  - Response Body:
```json
{
  "id": "example-uuid-1234",
  "name": "Product Name",
  "description": "Product Description",
  "price": 19.99,
  "image_path": "https//example.com/image.jpg"
}
```
- [ ]  Criar endpoint de atualização de produto **PUT** `/product/{productId}`
  - Request Body:
```json
{
  "name": "New Product Name",
  "description": "New Product Description",
  "price": 29.99,
  "image_path": "https//example.com/new_image.jpg"
}
```
- [ ]  Criar endpoint de visualização de todos os produtos **GET** `/products`
  - Response Body:
```json
{
  "id": "exampleuuid1234",
  "name": "Product Name",
  "description": "Product Description",
  "price": 19.99,
  "image_path": "https//example.com/image.jpg"
},
{
  "id": "exampleuuid5678",
  "name": "Another Product",
  "description": "Another Description",
  "price": 55.50,
  "image_path": "https//example.com/another_product_image.jpg"
}
```
- [ ]  Criar endpoint de exclusão de produto **DELETE** `/product/{productId}`
  - Sem request body ou response boy para este endpoint


## Testes
### Escrever testes para cada endpoint da API
- [ ]  Testar cadastro de produto com dados válidos
- [ ]  Testar cadastro de produto com dados inválidos (nome vazio, preço negativo)
- [ ]  Verificar se o produto é criado com sucesso
- [ ]  Testar consulta de produto por ID
- [ ]  Testar consulta de produto com ID inexistente
- [ ]  Verificar se a API retorna o produto correto
- [ ]  Testar atualização de produto com dados válidos
- [ ]  Testar atualização de produto com dados inválidos (nome vazio, preço negativo)
- [ ]  Verificar se o produto é atualizado com sucesso
- [ ]  Testar exclusão de produto com ID existente
- [ ]  Testar exclusão de produto com ID inexistente
- [ ]  Verificar se o produto é excluído com sucesso
- [ ]  Testar visualização de todos os produtos
- [ ]  Verificar se a API retorna a lista de produtos correta
- [ ]  Testar cadastro de produto com imagem inválida (string que não é uma URL válida)
- [ ]  Testar cadastro de produto com uma imagem muito grande
- [ ]  Testar cadastro de produto com tipo de arquivo de imagem inválido(arquivo de texto, csv, xlsx)
- [ ]  Testar atualização do produto com imagem nova
- [ ]  Testar atualização de produto sem imagem

## Visão da estrutura
```
API (Endpoints)
  |
  |-- Controller (Recebe requisições HTTP)
  |    |
  |    |-- Service (Processa requisições e chama repositórios)
  |    |    |
  |    |    |-- Repository (Acessa dados do banco de dados)
  |    |    |-- Domain Model (Define as regras de negócios)
```
# Front-end
- [ ]  Criar tela de cadastro de produtos que permita ao usuário cadastrar um produto informando o `nome`, `descrição`, `preço`.
- [ ]  Utilizar HTMX para lidar com as requisições de cadastro de produtos.
- [ ]  Utilizar Tailwind CSS para estilizar os elementos HTML.
- [ ]  Criar tela de listagem de produtos que exiba todos os produtos cadastrados no sistema.
