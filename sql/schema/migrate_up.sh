#!/bin/zsh
cd /Users/cameronharro/workspace/github.com/cameronharro/bootdev-aggregator-go/sql/schema
goose postgres $(cat ~/.gatorconfig.json | jq '.db_url' | xargs echo) $1
cd -
