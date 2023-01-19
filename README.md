# buddhabrot-go

[![Go](https://github.com/ebeeton/buddhabrot-go/actions/workflows/go.yml/badge.svg)](https://github.com/ebeeton/buddhabrot-go/actions/workflows/go.yml)

A Buddhabrot plotter written as a Go learning exercise.

![Buddhabrot image](/samples/sample.png)

# Usage

`go run .` starts an HTTP server on port 3000. The parameters used to plot the
image are posted in JSON, and a PNG image is written to the response.

```json
{
    "sampleSize":500000000,
    "maxIterations":1000,
    "region": {
        "minReal":-2.0,
        "maxReal":2.0,
        "minImag":-2.0,
        "maxImag":2.0
    },
    "width":2000,
    "height":2000,
    "gradient": [{
            "color":"#000000",
            "position":0.0
        }, {
           "color":"#00FF00",
           "position":0.25
        }, {
            "color":"#FFFFFF",
            "position":0.75
        }, {
            "color":"#FFFFFF",
            "position":1.0
        }]
}
```

If put in a file called params.json, it could be posted with:

```shell
curl -Ss -d @params.json -H "Content-Type: application/json" \
    http://localhost:3000
```

Note that the plotting is CPU intensive and blocks the request until complete.
