#!/bin/bash
curl -Ss -d @request.json -H "Content-Type: application/json" \
    http://localhost:3000
