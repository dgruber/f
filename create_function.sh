#!/bin/sh
if [ "$1x" = "x" ]; then
   echo "args: upper / reverse"
   exit 1
fi 
pfs function create $1 --git-repo https://github.com/dgruber/f --artifact lxamd64/f --image $REGISTRY/$REGISTRY_USER/$1 --env "FUNCTION=$1" --verbose

