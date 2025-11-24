# go-worker
This project demonstrates an asynchronous processing architecture using Go and RabbitMQ.
It consists of two main components:

- Producer API: A Go-based HTTP API that receives requests and publishes messages to a RabbitMQ queue.

- Worker Consumer: A Go worker service that continuously listens to the queue and processes incoming tasks.

The project showcases how to integrate Go applications with RabbitMQ, implement message publishing and consumption, and build a basic asynchronous workflow suitable for background jobs, task processing, and event-driven systems.
