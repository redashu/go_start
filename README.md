# GO Lang --> The modern way of programing 

## Info about GO lang 
Way fast core level programing  language.  <br/>
Can be installed and run on every platform. <br/>
Its a fully compiled Language.  <br/>

## Study material 
you can start [Go](https://golang.org/) from this  <br/>
In case you don't have platform to run then use google playground to run the code

## Developers 
<ul>
<li>Robert Griesemer</li>
<li>Rob Pike</li>
<li> Ken Thompson </li>
</ul>

##  Public available 
Nov 2009 

# Compiler

GCCGO is the name of compiler 

## Platform 
I am using ubuntu 18.04  <br/>
I am using visual studio to code  <br/>

##  Installation 
```
sudo apt install golang-go 
```
###  Version check 
```
root@fire: go version 
go version go1.10.4 linux/amd64
```
## demo code and usage 
```
package main

import "fmt"

func main() {

	fmt.Println("Hello world , welcome to goLang !.")
}
```
## Running and compiling  the code 
```
fire@fire:$ go run hello.go 
Hello world , welcome to goLang !.
```
##  Only compiling  code 
```
fire@fire:~$ ls
hello.go  LICENSE  README.md
fire@fire:~$ go build  hello.go 
fire@fire:~$ ls
<b>hello</b>  hello.go  LICENSE  README.md
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
you can start [contribution](https://github.com/golang) by clicking and opening an issue.

## License
[MIT](https://choosealicense.com/licenses/mit/)
