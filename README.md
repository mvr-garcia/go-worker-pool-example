# Sistema de Pedidos em Go

Este é um exemplo simples de um sistema de pedidos em Go, que simula o funcionamento de um restaurante com dois trabalhadores ("Candier" e "Stringer") que processam pedidos em paralelo.

## Conteúdo

- [Conceitos Básicos](#conceitos-básicos)
- [Estrutura do Código](#estrutura-do-código)
- [Como Executar](#como-executar)

## Conceitos Básicos

### Goroutine

Em Go, goroutines são unidades leves de execução, semelhantes a threads leves, gerenciadas pelo runtime.

### Canal (Channel)

Canais são estruturas de dados em Go usadas para comunicação entre goroutines. Eles permitem o envio e recebimento de dados entre goroutines.

### WaitGroup

WaitGroup é uma maneira de aguardar que várias goroutines terminem a execução antes de continuar a execução da goroutine principal.

## Estrutura do Código

O código consiste em:

- Definição de um tipo `Order` para representar pedidos.
- Funções para processar pedidos (`processOrder`) e esperar por novos pedidos (`waitForOrders`).
- Uma função `worker` que representa um trabalhador que recebe e processa pedidos de um canal.
- A função principal (`main`) cria um canal, duas goroutines para os trabalhadores e simula a criação e envio de pedidos.

## Como Executar

Para executar o código, certifique-se de ter o ambiente Go instalado e execute o seguinte comando:

```bash
go run main.go
