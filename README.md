// install migrate golang

// wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb

// sudo dpkg -i migrate.linux-amd64.deb

//cli migarate: migrate create -ext sql -dir db/migration -seq init_schema

// run file in make: make migrateup

// golang testing: golang testify
// install package : go get -t github.com/TTKirito/go/db/sqlc
// github.com/lib/pq
// go init: push github
// run cli : go mod init github.com/TTKirito/go
// run cli : go mod tidy
// install testify: go get github.com/stretchr/testify