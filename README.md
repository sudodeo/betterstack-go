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
- [ ] Monitors
- [ ] Monitor groups
- [ ] Heartbeats
- [ ] Heartbeat groups
- [ ] On-call calendar
- [ ] Escalation policies
- [ ] Incidents
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
- [ ] Status Pages
- [ ] Status Pages Sections
- [ ] Status Pages Resources 
- [ ] Status Pages Reports 
- [ ] Status Pages Report Updates 
- [ ] Metadata

### Logs
- [ ] Sources
- [ ] Query
<!-- ## Usage/Examples

```golang
package main

import (
    "github.com/sudodeo/betterstack-go"
)

func main(){
    dp, err := doppler.NewFromEnv()
    if err != nil {
        panic(err)
    }

    projects, err := dp.ListMonitors()
    if err != nil {
        panic(err)
    }

    log.Println(projects)
}
``` -->

## Author

- [@sudodeo](https://www.github.com/sudodeo)
