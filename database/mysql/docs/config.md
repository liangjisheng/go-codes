# mysqld.cnf

mysqlx-bind-address 作用

```text
mysqlx-bind-address 是MySQL服务器用来监听基于X Protocol的连接的地址，X Protocol是MySQL服务器的一种新协议，
可以通过MySQL Shell或其它的客户端软件使用。bind-address是MySQL服务器用来监听传统的MySQL协议的连接的地址。
如果您配置了mysqlx-bind-address，那么基于X Protocol的连接只能通过指定的地址连接到MySQL服务器。
如果您没有指定mysqlx-bind-address，则默认为0.0.0.0，意味着MySQL服务器会监听所有地址上的X Protocol连接。同理，
如果您配置了bind-address，那么传统的MySQL协议连接只能通过指定的地址连接到MySQL服务器。如果您没有指定bind-address，
则默认为0.0.0.0，意味着MySQL服务器会监听所有地址上的MySQL协议连接。总的来说，mysqlx-bind-address和bind-address的作用是类似的，
都是用来指定MySQL服务器监听连接的地址的。不同的是，mysqlx-bind-address用于监听基于X Protocol的连接，
而bind-address用于监听传统的MySQL协议连接。
```
