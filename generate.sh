#!/usr/bin/env bash

protoc ./proto/*/*.proto --jsonschema_out=./json --jsonschema_opt=file_extension=jsonschema --jsonschema_opt=json_fieldnames --jsonschema_opt=enforce_oneof -I ./proto