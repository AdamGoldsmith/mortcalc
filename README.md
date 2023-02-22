# mortcalc

Dabbling with Golang in a docker/podman container environment

## Build

```bash
podman build -t mortcalc .
```

## Run

```bash
podman run -d -p 8181:8081 mortcalc
```

## Test

```bash
curl http://localhost:8181
```

## Use

Point a browser at http://localhost:8181

## Reference

* Go webserver foundations based on [Rosie Hamilton's blog post](https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html)
* Writing [Web Apps in Go](https://go.dev/doc/articles/wiki/)
