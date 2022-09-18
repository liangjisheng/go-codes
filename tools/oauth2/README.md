# oauth2.0

[official](https://github.com/golang/oauth2)

[article](https://blog.csdn.net/qq_19018277/article/details/104935403)
[github](https://github.com/Lxy417165709/golang_github_Oauth)

[oauth1](http://www.ruanyifeng.com/blog/2019/04/oauth_design.html)
[oauth2](https://www.ruanyifeng.com/blog/2019/04/oauth-grant-types.html)
[oauth3](https://www.ruanyifeng.com/blog/2019/04/github-oauth.html)

第三方网站登录 Github 的原理如下:

1. 用户在第三方网站点击 Github授权登录 后，浏览器跳转到 Github 登录页面
2. 用户在 Github 登录成功后，浏览器重定向回第三方网站。此时浏览器会携带一个 code
3. 第三方网站服务器通过 code 向 Github 索取 token
4. Github 返回 token
5. 第三方网站服务器收到 token 后，就可以通过 token 获取用户信息了
