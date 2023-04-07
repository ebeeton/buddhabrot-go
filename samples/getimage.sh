#!/bin/bash
if [ -z "$1" ]
then
    echo "Please supply the plot ID."
    exit 1
fi

curl -Ss "http://localhost:3000/api/plots/$1"
