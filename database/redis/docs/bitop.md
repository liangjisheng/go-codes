# bit op

Perform a bitwise operation between multiple keys (containing string values) 
and store the result in the destination key.

The BITOP command supports four bitwise operations: AND, OR, XOR and NOT, 
thus the valid forms to call the command are:

BITOP AND destkey srckey1 srckey2 srckey3 ... srckeyN
BITOP OR destkey srckey1 srckey2 srckey3 ... srckeyN
BITOP XOR destkey srckey1 srckey2 srckey3 ... srckeyN
BITOP NOT destkey srckey

The result of the operation is always stored at destkey.

```redis
redis> SET key1 "foobar"
"OK"
redis> SET key2 "abcdef"
"OK"
redis> BITOP AND dest key1 key2
(integer) 6
redis> GET dest
"`bc`ab"
redis>
```