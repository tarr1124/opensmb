OpenSmb
============

## For what?
File sharing path from Windows users is troublesome for Mac users.
Mac user can open file/directory with CLI.

## Installation

```bash
$ go get -u github.com/tarr1124/opensmb
```

## Usage

```
$ opensmb '\\shared\Tech\開発部\全社員共有\'
This target path has been mounted.
open /Volumes/ec7d974583da4923cc7fda88d5bebb3db441fccc/全社員共有

$ opensmb '\\rfs\Tech\開発部\全社員共有\testing.xlsx'
This target path has been mounted.
open /Volumes/9a27e53f84549297b39569f87a6fda7315d130c1/testing.xlsx
```

## Author

[tarr1124](https://github.com/tarr1124)
