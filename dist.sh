#!/usr/bin/env bash
set -eux

# dist.sh is used by netlify when deploying the CUE playground within the wider
# cuelang.org site. It will, therefore, be run with a working directory that is
# within the module cache. Given we need to run npm install etc, this will fail
# because in this context all files/directories are read-only. Hence we make a
# copy of "ourselves" to a temp directory, make that writable, then run through
# all our dist steps.
#
# This script expects the following environment variables to have been set:
#
# * GOBIN - the target for serverless functions
# * NETLIFY_BUILD_BASE - the root of the netlify build, within which there will
#   be a cache directory
# * CUELANG_ORG_DIST - the directory into which we should run dist

if [ "${NETLIFY:-}" != "true" ]
then
	echo "Only intended to be run on Netlify"
	exit 1
fi

# cd to the directory containing the script
cd "$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

if [ "${BRANCH:-}" == "tip" ]
then
	 # We need to update our dependencies (as the main module)
	 # to the tip of CUE
	 go get -d cuelang.org/go@master
	 bash ./_scripts/revendorToolsInternal.sh
	 go mod tidy
	 go generate $(go list ./... | grep -v cuelang_org_go_internal)
fi

# Use the cache of playground node_modules
mkdir -p $NETLIFY_BUILD_BASE/cache/playground_node_modules
rsync -a $NETLIFY_BUILD_BASE/cache/playground_node_modules/ node_modules
npm install
rsync -a node_modules/ $NETLIFY_BUILD_BASE/cache/playground_node_modules

# Dist
echo "Install serverless functions to $GOBIN"
go install -tags netlify github.com/cue-sh/playground/functions/snippets

echo "Building WASM backend"
GOOS=js GOARCH=wasm go build -o main.wasm

echo "Running dist into $CUELANG_ORG_DIST"
npm run dist
