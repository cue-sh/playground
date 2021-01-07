#!/usr/bin/env bash

set -eux

go mod download
cd $(git rev-parse --show-toplevel)
path="cuelang.org/go"
version=$(go list -m -f={{.Version}} $path)

td=$(mktemp -d)
# trap "rm -rf $td" EXIT

pushd $td > /dev/null
modCache=$(go env GOMODCACHE)
if [ "$modCache" = "" ]
then
	modCache=${GOPATH%%:*}/pkg/mod
fi
unzip -q $modCache/cache/download/$path/@v/$version.zip
popd > /dev/null

regex='s+cuelang.org/go/internal+github.com/cue-sh/playground/internal/cuelang_org_go_internal+g'

for i in "" filetypes encoding third_party/yaml
do
	rsync -a --relative --delete $td/$path@$version/internal/./$i/ ./internal/cuelang_org_go_internal/
	find ./internal/cuelang_org_go_internal/$i -mindepth 1 -maxdepth 1 -type d -exec rm -rf {} +
done

find ./internal/cuelang_org_go_internal -name "*.go" -exec sed -i $regex {} +
find ./internal/cuelang_org_go_internal/ -name "*_test.go" -exec rm {} +
cp $td/$path@$version/LICENSE ./internal/cuelang_org_go_internal

go mod tidy
