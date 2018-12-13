#!/bin/sh
pfs function create upper --git-repo https://github.com/dgruber/f --artifact lxamd64/f --image $REGISTRY/$REGISTRY_USER/upper --env "FUNCTION=upper" --verbose

