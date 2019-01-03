timetracker
===========

[![Build Status](https://travis-ci.org/martinohmann/timetracker.svg)](https://travis-ci.org/martinohmann/timetracker)
[![codecov](https://codecov.io/gh/martinohmann/timetracker/branch/master/graph/badge.svg)](https://codecov.io/gh/martinohmann/timetracker)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinohmann/timetracker)](https://goreportcard.com/report/github.com/martinohmann/timetracker)
[![GoDoc](https://godoc.org/github.com/martinohmann/timetracker?status.svg)](https://godoc.org/github.com/martinohmann/timetracker)

Simple CLI tool to track time for different tasks. Can be queried for total time spent within different periods (e.g. by week, month, year). Uses SQLite as storage.

Installation
------------

```sh
go get -u github.com/martinohmann/timetracker
cd $GOPATH/src/github.com/martinohmann/timetracker
make install
```

Usage
-----

```sh
$ timetracker start foo
+----+-----+---------------------+------+----------+
| ID | Tag | Start               | End  | Duration |
+----+-----+---------------------+------+----------+
|  1 | foo | 2019/01/02 17:43:29 | open |       0s |
+----+-----+---------------------+------+----------+
interval with tag "foo" started

$ timetracker stop foo
+----+-----+---------------------+---------------------+----------+
| ID | Tag | Start               | End                 | Duration |
+----+-----+---------------------+---------------------+----------+
|  1 | foo | 2019/01/02 17:43:29 | 2019/01/02 17:43:36 |       7s |
+----+-----+---------------------+---------------------+----------+
interval with ID 1 closed

$ timetracker start bar
+----+-----+---------------------+------+----------+
| ID | Tag | Start               | End  | Duration |
+----+-----+---------------------+------+----------+
|  2 | bar | 2019/01/02 17:43:43 | open |       0s |
+----+-----+---------------------+------+----------+
interval with tag "bar" started

$ timetracker show day
All intervals between 2019/01/02 00:00:00 and 2019/01/03 00:00:00
+----+-----+---------------------+---------------------+----------+
| ID | Tag | Start               | End                 | Duration |
+----+-----+---------------------+---------------------+----------+
|  1 | foo | 2019/01/02 17:43:29 | 2019/01/02 17:43:36 |       7s |
|  2 | bar | 2019/01/02 17:43:43 | open                |      20s |
+----+-----+---------------------+---------------------+----------+
|                                                Total |      27s |
+----+-----+---------------------+---------------------+----------+
```

Run `timetracker help` for all available subcommands and options.

License
-------

The source code of this is released under the MIT License. See the bundled LICENSE
file for details.
