# procedure

存储过程

```sql
drop procedure if exists my_test_procedure;
create procedure my_test_procedure(num int)
begin
  declare sum int default 0;
  declare i int default 0;
  set i = 0;
  while i<num do
    set sum=sum+1;
    set i=i+1;
  end while;
      
  select sum;
end;

call my_test_procedure(3);
drop procedure if exists my_test_procedure;
```
