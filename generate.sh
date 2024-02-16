#!/usr/bin/env bash

protoc ./proto/*/*.proto --jsonschema_out=./json --jsonschema_opt=file_extension=jsonschema -I ./proto