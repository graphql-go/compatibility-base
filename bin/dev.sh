#!/bin/bash

if [[ "$@" == "single" ]]; then
  go run cmd/internal/single_model.go
elif [[ "$@" == "multiple" ]]; then
  go run cmd/internal/multiple_model.go
fi
