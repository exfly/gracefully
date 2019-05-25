# Gracefully [![Build Status](https://travis-ci.com/exfly/gracefully.svg?branch=master)](https://travis-ci.com/exfly/gracefully)

go app gracefully shutdown utils. This project is split from k8s, I want to do some progress here.

## Getting Started

```go
var stopCh = gracefully.SetupSignalHandler()
go func() {
    for {
        log.Println("alive...")
        time.Sleep(1 * time.Second)
    }
}()
<-stopCh
log.Println("gracefully shutdown")
```

### Installing

```sh
go get -u github.com/exfly/gracefully
```

## Deployment

This project is forked form k8s.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](github.com/exfly/gracefully/tags).

## Authors

- **exfly** - _Initial work_ - [exfly](https://github.com/exfly)

## License

This project is licensed under the Apache v2 License - see the [LICENSE.md](LICENSE) file for details
