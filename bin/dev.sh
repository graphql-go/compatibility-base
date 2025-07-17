#!/bin/bash

if [[ "$@" == "single" ]]; then
  go run cmd/internal/single/single_model.go
elif [[ "$@" == "multiple" ]]; then
  go run cmd/internal/multiple/multiple_model.go
fi
