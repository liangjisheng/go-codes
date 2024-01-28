# full index

[全文索引](https://juejin.cn/post/6989871497040887845)

```sql
-- CREATE FULLTEXT INDEX full_index_name ON tokens(symbol);

-- SELECT * FROM tokens WHERE MATCH ( symbol ) AGAINST ( 'xxx' );
-- SELECT * FROM tokens WHERE MATCH ( symbol ) AGAINST ( 'qqq' );
-- SELECT count(IF(MATCH ( symbol ) against ( 'T22' ), 1, NULL )) AS count FROM `tokens`;
-- SELECT *, MATCH ( symbol ) against ( 'xxx' ) AS Relevance FROM tokens;

-- select * from tokens where MATCH(symbol) AGAINST('+T22' IN BOOLEAN MODE);
-- select * from tokens where MATCH(symbol) AGAINST('t*' IN BOOLEAN MODE);

-- SELECT * FROM `tokens` WHERE MATCH ( symbol ) AGAINST ( 'T*' IN BOOLEAN MODE );
SELECT * FROM `tokens` WHERE MATCH ( symbol ) AGAINST ( 'aaa' IN BOOLEAN MODE );
```
