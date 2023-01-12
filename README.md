# buddhabrot-go

A Buddhabrot plotter written as a Go learning exercise.

![Buddhabrot image](/assets/images/sample.png)

# Usage

`go run .` starts an HTTP server on port 3000. The parameters used to plot the
image are posted in JSON, and a PNG image is written to the response.

```json
{
    "red": {
        "sampleSize":100000000,
        "maxSampleIterations":1000,
        "maxIterations":1000
    },
    "green": {
        "sampleSize":100000001,
        "maxSampleIterations":1001,
        "maxIterations":1001
    },
    "blue": {
        "sampleSize":100000002,
        "maxSampleIterations":1002,
        "maxIterations":1002
    },
    "region": {
        "minReal":-2.0,
        "maxReal":2.0,
        "minImag":-2.0,
        "maxImag":2.0
    },
    "width":2000,
    "height":2000
}
```

If put in a file called params.json, it could be posted with:

```shell
curl -Ss -d @params.json -H "Content-Type: application/json" \
    http://localhost:3000
```

Note that the plotting is CPU intensive and blocks the request until complete.
