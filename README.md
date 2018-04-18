# BART Go API

## Introduction

BART provides an API that works with XML, and more recently JSON.  This was originally a project to provide JSON output when it wasn't available from BART (they added it later).  It's now a project to provide "nicer" JSON output than that which comes from BART directly.

I originally wrote this in Node.js a while back and you can see the code for that [here](https://github.com/simonprickett/bartnodeapi).  There's a running example of the Node code [here](http://bart.crudworks.org/api).

## Running Locally

TODO - how to get up and running with dep

TODO - environment variables?

```
go run main.go
```

## Dependencies

* [dep dependency manager](https://github.com/golang/dep)
* [Gin web framework](https://gin-gonic.github.io/gin/)
* [Cron library](https://github.com/robfig/cron)