// install migrate golang

// wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb

// sudo dpkg -i migrate.linux-amd64.deb

//cli migarate: migrate create -ext sql -dir db/migration -seq init_schema

// run file in make: make migrateup