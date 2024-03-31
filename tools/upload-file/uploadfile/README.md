# upload file

[upload](https://mp.weixin.qq.com/s/OHzXxfcBaf5RNT4dA38LCQ)

```shell
curl --location --request POST 'http://127.0.0.1:8080/uploadMulti' --form 'name="alice"' --form 'age="23"' --form 'file1=@"/Users/liangjisheng/Downloads/2021926-12542.jpeg"' --form 'file2=@"/Users/liangjisheng/Downloads/2021926-125420.jpeg"'

curl --location --request POST 'http://127.0.0.1:8080/uploadMulti' --form 'file1=@"/Users/liangjisheng/Downloads/2021926-12542.jpeg"'
```


