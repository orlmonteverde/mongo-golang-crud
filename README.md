# Golang and MongoDB CRUD

Introduction to go-mongo-driver. This as an alternative to the well-known mgo package through a simple CRUD.

![MongoDB](https://img.shields.io/badge/DevsStore-MongoDB-brightgreen.svg?logo=mongodb&longCache=true&style=flat) ![Go badge](https://img.shields.io/badge/DevsStore-golang-blue.svg?logo=go&longCache=true&style=flat)

## Getting Started

This project uses the **Go** programming language (Golang) and the **MongoDB** database engine.

### Prerequisites

[MongoDB](https://www.mongodb.com/) is required in version 3 or higher and [Go](https://golang.org/) at least in version 1.12

### Installing

The dependency required for this project is the MongoDB driver for Go.

* go.mongodb.org/mongo-driver/mongo

#### Using GOPATH

```
go get -u -v go.mongodb.org/mongo-driver/mongo
```

#### Using GOMODULE
```
go build
```

## Deployment

Clone the repository
```
git clone git@github.com:orlmonteverde/mongo-golang-crud.git
```
Enter the repository folder
```
cd mongo-golang-crud
```
Build the binary
```
go build
```
Run the program
```
# In Unix-like OS
./mongo-golang-crud

# In Windows
mongo-golang-crud.exe
```

## Built With

* [go-mongo-driver](https://github.com/mongodb/mongo-go-driver) - The Go driver for MongoDB

## Authors

* **Orlando Monteverde** - *Initial work* - [orlmonteverde](https://github.com/orlmonteverde)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
