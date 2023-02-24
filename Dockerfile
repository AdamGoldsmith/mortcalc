FROM docker.io/golang
COPY mortcalc.go mortcalc.go
COPY templates/main.html templates/
RUN go build mortcalc.go
CMD ["./mortcalc"]
EXPOSE 3001
