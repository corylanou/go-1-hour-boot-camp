
# Installing Go

## Windows

Download the latest installer from go.  [http://golang.org/dl/](http://golang.org/dl/)

## Mac

### Preferred

*Using Homebrew:*

```bash
brew update
brew install go
```

### Alternate

Download the latest installer from go.  [http://golang.org/dl/](http://golang.org/dl/)

## Linux

### Preferred

Download the latest installer from go.  [http://golang.org/dl/](http://golang.org/dl/)

### Alternate Install Methods - Not preferred

*Using APT*

```bash
sudo apt-get update
sudo apt-get install golang
```

*Note: This can leave you with an old version of go as they don't update
the package manager as timely as they should*

# Environment - GOPATH

The GOPATH environment variable specifies the location of your
workspace. It is likely the only environment variable you'll need to set
when developing Go code.

To get started, create a workspace directory and set GOPATH accordingly.
Your workspace can be located wherever you like, but we'll use $HOME/go
in this document. Note that this must not be the same path as your Go
installation.

*NOTE:* This is only necessary if the method of installation used above did NOT do this.

## Instructions for Linux or Mac

```bash
mkdir $HOME/go
export GOPATH=$HOME/go
```

## Instructions for Windows

From a command prompt:

```bash
mkdir "%USERPROFILE%\go"
```

Go to the Control Panel > System > Advanced Tab > Environment Variables.

Add a new *User* Variable (not a system variable)

Variable name: GOPATH
Variable value: %USERPROFILE%\go

*NOTE:* It's very likely you will need to reboot for this variable to take effect.

# Add Go Bin directory to PATH

For convenience, add the workspace's bin subdirectory to your PATH:

*NOTE:* If you used an installer, you typically do not need to do this step.


## Instructions for Linux or Mac

```bash
export PATH=$PATH:$GOPATH/bin
```

## Instructions for Windows

Go to the Control Panel > System > Advanced Tab > Environment Variables.

Under System variables, select "Path" from the box, and click "edit".  Add `;C:\Go\bin` to the value and click ok.


## Test your installation

Check that Go is installed correctly by building a simple program, as follows.

Create a file named `hello.go` and put the following program in it:

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```
Then run it with the go tool:

```bash
$ go run hello.go
hello, world
```

If you see the "hello, world" message then your Go installation is working.

*NOTE:* This was taken directly from [golang.org/doc/install](http://golang.org/doc/install)

# Revision Control Systems

Go has the ability to import remote packages via revision control systems with the `go get` command.  To ensure that you can retrieve any remote package, be sure to install the following rcs software to your system.

## Git

[Install Git](http://git-scm.com/book/en/Getting-Started-Installing-Git)

## Mercurial

[Install Mercurial](http://mercurial.selenic.com/wiki/Download)

## Subersion

[Install Subversion](https://subversion.apache.org/packages.html)

## Bazaar

[Install Bazaar](http://wiki.bazaar.canonical.com/Download)

# Go Tools - [Applies to Go 1.3 and prior]

`go vet` and `go doc` are now part of the go.tools sub repo:
[golang.org/doc/go1.2#go_tools_godoc](http://golang.org/doc/go1.2#go_tools_godoc)

To get `go vet` and `go doc`, from a console run:

```bash
go get code.google.com/p/go.tools/cmd/godoc
go get code.google.com/p/go.tools/cmd/vet
```

*Troubleshooting*

If you are on a mac and getting permission denied and used the package installer to install go, you may need to run this command to get the tools installed:

```bash
sudo chown -R $USER /usr/local/go
```
