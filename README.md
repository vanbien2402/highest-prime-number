# Highest-prime-number

This is application by Golang which takes input number and return to user the highest prime number.
### Feature

- [x] Build simple API using Gin framework
- [x] Find prime numbers using Sieve of Eratosthenes algorithm
- [x] Containerize Application Using Docker
- [x] Deploy app to AWS ECS using GitHub workflows ci-cd
- [x] API Documentation Using Swagger

### The following diagram is my app infrastructure design

![Alt text](/docs/infra_design.drawio.png?raw=true "infra")

### How to run app local
- Clone source code
```bash
git clone https://github.com/vanbien2402/highest-prime-number.git -b master
```

- Navigate to source code directory
- Run app in local
```sh
go mod download
```
```sh
go run cmd/main.go    
```
- Run app in docker
```sh
docker compose build
```
```sh
docker compose up
```
- Try to call API with Postman
```sh 
http://localhost:8080/prime_number?inputNumber=10
```
You should get "The highest prime number lower than 10 is 7"
- Unit test
```sh 
go test ./...
```
## Thank you for visiting me!
### Author
- Van Bien