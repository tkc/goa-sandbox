#!/usr/bin/env bash

goagen bootstrap -d github.com/tkc/goa-sandbox/design
go build
./goa-sandbox