# Basic Proxy
A proxy server written in go that supports [GET, POST] for XML or Json requests

## Installation
1) Clone repo
2) create .env file in directory, set port and GIN\_MODE
3) go build
4) go run ./main.go

### Configuring SSL
> todo

## Basic Usage
1) Call either /proxy or /proxy\_xml depending on request body
2) Set header (proxy\_url) with destination URL
3) Enjoy

## Todo
- [ ] Callback support 
