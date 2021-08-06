# NekoDNS
NekoDNS is a rudimentary DNS server

```
+-----------+
|  web ui   |
+-----+-----+
      |
      |
+-----v-----+
| apiserver |
+-----+-----+
      |
      |
+-----v-----+       +---------+
|   etcd    <-------+ coredns |
+-----------+       +---------+
```
