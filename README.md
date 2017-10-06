# Gopulse

Simple employee engagement web tool

## Prerequisites

```shell
$ go get github.com/stretchr/gomniauth
$ go get github.com/clbanning/x2j
$ go get github.com/ugorji/go/codec
$ go get gopkg.in/mgo.v2/bson
$ go get gopkg.in/flosch/pongo2.v3
$ go get github.com/astaxie/beego
$ go get github.com/beego/bee
$ go get github.com/go-sql-driver/mysql
```

## Database

```mysql
CREATE TABLE `session` (
`session_key` char(64) NOT NULL,
`session_data` blob,
`session_expiry` int(11) unsigned NOT NULL,
PRIMARY KEY (`session_key`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
```

## Resources

  * https://github.com/stretchr/gomniauth
  * https://github.com/valyala/quicktemplate -
  * https://github.com/flosch/pongo2 ?
  * https://beego.me/quickstart (https://github.com/astaxie/beego)

## Hacks

  * https://github.com/golang/go/issues/19734
  * https://github.com/joiggama/beego-example
  * https://beego.me/docs/mvc/model/orm.md