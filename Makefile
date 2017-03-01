# Makefile - base
#

__PWD=$(shell pwd)
__BIN=$(__PWD)/bin
__PKG=$(__PWD)/pkg
__DIST=$(__PWD)/dist

__go=$(shell which go)


__PROG=$(__DIST)/$(__PROG_NAME)

__PATH=$(__BIN):$$PATH

TARGET=$(target)
__PROG_NAME=goa-as-design-for-api
__MAIN=$(__GOPATH_THIS_PROJECT)/$(TARGET)/*.go

__GOPATH=$(__PWD)
__GITHUB_THIS_PROJECT_ORG=github.com/syama666
__GITHUB_THIS_PROJECT=$(__GITHUB_THIS_PROJECT_ORG)/$(__PROG_NAME)
__GOPATH_THIS_PROJECT=src/$(__GITHUB_THIS_PROJECT)


all:goa-gen api

clean:
	rm -f $(__DIST)/$(__PROG_NAME)
	rm -rf $(__PKG)

deps:
	$(__PWD)/vendor.sh

goa-gen: goa-gen-cmd goa-gen-app goa-gen-swagger

goa-gen-cmd:
	cd $(__GOPATH)/src/$(__GITHUB_THIS_PROJECT) && GOPATH=$(__GOPATH) $(__GOPATH)/bin/goagen main -d $(__GITHUB_THIS_PROJECT)/design -o ./controller && rm ./controller/main.go

goa-gen-app:
	cd $(__GOPATH)/src/$(__GITHUB_THIS_PROJECT) && GOPATH=$(__GOPATH) $(__GOPATH)/bin/goagen app -d $(__GITHUB_THIS_PROJECT)/design -o ./gen/

goa-gen-swagger:
	cd $(__GOPATH)/src/$(__GITHUB_THIS_PROJECT) && GOPATH=$(__GOPATH) $(__GOPATH)/bin/goagen swagger -d $(__GITHUB_THIS_PROJECT)/design -o ./gen/
	cd $(__GOPATH)/src/$(__GITHUB_THIS_PROJECT) && GOPATH=$(__GOPATH) $(__GOPATH)/bin/go-bindata -pkg swagger -o ./gen/swagger/swagger-gen.go ./gen/swagger/swagger.json

setup: deps

api:
	GOPATH=$(__GOPATH) $(__go) build -o $(__PROG) src/$(__GITHUB_THIS_PROJECT)/main.go

run:
	THIS_PROJECT_DEBUG=true $(__PROG)

help:
	@echo "Target:"
	@echo "\t$$ make setup"
	@echo "\t$$ make all"
	@echo "\t$$ make run"
	@echo ""

.PHONY: all clean setup api deps
