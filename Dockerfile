FROM docker.io/golang
# COPY server.go server.go
# RUN go build server.go
COPY template.go template.go
COPY homepage.html homepage.html
RUN go build template.go
CMD ["./template"]
EXPOSE 8181
