# set bit

set bit
Sets or clears the bit at offset in the string value stored at key.
The bit is either set or cleared depending on value, which can be either 0 or 1.

```redis
> setbit bitmapsarestrings 2 1
> setbit bitmapsarestrings 3 1
> setbit bitmapsarestrings 5 1
> setbit bitmapsarestrings 10 1
> setbit bitmapsarestrings 11 1
> setbit bitmapsarestrings 14 1
> get bitmapsarestrings
"42"
```
