#!/bin/bash
curl -Ss -d @params.json -H "Content-Type: application/json" \
    http://localhost:3000
