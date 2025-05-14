# Physical pin code generator

Generates pin codes for use with physical locks, where the wear of keys needs to be taken into account. If a single pin code is used, then it will quickly be obvious which keys are being used.

The optimal solution would be to have codes that distribute the wear as evenly as possible. This little program aims to help with generating these codes based on per-code length and expected use frequency.

## This is written in golang

```shell
$ go mod download
$ make run
```
