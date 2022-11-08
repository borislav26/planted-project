# Planted-project

A couple of notes about project:
  - On the backend I have used go programming language and on frontend is React.js library
  - Backend is organized like 3 layer application with repo, service and api layer
  - Gin is used as package for server running and api endpoints making
  
In order you want to run unit tests on backend, checkout project and position yourself in /backend repository and run command:
  - `go test ./...`
  or
  - `go test ./...`
  
Both commands will run all tests on backend, but second one will show tests that has been run with statuses.

For ranking of zipcodes, I think there is mistake with data where for every entry we have same zipcode, but code is implemented in way to support different
values and ranking.
