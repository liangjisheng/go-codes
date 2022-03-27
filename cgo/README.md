# README

[article](https://blog.csdn.net/fuyuande/article/details/89178640)
[article](https://blog.csdn.net/fengfengdiandia/article/details/82747515)
[article](https://blog.csdn.net/fengfengdiandia/article/details/82748007)
[article](https://www.cnblogs.com/linguanh/p/8323487.html)
[article](http://blog.codeg.cn/post/blog/2016-04-20-golang-cgo/)
[linux-so](https://blog.csdn.net/qq_30549833/article/details/99714237)

开发注意事项：

1. 在注释和import”C”之间不能有空行
2. 使用C.CString函数转换GoString为CString时要手动释放该字符串。
3. CGO不支持使用变参的函数，例如printf,如果要使用的话，可以写个包裹函数m'yprintf,使用传参的方式调用。
4. Go支持使用//export导出函数给C使用，但是有一点需要注意就是不能在export导出的同一个文件里定义c函数，不然会出现

multiple definition of "xxx"编译错误，如果函数非常tiny的话，还有一个方法是使用static inline 来声明该函数
