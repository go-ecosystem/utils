# utils

![Go](https://github.com/go-ecosystem/utils/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/go-ecosystem/utils/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ecosystem/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ecosystem/utils)](https://goreportcard.com/report/github.com/go-ecosystem/utils)
[![Release](https://img.shields.io/github/release/go-ecosystem/utils.svg)](https://github.com/go-ecosystem/utils/releases)

> go utils

## Install

```shell
go get github.com/go-ecosystem/utils/v2
```

## Packages

- array
  - Contains
  - HasDuplicateItem
- convert
  - StringToBytes
  - BytesToString
  - JSONToMap
  - MapToJSON
- crypto
  - EncryptPwd
  - ComparePwd
- file
  - Exists
  - Exist
  - Mode
  - WriteStringToFile
  - AppendStringToFile
  - GetDirList
  - GetDirListWithFilter
  - RecreateDir
  - GetFilepaths
  - GetFiles
- flag
  - IsTesting
- http
  - ClientIP
- json
  - ToString
  - FromString
  - ToMap
  - Convert
- log
  - E
  - W
  - H
  - I
  - V
- middleware
  - interceptor
    - Recovery
  - Cors
  - Recovery
- net
  - GetIP
- os
  - RunBashCommand
  - RunCommand
  - RunCommandWithStdin
- response
  - UnknownError
  - DBConnectError
  - DBOperationError
  - RequestParamError
  - TokenInvalidError
  - PermissionDeniedError
  - JSON
  - OK
  - Err
  - Abort
- strings
  - Capitalize
  - IsCapitalize
  - SplitToChunks
  - Rand
  - IsNumber
  - RemoveEmpties
  - Remove
- version
  - Stringify
  - StringifyWithOps

## Usage

see xxx_test.go file.
