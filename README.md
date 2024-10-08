# Go Simple Reverse Proxy

This repository aims to show how to create simple reverse proxy in Go.

## Tech

- Go

## Intro

![image](https://github.com/user-attachments/assets/6ba74ca2-1cab-4541-8f0d-0ac65f36aa3d)


A reverse proxy is a server that sits between client devices and a backend server. It receives requests from clients and forwards them to the appropriate backend server. 

The backend server processes the request and sends the response back to the reverse proxy, which then forwards it to the client. 

This setup helps improve security, load balancing, and scalability by abstracting the backend server's details from the client and managing multiple servers efficiently.

According to our project, requests are forwarded to different services based on the paths to which the requests are made.

Example: /server1 http://server1:8081

## Run

docker-compose up --build
