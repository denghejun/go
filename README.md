## This is A Standard GOPATH Project Workspace Structure
http://docscn.studygolang.com/doc/code.html

src: Go source files directory

pkg: Go packages directory

bin: Go executable file directory

#### Step1: Checkout Current Project to Your Local Directory
```git clone https://github.com/denghejun/go.git YourLocalDirectoryPath```

#### Step2: Set/Update Your $GOPATH ENV
```export GOPATH=YourLocalDirectoryPath```

```export PATH=$PATH:$GOPATH/bin```

#### Step3: Go Install Any Project
```cd $GOPATH/src/github.com/hakuna/01.helloworld/```

```go install```

```01.helloworld```
