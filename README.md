# buddhabrot-go

[![Go](https://github.com/ebeeton/buddhabrot-go/actions/workflows/go.yml/badge.svg)](https://github.com/ebeeton/buddhabrot-go/actions/workflows/go.yml)

A Buddhabrot plotter written as a Go, React, and TypeScript learning exercise.

![Buddhabrot image](/samples/sample.png)

This plot took about a minute on an eight-core machine.

## Usage

### Set the MySql Password

You'll need to the MySQL root password. Create a file in the root of the
repository called `.env` and set the contents to `DB_ROOT_PASSWORD=yourchoice`.

### Run It

In the same directory run `docker compose up`, which starts the web application
on [http://localhost:8000](http://localhost:8000).

### Request a Plot

The form on the left of the app defines the parameters that are used to plot an
image. The table on the right shows the plot history and links to images of
completed plots. Click the plot button to enqueue a plot.

## Under the Hood

The parameters used to plot the image are posted to an API running on port 3000,
and the API records them in the database. The row ID returned will be used to
obtain a PNG image of the plot after it is complete. Because the plotting is CPU
intensive, RabbitMQ is used to enqueue the plot request for a separate plotter
process so the API request doesn't block.

A sample plot request:

```json
{
    "sampleSize":100000000,
    "maxIterations":5000,
    "region": {
        "minReal":-2.0,
        "maxReal":1.6,
        "minImag":-2.0,
        "maxImag":2.0
    },
    "width":512,
    "height":512,
    "gradient": [{
            "color":"#000000",
            "position":0.0
        }, {
           "color":"#ff8000",
           "position":0.5
        }, {
           "color":"#ffff00",
           "position":0.75
        }, {
            "color":"#FFFFFF",
            "position":1.0
        }
    ],
    "dumpCounterFile": false
}
```

If put in a file called params.json, it could be posted with:

```shell
curl -Ss -d @params.json -H "Content-Type: application/json" \
    http://localhost:3000/api/plots
```

This is included in the samples directory in a file named plotrequest.sh along
with the JSON file above, so to test the plotting you can run:

```shell
./request.sh
```

You will get a response similiar to this; note the `Id` property.

```json
{
    "Id": 25,
    "Plot": {
        "SampleSize": 10000000,
    ...
```

### Getting Images

Given the `Id` from the previous step, you can do a get to `/api/plots/25`. If
the plotter process has completed the plot, this will return a PNG image.
Otherwise 404 is returned until the plot is complete. To do this with curl:

```shell
curl -Ss "http://localhost:3000/api/plots/25 > image.png
```

A script for this is also included in the samples directory and can be run as:

```shell
./getimage.sh 25 > image.png
```

### Parameters

For a thorough explanation of the Buddhabrot
[see Wikipedia's page here](https://en.wikipedia.org/wiki/Buddhabrot).

| Parameter | Description |
| --- | --- |
| `sampleSize` | The number of random sample points not in the Mandelbrot set. |
| `maxIterations` | The maximum number of orbits per point.
| `region` | Defines the complex plane region points are sampled from. |
| - `minReal` | Minimum real component for each point. |
| - `maxReal` | Maximum real component for each point. |
| - `minImag` | Minimum imaginary component for each point. |
| - `maxImag` | Maximum imaginary component for each point. |
| `width` | Width of plot image in pixels. |
| `height` | Height of plot image in pixels. |
| `gradient` | Array of "stops" which are used to color the plot. |
| - `color` | Hex color of the stop. |
| - `position` | Position of the stop in [0, 1]. |
| `dumpCounterFile` | If `true`, dump the orbit counts per pixel to a file called counter.txt in the log directory. |

