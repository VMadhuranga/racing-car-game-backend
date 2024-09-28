#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
goose postgres $POSTGRES_URI up
