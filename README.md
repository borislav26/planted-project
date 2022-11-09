# Planted project

A couple of notes about project:
  - On the backend I have used go programming language and on frontend is React.js library
  - Backend is organized like 3 layer application with repo, service and api layer
  - Gin is used as package for server running and api endpoints making
  
In order you want to run unit tests on backend, checkout project and position yourself in /backend repository and run command:
  - `go test ./...` or `go test ./... -v`
  
Both commands will run all tests on backend, but second one will show tests that has been run with statuses.

For ranking of zipcodes, I think there is mistake with data where for every entry we have same zipcode, but code is implemented in way to support different
values and ranking.


There is two ways to run this application in your local environment:
  - You can checkout this project and position yourself in planted directory and run command `docker-compose up --build`
  - Pull two docker images from pulic dockerhub repositories:
    - docker pull `boradocker26/planted-project:frontend`
    - docker pull `boradocker26/planted-project:backend`
  - After that you can build these images with `docker image build` and give names as your choice.
  - Finally, you can run these built images with `docker run`
  
  
Frontend application wiil be available on port 3000 and backend will be on 80 port.
