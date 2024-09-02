# MenuFy 🍴
### Sobre o projeto
<p>O projeto tem como objetivo criar uma aplicação de cardápio digital para restaurantes e estabelecimentos de alimentação.<br>
A aplicação permitirá que os usuários criem, editem, apaguem e visualizem os produtos que farão parte do seu menu.<br>
</p>

### Por que "Menufy"?
<p>O nome "Menufy" foi escolhido para refletir o propósito e os valores do projeto. "Menufy" é uma combinação de "Menu" e "Simplificar".<br>
É uma escolha que busca combinar a essência do projeto com uma palavra que seja fácil de lembrar e pronunciar..<br>
</p>

### Por que Go?
<p>Go foi escolhida para este projeto devido ao seu desempenho, concorrência eficiente e simplicidade.<br>
É amplamente utilizada em empresas de grande escala, como Google, Netflix e Dropbox.<br>
E utilizada em projetos de código aberto, como Docker, Kubernetes e Prometheus.<br>
</p>

### Requisitos funcionais
-  O usuário pode se registrar enviando uma solicitação de registro para o administrador;
-  O administrador pode aprovar ou rejeitar a solicitação de registro;
-  O usuário pode se logar com seu `nome de usuário` e `senha` após aprovação da solicitação de registro
-  O usuário cadastra um produto informando o ```nome```, ```descrição```, ```preço```;
-  O usuário pode editar os dados de um produto cadastrado;
-  O usuário pode apagar um produto cadastrado;
-  O usuário pode visualizar todos os produtos cadastrados em forma de um cardápio;

### Fluxo da aplicação
- **Registro:** O usuário solicita seu registro na plataforma preenchendo um formulário com informações como nome, senha, email e número de telefone.
- **Aprovação:**  A solicitação de registro é enviada ao administrador do sistema, que pode aprovar ou rejeitar a solicitação.
- **Login:** Se a solicitação for aprovada, o usuário pode se logar no sistema com seu nome de usuário e senha.
- **Configuração do Estabelecimento:** Ao acessar o sistema pela primeira vez, o usuário é solicitado a preencher um formulário com informações do estabelecimento, como nome, CNPJ, endereço e número de contato.
- **Criação do Menu:** O usuário pode criar seu próprio menu digital, adicionando produtos com informações como nome, descrição, preço e imagem.
- **Gerenciamento do Menu:** O usuário pode editar, apagar e visualizar todos os produtos em seu menu personalizado.
- **Exibição do Menu:** O menu digital é exibido em uma URL exclusiva e personalizada do estabelecimento, com a logo do estabelecimento no topo e as informações do estabelecimento no rodapé.
 
### Tecnologias utilziadas
- Go
- Gin
- Docker
- PostgreSQL
- Tailwind CSS

# Back-end
### Configuração e criação do projeto
- [ ]  Criar projeto usando Go e Gin
- [ ]  Configurar conteinerização com Docker
- [ ]  Configurar banco de dados na aplicação (PostgreSQL)
- [ ]  Utilizar Docker para containerização da aplicação

## Infrastructure
- [ ]  Criar tabela de `User`
- [ ]  Criar tabela de `RegistrationRequest`
- [ ]  Criar tabela de `Product`
- [ ]  Configurar conexão com o banco de dados
- [ ]  Criar migrations para alterar a estrutura do banco de dados

## Domain Model
- [ ]  Criar entidades que irá representar um `User`
  - User
    - `id`(primary key, uuid): identificado único do usuário
    - `name`(string): nome do usuário
    - `email`(string): email do usuário
    - `password`(string): senha do usuário (hashed)
    - `phone_number`(string): contato do usuário
    - `status`(string): status do usuário ("active", "inactive")
```go
type User struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    Password string    `json:"password"`
    PhoneNumber string `json:"phone_number"`
    Status      string `json:"status"`
}
```

- [ ]  Criar entidades que irá representar um `RegistrationRequest`
  - RegistrationRequest
    - `id`(primary key, uuid): identificado único do usuário que solicitou registro
    - `name`(string): nome do usuário que solicitou o registro
    - `password`(string): senha do solicitante do registro
    - `email`(string): email do solicitante do registro
    - `phone_number`(string): número de telefone do solicitante do registro
    - `created_at`(timestamp): data e hora em que a solicitação de registro foi criada
    - `status`(string): status da solicitação de registro (por exemplo, "pending", "approved", "rejected")
    - `ApprovedAt`(timestamp): data e hora em que a solicitação de registro foi aprovada
    - `ApprovedBy`(integer): identificador do usuário que aprovou a solicitação de registro
    - `RejectedAt`(timestamp): data e hora em que a solicitação de registro foi rejeitada
    - `RejectedReason`(string): motivo da rejeição da solicitação de registro
```go
type RegistrationRequest struct {
    ID          uint           `json:"id"`
    Name        string         `json:"name"`
    Password    string         `json:"password"`
    Email       string         `json:"email"`
    PhoneNumber string         `json:"phone_number"`
    CreatedAt   time.Time      `json:"created_at"`
    Status      string         `json:"status"`
    ApprovedAt  *time.Time     `json:"approved_at"`
    ApprovedBy  *uint          `json:"approved_by"`
    RejectedAt  *time.Time     `json:"rejected_at"`
    RejectedReason *string     `json:"rejected_reason"`
}
```

- [ ]  Criar entidades que irá representar um `Product`
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
- [ ] Criar tabela de usuários

```sql
CREATE TABLE user (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20),
  status VARCHAR(20) NOT NULL CHECK(status IN ('active', 'inactive'))
);
```
- [ ] Criar tabela de solicitação de registro

```sql
CREATE TABLE registration_request (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  status VARCHAR(20) NOT NULL CHECK(status IN ('pending', 'approved', 'rejected')),
  approved_at TIMESTAMP,
  approved_by INTEGER,
  rejected_at TIMESTAMP,
  rejected_reason VARCHAR(255),
  FOREIGN KEY (approved_by) REFERENCES users(id)
);
```
- [ ] Criar tabela de produtos

```sql
CREATE TABLE product (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  image_path VARCHAR(255) NOT NULL
);
```

## Repository
- [ ]  Criar repository da entidade `User`
  - UserRepository
    - Métodos para acessar e manipular os dados do usuário no banco de dados

- [ ]  Criar repository da entidade `RegistrationRequest`
  - RegistrationRequestRepository
    - Métodos para acessar e manipular os dados da solicitação de registro no banco de dados
   
- [ ]  Criar repository da entidade `Product`
  - ProductRepository
    - Métodos para acessar e manipular os dados do produto no banco de dados

## Service
- [ ]  Criar serviço que irá lidar com as operações de `User`
  - UserService
    - Métodos para criar, editar, excluir e consultar usuários
    - Utilizar o repository para acessar os dados do usuário
    - Utilizar o domain model para validar e lidar com as regras de negócios do usuário
   
- [ ]  Criar serviço que irá lidar com as operações de `RegistrationRequest`
  - RegistrationRequestService
    - Métodos para criar, editar, excluir e consultar as solicitações de registro
    - Utilizar o repository para acessar os dados da solicitação de registro
    - Utilizar o domain model para validar e lidar com as regras de negócios da solicitação de registro
   
- [ ]  Criar serviço que irá lidar com as operações de `Produto`
  - ProductService
    - Métodos para criar, editar, excluir e consultar produtos
    - Utilizar o repository para acessar os dados do produto
    - Utilizar o domain model para validar e lidar com as regras de negócios do produto

## Endpoints
- [ ]  Criar endpoint de registro de usuário **POST** `/register`
  - Request Body:
```json
{
  "username": "username",
  "email": "email",
  "password": "password",
  "phone_number": "phone_number"
}
```
```
+ Response Body:
```
```json
{
  "message": "Solicitação de registro enviada com sucesso.",
  "registration_request_id": "example-uuid-1234"
}
```
- [ ]  Criar endpoint de aprovação de solicitação de registro **POST** `/approve-registration`
  - Request Body:
```json
{
  "registration_request_id": "example-uuid-1234",
  "approved_by": "admin_username"
}
```
- [ ]  Criar endpoint de login **POST** `/login`
  - Request Body:
```json
{
  "username": "username",
  "password": "password"
}
```
```
+ Response Body:
```
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaGFuIjoiMjMwfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```
     
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

# Front-end
- [ ]  Criar tela de cadastro de produtos que permita ao usuário cadastrar um produto informando o `nome`, `descrição`, `preço`.
- [ ]  Criar tela de listagem de produtos que exiba todos os produtos cadastrados no sistema.
