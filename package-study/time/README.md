# time

[你确信你了解时间吗](https://coolshell.cn/articles/5075.html)
[关于闰秒](https://coolshell.cn/articles/7804.html)
[RFC 3339](https://tools.ietf.org/html/rfc3339)

## GMT

GMT(Greenwich Mean Time), 格林威治平时(也称格林威治时间)
它规定太阳每天经过位于英国伦敦郊区的皇家格林威治天文台的时间为中午12点
1972年之前，格林威治时间（GMT）一直是世界时间的标准。1972年之后，GMT 不再是一个时间标准了

## UTC

UTC(Coodinated Universal Time), 协调世界时，又称世界统一时间、世界标准时间、国际协调时间
由于英文（CUT）和法文（TUC）的缩写不同，作为妥协，简称UTC
UTC 是现在全球通用的时间标准，全球各地都同意将各自的时间进行同步协调。
UTC 时间是经过平均太阳时（以格林威治时间GMT为准）、地轴运动修正后的新时标以及以秒为单位的国际原子时所综合精算而成

UTC 由两部分构成：
原子时间（TAI, International Atomic Time）:
结合了全球400个所有的原子钟而得到的时间，它决定了我们每个人的钟表中，时间流动的速度。

世界时间（UT, Universal Time）:
也称天文时间，或太阳时，他的依据是地球的自转，我们用它来确定多少原子时，对应于一个地球日的时间长度

UTC的历史
1960年，国际无线电咨询委员会规范统一了 UTC 的概念，并在次年投入实际使用。
“Coordinated Universal Time”这个名字则在1967年才被正式采纳。
1967年以前， UTC被数次调整过，原因是要使用闰秒（leap second）来将 UTC 与地球自转时间进行统一

GMT vs UTC
GMT是前世界标准时，UTC是现世界标准时。
UTC 比 GMT更精准，以原子时计时，适应现代社会的精确计时。
但在不需要精确到秒的情况下，二者可以视为等同。
每年格林尼治天文台会发调时信息，基于UTC

## 使用注意

- 一定要使用 time.Time 和 time.Duration  这两个类型
- 在命令行上，flag 通过 time.ParseDuration 支持了 time.Duration
- JSON 中的 encoding/json 中也可以把time.Time 编码成 RFC 3339 的格式
- 数据库使用的 database/sql 也支持把 DATATIME 或 TIMESTAMP 类型转成 time.Time
- YAML 也可以使用 gopkg.in/yaml.v2 支持 time.Time 、time.Duration 和 RFC 3339 格式
- 如果要和第三方交互，实在没有办法也请使用 RFC 3339 的格式
- 如果要做全球化跨时区的应用，一定要把所有服务器和时间全部使用 UTC 时间
