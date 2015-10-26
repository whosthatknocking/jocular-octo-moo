Snuffles - Collect Docker stats and report to monitoring application
===

This command line tool gathers information about docker containers, images, and devicemapper storage driver and reports them to monitoring application. To run this check make sure you have appropriate privileges to access docker socket.


### Exported docker info

  - `Containers` reports total number of containers
  - `Images` reports total number of downloaded images


### Exported storage driver info

Items in the `Storage Driver: devicemapper` section are status information
about the used driver.
  - `Data Space Used` reports how much of `Data file` is currently used
  - `Data Space Total` reports max size the `Data file`
  - `Data Space Available` reports how much free space there is in the `Data file`. If you are using a loop device this will report the actual space available to the loop device on the underlying filesystem.
  - `Metadata Space Used` reports how much of `Metadata` is currently used
  - `Metadata Space Total` reports max size the `Metadata`
  - `Metadata Space Available` reports how much free space there is in the `Metadata`. If you are using a loop device this will report the actual space available to the loop device on the underlying filesystem.


### Spec

  - API (v1.20)
  - [Docker Remote API](https://docs.docker.com/reference/api/docker_remote_api/)
  - [Docker Remove API v1.20](https://docs.docker.com/reference/api/docker_remote_api_v1.20/)
  - [Docker Device Mapper driver](https://github.com/docker/docker/tree/master/daemon/graphdriver/devmapper)
