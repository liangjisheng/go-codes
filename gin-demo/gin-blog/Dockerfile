FROM ubuntu:18.04

# WORKDIR $GOPATH/src/go-demos/gin-demo/gin-blog
# COPY . $GOPATH/src/go-demos/gin-demo/gin-blog
# RUN go build .
WORKDIR /app
COPY . /app
# COPY gin-blog /app/gin-blog

EXPOSE 8080
CMD ["./gin-blog"]