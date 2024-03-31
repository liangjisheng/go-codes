FROM    hub.c.163.com/library/golang:latest

WORKDIR /data
COPY main.go /data
RUN cd /data && go build main.go

EXPOSE 8080

ENTRYPOINT ["/data/main"]
