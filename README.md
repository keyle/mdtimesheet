# mdtimesheet  [![Build Status](https://github.com/keyle/mdtimesheet/workflows/Go/badge.svg)]

`mdtimesheet` calculates time spent on projects, based on a markdown `plan` style file.

Consider the following format for recording DONE work and TODO in a markdown file:

```
[...]

### 2020-02-01-mon :0945-2150-120m

* read status, save last message id seen, per channel into the localstorage & memory
* massive lag when typing due to nick parsing [bug]
* reimpl nick completion

### 2020-02-02-tue :0900-1600-50m

* fix for messaging incl. private messaging is broadcasted to everyone
* reimpl harness testing with go backend [wip]

### 2020-02-03-wed :0800-1630-50m

* fix for we can still autcomplete self [bug]
* fix for we can still /message @self [bug]
* initdb tool and library

### @next 

- [other tasks...]

### @backlog

- [other tasks...]

### @consider

- [thoughts...]

```

`mdtimesheet` will parse the title H3 (###), find the date, and the time spent where

* `:` indicates the time tagging starts
* `0945` is military time for 9:45 AM
* `2150` is military time for 9:50 PM
* `-220m` is the amount of minutes to substract for lunch, breaks, etc.

### Sample Result

```
[...]

2020-02-01-mon :0945-2150-120m
	 2020-02-01 : 12h5m0s : 605 minutes

2020-02-02-tue :0900-1600-50m
	 2020-02-02 : 7h0m0s : 370 minutes

2020-02-03-wed :0800-1630-50m
	 2020-02-03 : 8h30m0s : 460 minutes
-----------------------------------------
Total minutes spent 15615 minutes
Total duration spent 260h15m0s
```

### Build

```
# go version go1.15.6 darwin/amd64
> go build .
```

### Usage

`./mdtimesheet ~/filepath/plan.md`

### License

Apache 2.0
