#!/bin/bash

export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
PATH1=${PATH%:${GOPATH}}
export PATH=${PATH1}:${GOPATH}
PATH1=${PATH%:${GOBIN}}
export PATH=${PATH1}:${GOBIN}
