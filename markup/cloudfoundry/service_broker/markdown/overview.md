# Service Broker API

## Overview
The Cloud Foundry services API defines the contract between the Cloud Controller and the service broker. The broker is expected to implement several HTTP (or HTTPS) endpoints underneath a URI prefix. One or more services can be provided by a single broker, and load balancing enables horizontal scalability of redundant brokers. Multiple Cloud Foundry instances can be supported by a single broker using different URL prefixes and credentials. [Learn more about the Service Broker API.](http://docs.cloudfoundry.org/services/api.html)


### Version information
Version: 2.5

### URI scheme
Host: 127.0.0.1:8888
BasePath: /v2
Schemes: HTTP

