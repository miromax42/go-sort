# go-sort
Sorter for txt files
## Type of lines
### Line type
>rule: field are splited by llast number sequence with leading zeros
Line filter fields:
```txt
[prefix][number][comment]
```
Example1:
```txt
a1a2a034b1b2b3
```
* `a1a2a` - group id
* `34` - number
*  `b1b2b3` - comment

Example2:
```txt
a01a2a034b1b2b3
```
* `a01a2a` - group id
* `34` - number
*  `b1b2b3` - comment

Example2:
```txt
a01a2a034b1b02b3
```
* `a01a2a034b1b` - group id
* `2` - number
*  `b3` - comment

### List type
> rule: starts with `--`
```txt
--exampleListName
--list234
--234
```

### Sublist type
> rule: starts with ` -`
```txt
 -1May
 -last
 -01.02.34
```

## FullExample
>command go-sort sort test.txt

file:
```txt
--l1
 -010203
e1e02e3
b03e3
c03e345
 -020203
e1e03e3
b04e3
c02e345
--l2
e1e01b7
b01e7
c0999e7
```
output:
```txt
| TYPE | NUMBER | COMMENT | REGISTRY |  DATE  |
+------+--------+---------+----------+--------+
| b    |      1 |   e7    | l2       |     NO |
+      +--------+---------+----------+--------+
|      |      3 |   e3    | l1       | 010203 |
+      +--------+---------+----------+--------+
|      |      4 |   e3    | l1       | 020203 |
+------+--------+---------+----------+--------+
| c    |      2 |  e345   | l1       | 020203 |
+      +--------+---------+----------+--------+
|      |      3 |  e345   | l1       | 010203 |
+      +--------+---------+----------+--------+
|      |    999 |   e7    | l2       |     NO |
+------+--------+---------+----------+--------+
| e1e  |      1 |   b7    | l2       |     NO |
+      +--------+---------+----------+--------+
|      |      2 |   e3    | l1       | 010203 |
+      +--------+---------+----------+--------+
|      |      3 |   e3    | l1       | 020203 |
+------+--------+---------+----------+--------+
|                            TOTAL   |   9    |
+------+--------+---------+----------+--------+
```

