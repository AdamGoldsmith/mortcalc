# mortcalc

Dabbling with Golang in a docker/podman container environment

## Build

```bash
podman build -t mortcalc .
```

## Run

```bash
podman run -d -p 3001:3001 mortcalc
```

## Test

```bash
curl http://localhost:3001
```

## Use

Point a browser at http://localhost:3001

## Reference

* Go webserver foundations based on [Rosie Hamilton's blog post](https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html)
* Writing [Web Apps in Go](https://go.dev/doc/articles/wiki/)
