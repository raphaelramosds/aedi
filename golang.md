# Golang Environment

All implementations are built in Golang, a modern programming language. 

How to set up a golang enviroment? Follows these steps

1. Go to the directory you want to start your project

2. Create a module and name its domain and subdomain

```bash
go mod init domain.com/subdomain
```

3. A go.mod file will be created. Your external functions should be declared on helper folder, as the file structure below suggests

```bash
helper/arraylist2.go
helper/linkedlist2.go
...
main.go
go.mod
```

Then, say, you want to use Init() method of ArrayList and the struct ArrayList is defined on arraylist2.go. First go to main.go and call them by means of a instance of helper package, like this

```go
package main
import "exemplo.com/subdominio/helper"

func main() {
	var arraylist helper.ArrayList
	arraylist.Init(10)
}
```

5. Then, on arraylist2.go, you must invoke, on the first line, the helper package

```go
package helper
```

So, all methods, variables and structs declared on arraylist2.go will be visible for main.go
