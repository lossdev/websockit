docsurl := "http://localhost:8080/pkg/github.com/lossdev/websockit"

_default:
  @just --list --list-prefix '  > '

# host manium docs with godoc and open documentation page
docs:
  sleep 2 && if [ {{os()}} == "macos" ]; then open {{docsurl}}; else xdg-open {{docsurl}}; fi &
  godoc -http=localhost:8080

# test all packages
test:
  go test -v ./...

# test all packages, generate coverfile, and open coverfile in browser
test-cover:
  go test -v ./... -coverprofile=c.out; go tool cover -html=c.out

# remove manium binary and any residual test files
clean:
  rm -f c.out
