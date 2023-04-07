# buddhabrot-go

[![Go](https://github.com/ebeeton/buddhabrot-go/actions/workflows/go.yml/badge.svg)](https://github.com/ebeeton/buddhabrot-go/actions/workflows/go.yml)

A Buddhabrot plotter written as a Go learning exercise.

![Buddhabrot image](/samples/sample.png)

This plot took about eight minutes on an eight-core machine using the parameters
below.

## Usage

### Set the MySql Password

You'll need to the MySQL root password. Create a file in the root of the
repository called `.env` and set the contents to `DB_ROOT_PASSWORD=yourchoice`.

### Run It

In the same directory run `docker compose up --build`, which starts the API on
`http://localhost:3000`. After all the containers are running you can request a
plot.

### Request a Plot

The parameters used to plot the image are posted in JSON by clients 
to '/api/plots', and the API records them in the database. The row ID and 
parameters are returned with HTTP status 201. The ID will be used to obtain a
PNG image of the plot after it is complete. Because the plotting is CPU
intensive, RabbitMQ is used to enqueue the plot request for a separate plotter
process so the API request doesn't block.

A sample plot request:

```json
{
    "sampleSize":1000000000,
    "maxIterations":5000,
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

Given the `Id` from the previous step, you can do a get to `/api/plots/25'`. If
the plotter process has completed the plot, this will return a PNG image.
Otherwise 404 is returned until the plot is complete. To do this with curl:

```shell
curl -Ss "http://localhost:3000/api/plots/25 > image.png
```

A script for this is also included in the samples directory and can be run as:

```shell
./getimage.sh 25 > image.png
```
