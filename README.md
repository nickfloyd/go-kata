# go personal kata

Uses the [Go (Golang) Codespaces starter template](https://github.com/codespaces-examples/go)

## Play the Kata

1. Fork the repo / or clone and branch
2. run `go test ./...`
3. Note all the tests are failing
4. Pick a section under `./src` that you want to practice (i.e. `basics.go`)
5. Write logic to get the tests to pass
6. Run the tests for the section you chose until all :white_check_mark:

### Flow

In this example we'll use `basics.go` as the chosen practice section

1. Run ` go test src/basics_test.go src/basics.go`
2. Note that they all fail
3. Open `basics.go` and add code to make the tests run :white_check_mark:

Note: You can also do the following if you want to run individual tests

1. Run `cd src`
2. Run ` go test -v -run TestIfStatements` (ex)
