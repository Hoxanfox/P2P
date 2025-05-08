module migration

go 1.21

replace (
	model => ../../04-DomainLayer/model
	pool => ../pool
)

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.4
	pool v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
