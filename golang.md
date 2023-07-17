# Golang Environment

All implementations are built in Golang, a modern programming language. 

How to set up a golang enviroment? Follows these steps

1. Go to the directory you want to start your project

2. Create a module and name its domain and subdomain

```bash
go mod init domain.com/subdomain
```

3. A go.mod file will be created. Your external functions should be declared on a folder. Here, I've named this folder as `helper`

```bash
helper/mylib.go
...
main.go
go.mod
```

Then, say, you want to use Init() method of a struct List that is defined on mylib.go. For doing so, first go to main.go and call it by means of an instance of helper package, like this

```go
package main
import "exemplo.com/subdominio/helper"

func main() {
	var instance helper.List
	instance.Init(10)
}
```

5. Then, on mylib.go, you must invoke, on the first line, the helper package

```go
package helper
```

So, all methods, variables and structs declared on mylib.go will be visible for main.go
