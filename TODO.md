# To Do List for `meinberg_ltos_exporter`

- [x] Add flag to skip SSL cert verification of Meinberg device
- [x] Extend mock-server to support HTTPS endpoint
- [x] Add support for basic auth in mock-server.go
- [x] Test code with basic auth
- [ ] Refactor code to make collector slimmer, possibly move API response parsing and validation code to client or even to a dedicated model
- [ ] Verify units of all metrics (milliseconds vs seconds, kB vs bytes, etc.)
- [ ] Testing with live M600 system
- [x] Release "first" version
- [x] Create more build artifacts
- [ ] Create Docker images via Dockerfile
- [ ] Add network and interface metrics
- [ ] (?) Extend mock-server.go to provide separate HTTP port with delayed response (for timeout testing)
- [ ] (?) Remove password command line option (b/c insecure and leaks to shell history)
