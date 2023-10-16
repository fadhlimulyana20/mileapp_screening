# Mile App Screening Test

## How To Run
### Run Using Docker
```
cp .env.example .env
docker compose up
```
then open http://localhost:5000

### Run Manually
- Set up .env file. The mandatory application are postgres and mongo
- then run,
```
go run main go serve ---port=5000
```
- then open http://locahost:5000

## How To Access Documentation
Documentation can be accessed by open the http://localhost:5000/swagger link after running the REST API.

## Unit Test
Unit test are located in /internal/handler directory and /database directory.

## Created By Golang-Kit
Golang Kit is a framework to build microservices  a boilerplate for building microservice. It is written with Go with clean architecture. Currently, it only can be used to build REST API microservice.

Golang Kit can be accessed [here](https://github.com/fadhlimulyana20/golang_kit)

## Documentation
The Golang Kit documentation can be found in this [link](https://fadhlimulyana.notion.site/Golang-Go-Kit-3dd318e1eacd46fcbd99e20df8fa57e8)

## Feature
- HTTP Server
- Database
- Migration
- Seeder
- Authorization (Casbin)
- JWT
- AES Encryption
- Mailer
- Minio (S3)
- Stub
- Clean Architecture

## Supported DB
- Postgres
- MongoDB
