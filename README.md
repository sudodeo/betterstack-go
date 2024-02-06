# BetterStack-Go

Explore seamless integration with BetterStack's powerful APIs using this Golang SDK. This personal project is crafted to streamline your development process, offering a robust set of tools for interacting with [Betterstack](https://betterstack.com) services. While it's not officially affiliated with BetterStack, this SDK aims to provide a convenient and efficient way for developers to enhance their projects with the flexibility of the BetterStack Golang SDK.

## Installation

```bash
go get -u github.com/sudodeo/betterstack-go/sdk
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`UPTIME_API_TOKEN`

To obtain your uptime API token, view instructions from [BetterStack](https://betterstack.com/docs/uptime/api/getting-started-with-uptime-api/)

`LOGS_API_TOKEN`

To obtain your logs API token, view instructions from [BetterStack](https://betterstack.com/docs/logs/api/getting-started/)

## Documentation

[Uptime API Documentation](https://betterstack.com/docs/uptime/api/getting-started-with-uptime-api/)

[Logs API Documentation](https://betterstack.com/docs/logs/api/getting-started/)

## Features

### Uptime

- [x] Monitors
- [x] Monitor groups
- [x] Heartbeats
- [x] Heartbeat groups
- [x] On-call calendar
- [ ] Escalation policies
- [x] Incidents
- [x] New Relic Integrations
- [ ] Datadog Integrations
- [ ] AWS Cloudwatch Integrations
- [ ] Azure Integrations
- [ ] Google Monitoring Integrations
- [ ] Grafana Integrations
- [ ] Prometheus Integrations
- [ ] Email Integrations
- [ ] Incoming webhooks
- [ ] Splunk On-call Integrations
- [ ] PagerDuty Integrations
- [x] Status Pages
- [x] Status Pages Sections
- [x] Status Pages Resources
- [x] Status Pages Reports
- [x] Status Pages Report Updates
- [x] Metadata

### Logs

- [x] Sources
- [x] Query

## Usage/Examples

```golang
package main

import (
    "github.com/sudodeo/betterstack-go/sdk"
    betterstack "github.com/sudodeo/betterstack-go/sdk/uptime" // for uptime API
)

func main(){
    bs, err := betterstack..New(&sdk.Config{APIToken: token, APIType: "uptime"})
    if err != nil {
        panic(err)
    }

    incidents, err := bs.ListIncidents()
    if err != nil {
        panic(err)
    }

    log.Println(incidents)
}
```

## Author

- [@sudodeo](https://www.github.com/sudodeo)
