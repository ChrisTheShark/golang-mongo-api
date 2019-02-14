# golang-mongo-api

![Build Status](https://travis-ci.org/ChrisTheShark/golang-mongo-api.svg?branch=master) 

Lead Maintainer - [Chris Dyer](https://github.com/ChrisTheShark)

Simple, testable api to retrieve users from MongoDB.

## Gettting Started

This repository uses the [dep](https://github.com/golang/dep) tool for dependency management. First, clone this repository and at the root of the project execute ```dep ensure```. This command will go get all dependencies. Provide your connection string in the form ```mongodb://localhost:27017``` as an environment variable named MONGO_HOST. Build or run the application using ```go run *.go``` or ```go build *.go```. If using build follow up with an execution of the created binary. 
