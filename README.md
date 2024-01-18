# BetterStack-Go

Golang SDK client for betterstack.com APIs.

## Installation

```bash
go get -u github.com/sudodeo/betterstack-go
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
- [ ] New Relic Integrations
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

<!-- ## Usage/Examples

```golang
package main

import (
    betterstack "github.com/sudodeo/betterstack-go"
)

func main(){
    bs, err := betterstack.NewFromENV()
    if err != nil {
        panic(err)
    }

    monitors, err := bs.ListMonitors()
    if err != nil {
        panic(err)
    }

    log.Println(monitors)
}
``` -->

## Author

- [@sudodeo](https://www.github.com/sudodeo)
