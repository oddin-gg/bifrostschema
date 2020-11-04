#!/usr/bin/env bash

protoc ./proto/*/*.proto --jsonschema_out=./json -I ./proto