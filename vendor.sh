#!/bin/bash

export LANG=C

__dirname=`which dirname`
__basename=`which basename`

LOG="1"

log_e() {
  if [ "x$LOG" != "x" ]; then
    echo -e "[LOG] $1"
  fi
}

log_ln() {
  if [ "x$LOG" != "x" ]; then
    echo ""
  fi
}

log_n() {
  if [ "x$LOG" != "x" ]; then
    echo -n $1
  fi
}

log() {
  if [ "x$LOG" != "x" ]; then
    echo "[LOG] $1"
  fi
}

__execpath(){
  if [ ! -e "$1" ]; then
    return 1
  fi
  pushd `${__dirname} $1` >/dev/null 2>&1
  [ $? -eq 1 ] && exit 1
  _dir=`pwd`
  popd >/dev/null 2>&1
  echo ${_dir}
}
__script_dir=`__execpath $0`

log ${__script_dir}

goget() {
  vcs=$1
  pkg=$2
  rev=$3

  pkg_url=https://${pkg}
  target_dir=${__script_dir}/src/${pkg}

  if [ "x${rev}" == "x" ]; then
    log "Getting dependency (HEAD) -> ${vcs}:${pkg}"
  else
    log "Getting dependency        -> ${vcs}:${pkg}@${rev}"
  fi

  if [ ! -d ${target_dir} ]; then
    log "Start install  -> GOPATH=${__script_dir} go get $pkg"
    GOPATH=${__script_dir} go get -u $pkg
    log "Finish install -> GOPATH=${__script_dir} go get $pkg"
  fi

  case $vcs in
    git)
      pushd ${target_dir} >/dev/null 2>&1
      log "git reset --quiet --hard $rev"
      git reset --quiet --hard $rev
      git --no-pager log -n 1 --pretty=oneline
      popd >/dev/null 2>&1
      ;;
  esac
  log_ln
}

# goa 
goget git github.com/goadesign/goa/goagen     35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa            35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/middleware 35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/dslengine  35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/client     35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/cors       35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/design     35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/encoding   35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/goatest    35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/logging    35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/uuid       35e703a31f2552f5847cf9ab06db07e84b390d89
goget git github.com/goadesign/goa/version    35e703a31f2552f5847cf9ab06db07e84b390d89
goget git golang.org/x/net/context

# goa dependency 3rd party
goget git github.com/armon/go-metrics         93f237eba9b0602f3e73710416558854a81d9337
goget git github.com/dimfeld/httptreemux      86f7c217d9043ebc6adfd8e2ed04a0bb1e1db651

GOPATH=${__script_dir} go get -u github.com/jteeuwen/go-bindata/...

# other tools
goget git golang.org/x/tools/cmd/cover
goget git github.com/smartystreets/goconvey
goget git github.com/nsf/gocode
goget git github.com/rogpeppe/godef
goget git github.com/zmb3/gogetdoc
goget git github.com/lukehoban/go-outline
goget git sourcegraph.com/sqs/goreturns
goget git golang.org/x/tools/cmd/gorename
goget git github.com/tpng/gopkgs
goget git github.com/newhook/go-symbols
goget git golang.org/x/tools/cmd/guru
goget git golang.org/x/tools/cmd/godoc
GOPATH=${__script_dir} go get -u github.com/cweill/gotests/...

# global
go get -u github.com/golang/lint/golint
