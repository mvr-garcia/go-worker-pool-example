# Go Order Processing System

This is a simple example of an order processing system in Go, simulating the operation of a restaurant with two workers ("Candier" and "Stringer") processing orders in parallel.

## Contents

- [Basic Concepts](#basic-concepts)
- [Code Structure](#code-structure)
- [How to Run](#how-to-run)

## Basic Concepts

### Goroutine

In Go, goroutines are lightweight units of execution, similar to lightweight threads, managed by the Go runtime.

### Channel

Channels are data structures in Go used for communication between goroutines. They allow the sending and receiving of data between goroutines.

### WaitGroup

WaitGroup is a way to wait for several goroutines to finish execution before continuing the execution of the main goroutine.

## Code Structure

The code consists of:

- Definition of an `Order` type to represent orders.
- Functions to process orders (`processOrder`) and wait for new orders (`waitForOrders`).
- A `worker` function representing a worker that receives and processes orders from a channel.
- The `main` function creates a channel, two goroutines for the workers, and simulates the creation and sending of orders.

## How to Run

To run the code, make sure you have the Go environment installed and execute the following command:

```bash
go run main.go
```

## Credits
[Go Concurrency Visually Explained](https://blog.stackademic.com/go-concurrency-visually-explained-channel-c6f88070aafa)