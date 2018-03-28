commit=`git rev-parse --short HEAD`
go build -ldflags "-X main.version_commit=${commit}" gitcredential.go
