# go nessus

go client for the nessus  

https://developer.tenable.com/reference#read-the-docs

[![GitHub license](https://img.shields.io/github/license/xanzy/go-gitlab.svg)](https://github.com/autonubil/go-nessus/blob/master/LICENSE)
[![Sourcegraph](https://sourcegraph.com/github.com/autonubil/go-nessus/-/badge.svg)](https://sourcegraph.com/github.com/autonubil/go-nessus?badge)
[![GoDoc](https://godoc.org/github.com/autonubil/go-nessus?status.svg)](https://godoc.org/github.com/autonubil/go-nessus)



## Usage

```
import "github.com/autonubil/go-nessus"
```


There are a few With... option functions that can be used to customize the API client:

- `WithBaseURL` custom base url
- `WithLogin` (username, password)
- `WithContext` (custom Context)
- `WithInsecure` allow insecure certificates
- `WithUserAgent` to set custom user agent
- `WithTrace` traces all calls

 go-nessus supports following environment variables for easy construction of a client:

- `NESSUS_URL`
- `NESSUS_USER`
- `NESSUS_PASSWORD`
- `NESSUS_INSECURE`



Construct a new nessus client, then use the various service on the client to access different parts of the wazuh API. For example, to list all agents:

```
c, err := NewClientFromEnvironment(WithTrace(true))
if err != nil {
    t.Error(err)
    return
}
i, err := c.Wellness.GetIssues()
if err != nil {
    t.Error(err)
    return
}

fmt.Printf("%v", i)
```

## ToDo

- more test cases

## Issues

- If you have an issue: report it on the [issue tracker](https://github.com/autonubil/go-nessus/issues)

## Author

Carsten Zeumer (<carsten.zeumer@autonubil.net>)

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>
