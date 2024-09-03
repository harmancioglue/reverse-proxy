# Go Simple Reverse Proxy

This repository aims to show how to create simple reverse proxy in Go.

## Tech

- Go

## Intro

![image]([https://github.com/harmancioglue/go-worker-pool-pattern/assets/27441734/93f76717-0518-4f09-ae01-4c8b1fcf2395](https://miro.medium.com/v2/resize:fit:1400/1*_BUaIvoqUTqxyZvyH9ky-A.png))

A reverse proxy is a server that sits between client devices and a backend server. It receives requests from clients and forwards them to the appropriate backend server. 

The backend server processes the request and sends the response back to the reverse proxy, which then forwards it to the client. 

This setup helps improve security, load balancing, and scalability by abstracting the backend server's details from the client and managing multiple servers efficiently.

## Run

docker-compose up --build
