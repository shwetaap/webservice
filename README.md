# Webserver Application

This is Webserver application implementing REST API calls to manipulate Objects

---
## API Specification

The sample http request looks like this

GET "/objects/{bucket}/{objectID}"

PUT "/objects/{bucket}/{objectID}" -d '{data":"obj"}

DELETE "/objects/{bucket}/{objectID}"

Data is stored in memory using the Object Strut. These Objects are distributed and stored in multiple buckets

---
## Getting Started
1. A user needs to assign a port to start the webserver. This can be done by setting the BIND_PORT environment variable
> export BIND_PORT=8080

2. Start the server
  * > go run server.go
  * > ./server.go
3. The webserver starts. In another terminal send the curl commands for GET/PUT/DELETE requests
4.  Add/Modify data
> curl localhost:8080/objects/1/2 -X PUT -d '{"data":"sec"}' -v

5. Get data
> curl localhost:8080/objects/1/2

6. Delete Data
> curl localhost:3000/objects/0/2 -X DELETE

---
## Unit Test Execution
Unit tests can be executed using the following command
> go test -v
