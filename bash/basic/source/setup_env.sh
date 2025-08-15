#!/bin/bash

export ENV=production
echo "Setting up $ENV environment"

calculate() {
    echo $(($1 + $2))
}
