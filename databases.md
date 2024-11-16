
# databases

[sql](./sql.md)
[postgresql](./postgresql.md)

### ER Diagram

- strong entity: rectangle/square
  - can be uniquely identified by its own attributes without relying on other entity
- weak entity: rectangle that contains another rectangle or balloon with dotted underline
  - cannot be uniquely identified by its own attributes
  - depends on a strong entity to be identified
- relationship: diamond
  - id relationship: diamond that contains another diamond
  - recursive relationship: occurs when an entity has multiple functions with itself
- attribute: empty dot
  - identifying attribute: black dot or balloon with underlined attribute
  - multivalued attributes: balloon that contains another balloon
  - derived attributes: dotted baloon
    - attribute that can be defined from another attributes
- inheritance: triangle
  - total or partial
    - total: all instances of superclass must be instances of subclass
    - partial: some instances of superclass must be instances of subclass
  - overlapped or disjoint
    - disjoint: entities in superclass can only be instance of one subclass
    - overlapped: entities in superclass can instances of multiples subclasses simultaneously


#### crow's notation

- relationships are bidirected
- there are symbols on both sides that represent min:max cardinality

- `[ student ] >|--has--|| [ school ]`
  - student relates with one and exactly one school
  - school relates to one or many students

- `--------0|` = 0:1
- `--------0<` = 0:n
- `--------||` = 1:1 (one and exactly one)
- `--------|<` = 1:n

#### chen notation

- `[ student ] 1----<has>---1 [ seat ]`
  - one student can have maximum of 1 seat
  - one seat belongs to maximum of 1 student

#### barker notation

- `#`: primary key
- `*` or `·`: obligatory attribute
- `o`: optional attribute

```
+----------------+          +----------------+          +----------------+
|   Customer     |          |     Order      |          |    Product     |
|----------------|          |----------------|          |----------------|
|*CustomerID     |     1,M  |*OrderID        |          |*ProductID      |
| Name           |----------| OrderDate      |     1,M  | ProductName    |
| Email          |          |#CustomerID     |----------| Price          |
+----------------+          +----------------+          +----------------+
                              0,M  |                         0,M   |
                                   |                               |
                                   |                               |
                                   +-------------------------------+
                                                1,M
                                           +------------+
                                           |  OrderItem |
                                           |------------|
                                           |*OrderItemID|
                                           |#OrderID    |
                                           |#ProductID  |
                                           | Quantity   |
                                           +------------+
```

### Class diagram

- cardinality symbols
  - how to read class diagrams:
    - Entity1 ----------------------> 0..* Entity2
    - Entity1 can have 0..* relations with Entity2
      - the multiplicity is closer to the entity it applies to
  - multiplicity
    - 1: one to one
    - *: many
    - 0..1: zero or one
    - 0..*: zero or many
    - 1..*: one or many
    - n..m: n to m

## theory

- SGBD: Sistema Geral de Banco de Dados
- SBD: Sistema de Banco de Dados
- SBD = SGBD + BD

- três categorias de modelos de dados:
  - conceituais (alto nível)
    - diagrama Entidade-Relacionamento (ER)
    - usuários típicos: analistas de negócios, designers
    - registra quais dados são registrados, mas não como
  - lógicos (representativos)
    - modelo lógico de dados
    - usuários típicos: desenvolvedores, DBAs
  - físicos (baixo nível)
    - modelo físico de dados
    - usuários típicos: administradores de banco de dados

Modelo Lógico de Dados para um Sistema de Vendas:

| Tabela   | Colunas |
| --- | --- |
| Clientes | ClienteID (INT, PK), Nome (VARCHAR), Email (VARCHAR)|
| Pedidos  | PedidoID (INT, PK), ClienteID (INT, FK), Data (DATE)|
| Produtos | ProdutoID (INT, PK), NomeProduto (VARCHAR), Preço (DECIMAL)|
| ItensPedido | ItemID (INT, PK), PedidoID (INT, FK), ProdutoID (INT, FK), Quantidade (INT) |

Modelo Físico de Dados para um DBMS Relacional (e.g., MySQL):

```sql
CREATE TABLE Clientes (
    ClienteID INT PRIMARY KEY,
    Nome VARCHAR(100) NOT NULL,
    Email VARCHAR(100) UNIQUE
);

CREATE TABLE Pedidos (
    PedidoID INT PRIMARY KEY,
    ClienteID INT,
    Data DATE,
    FOREIGN KEY (ClienteID) REFERENCES Clientes(ClienteID)
);

CREATE INDEX idx_email ON Clientes(Email);
```
