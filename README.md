# Overview

Earlier this year I bought an [Awair Element](https://www.getawair.com/products/element).  I wanted to understand my indoor air quality better.  I looked for a device that could measure CO2 + particulate matter + VOC, and that most importantly had a local API.  I don't want to rely on a vendor's cloud service or app that might be discontinued in 6 months.  The Awair Element fit these requirements perfectly.

I wrote this small go service to collect and store the measurements from the Awair Element device.  It polls the Awair API on the sensor itself and stores the retrieved measurements in InfluxDB, where they can be easialy monitored using other tools like Grafana.

The local API has to be enabled on the Awair Element.  Awair has instructions on how to do that [here](https://support.getawair.com/hc/en-us/articles/360049221014-Awair-Element-Local-API-Feature).

# Building
`go build .`

# Running
`AM_CONFIG=path/to/config.yaml ./airmon`

By default, configuration will be loaded from a YAML file specified by the AM_CONFIG environment variable.

Also by default, this will run continuously.  You can use `-once` to only collect metrics once, then exit.

# Configuration Options
- Interval - how often to poll API (in seconds)
- Sensors - list of sensors to collect data from.  Only 'awair' type sensors are supported.
- Datastores - list of datastores to send collected data to.  Only 'influx' type datastores are supported.
- Monitor endpoint (optional) - endpoint to call each time job runs successfully

An example YAML config file is available [here](./sample-config.yml).

# Monitoring
I use [Uptime Kuma](https://github.com/louislam/uptime-kuma) to monitor this service and alert if it hasn't run.  If a monitor endpoint is specified in the configuration, airmon will perform an HTTP GET of the specified URL on each run.

This is optional - if the monitor endpoint is not included in the configuration, it will just skip this.
