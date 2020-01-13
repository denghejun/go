## go public packages: https://godoc.org/
## Install Go
https://golang.org/dl/

## Setup ENV GOPATH & GOROOT 
- export GOPATH=~/coding/go/
- export GOROOT=/usr/local/opt/go/libexec
- export PATH=$PATH:$GOPATH/bin
- export PATH=$PATH:$GOROOT/bin

## go install 
在go文件所在目录执行go install后：
- 如果go文件对应的package中包含main函数，则会在$GOPATH/bin目录生成对应平台的可执行文件。
- 如果go文件对应的package中不包含main函数，则会在$GOPATH/pkg目录下生成对应的.a后缀的package包。

## go install VS go build
Go build just compile the executable file and move it to destination.
Go install do a little more. It move the executable file to $GOPATH/bin and cache all non-main packages which imported to $GOPATH/pkg. the cache will be use in the next compile if it not changed yet.

## go test
test文件以"_test.go"结尾，在其目录下执行`go test`,运行单元测试.

## go run xxx.go
go文件中必须包含main函数, 执行go文件中的main函数。



