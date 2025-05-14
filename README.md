# Physical pin code generator

Generates pin codes for use with physical locks, where the wear of keys needs to be taken into account. If a single pin code is used, then it will quickly be obvious which keys are being used.

The optimal solution would be to have codes that distribute the wear as evenly as possible. This little program aims to help with generating these codes based on per-code length and expected use frequency.

## This is written in golang

```shell
$ go mod download
$ make run
```

## Storage

The input for needed codes and generated ones is stored in `codes.json`.

See [codes.example.json].

```json
{
  "keys": "0123456789", // all the allowed keys on your keypad
  "codes": [
    {
      "name": "lovely_wife", // human readable name
      "frequency": 14, // how often this code will be used per week (minimum 1)
      "length": 5, // the length of this code
      "code": ""  // the actual code - if empty, one will be generated
    },
    {
      "name": "meh_husband",
      "frequency": 14,
      "length": 5,
      "code": ""
    },
    {
      "name": "1st_child",
      "frequency": 7,
      "length": 4,
      "code": ""
    },
    {
      "name": "wife_parents",
      "frequency": 1,
      "length": 4,
      "code": ""
    }
  ]
}
```