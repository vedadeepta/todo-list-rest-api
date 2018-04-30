golang-rest-api for todolist

Run **``` go get github.com/vedadeepta/todo-list-rest-api ```**

To run in local:
**``` cd $GOPATH/src/github.com/vedadeepta/todo-list-rest-api ```**
**``` go build server.go ```**
**``` ./server ```**

Visit **localhost:3000**

To test **/create** endpoint send a **``` POST ```** requet with the following structure
```
{
  ID: string,
  Name: string,
  Description: string
}

```
cUrl Example
```
 curl  -H "application/json" -d '{"ID": "1", "Name": "my 1st todo", "Description": "something 1"}' -X POST http://localhost:3000/create

```
