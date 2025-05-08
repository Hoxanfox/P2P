module repository

go 1.21

replace (
	dao => ../dao
	model => ../../04-DomainLayer/model
	pool => ../pool
	repository.interfaces => ../../04-DomainLayer/repository.interfaces
)

require (
	dao v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	model v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.15.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	pool v0.0.0-00010101000000-000000000000 // indirect
)
