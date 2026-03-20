# To Do List for `meinberg_ltos_exporter`

- [x] Add flag to skip SSL cert verification of Meinberg device
- [x] Extend mock-server to support HTTPS endpoint
- [x] Add support for basic auth in mock-server.go
- [x] Test code with basic auth
- [x] Refactor code to make collector slimmer, possibly move API response parsing and validation code to client or even to a dedicated model
- [ ] Split models.go into smaller logical chunks (system, network, notification, etc.)
- [ ] Split collector.go into smaller chunks (following API response structure) that can be enabled/disabled
- [x] Verify units of all metrics (milliseconds vs seconds, kB vs bytes, etc.)
- [ ] Testing with live M600 system
- [x] Release "first" version
- [x] Create more build artifacts
- [x] Create Docker images via Dockerfile
- [x] Add (more) (debug) logging
- [ ] Add network and interface metrics
