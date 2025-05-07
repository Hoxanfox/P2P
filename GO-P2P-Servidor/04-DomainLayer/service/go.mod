module service

go 1.21

replace (
	factory => ../factory
	model => ../model
	observer => ../observer
	repository.interfaces => ../repository.interfaces
)

require (
	github.com/google/uuid v1.6.0
	model v0.0.0-00010101000000-000000000000
)
