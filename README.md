golang-rest-api for todolist

Steps for local setup

**``` go get github.com/vedadeepta/todo-list-rest-api ```**

**``` cd $GOPATH/src/github.com/vedadeepta/todo-list-rest-api ```**

**``` go build server.go ```**

**``` ./server ```**

Visit **localhost:3000**

To test **/create** endpoint send a **``` POST ```** request with the following structure
```
{
  ID: string,
  Name: string,
  Description: string
}

```
curl Example
```
 curl  -H "application/json" -d '{"ID": "1", "Name": "my 1st todo", "Description": "something 1"}' -X POST http://localhost:3000/create

```

To get all the todos send a **``` GET ```** request to **/todo** endpoint

curl Example

```
 curl -X GET http://localhost:3000/todos

```

To get a single todo send a **``` GET ```** request to **/todos/{id}** endpoint

```
 curl -H "application/json" -X GET http://localhost:3000/todos/2

```

To delete a todo send a **``` DELETE ```** request to **/todos/{id}** endpoint

```
 curl  -H "application/json" -X DELETE http://localhost:3000/todos/2
 
```

To update a todo send a **``` PUT ```** request to **/update** endpoint

```
 curl  -H "application/json" -d '{"ID": "2", "Name": "my 2nd todo"}' -X PUT http://localhost:3000/update

```
