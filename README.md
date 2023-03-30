# testing-demo

This is a demo repository to implement testing in a project.
It uses some trivial unit tests, then some unit tests with mocks, and finally some integration tests.

### Unit tests

Run the unit tests with
```bash
 go test ./...
```

Run the unit tests with coverage
```bash
 go test -cover ./...
```

### Run integration tests

Run the integration tests with
```bash
go test -count=1 -tags=integration ./integration_tests/...
```