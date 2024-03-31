build:
	go build
	docker build -t gin-blog-docker .
runBlog:
	docker run --link mysql:mysql -p 8080:8080 gin-blog-docker
runMysql:
	docker run --name mysql \
		-p 3306:3306 \
		-e MYSQL_ROOT_PASSWORD=root \
		-v /root/gopath/src/go-demos/gin-demo/gin-blog/data/docker-mysql:/var/lib/mysql \
		-d mysql