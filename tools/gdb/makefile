main:
	# ldflags 加入调试信息, gcflags 禁止编译器优化
	go build -o main -ldflags=-compressdwarf=false -gcflags "-N -l"	

clean:
	rm -f main
	