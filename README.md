# BadgerDB [![Go Reference](https://pkg.go.dev/badge/github.com/dgraph-io/badger/v3.svg)](https://pkg.go.dev/github.com/dgraph-io/badger/v3) [![Go Report Card](https://goreportcard.com/badge/github.com/dgraph-io/badger)](https://goreportcard.com/report/github.com/dgraph-io/badger) [![Sourcegraph](https://sourcegraph.com/github.com/dgraph-io/badger/-/badge.svg)](https://sourcegraph.com/github.com/dgraph-io/badger?badge) [![Build Status](https://teamcity.dgraph.io/guestAuth/app/rest/builds/buildType:(id:Badger_UnitTests)/statusIcon.svg)](https://teamcity.dgraph.io/viewLog.html?buildTypeId=Badger_UnitTests&buildId=lastFinished&guest=1) ![Appveyor](https://ci.appveyor.com/api/projects/status/github/dgraph-io/badger?branch=master&svg=true) [![Coverage Status](https://coveralls.io/repos/github/dgraph-io/badger/badge.svg?branch=master)](https://coveralls.io/github/dgraph-io/badger?branch=master)

![Badger mascot](images/diggy-shadow.png)

（注：源码分析文档位于badger/myDocs中，本仓库是本人个人在学习badgerDB实现时自己魔改的注释版）

BadgerDB is an embeddable, persistent and fast key-value (KV) database written
in pure Go. It is the underlying database for [Dgraph](https://dgraph.io), a
fast, distributed graph database. It's meant to be a performant alternative to
non-Go-based key-value stores like RocksDB.
