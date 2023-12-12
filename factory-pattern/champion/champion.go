package champion

// Champion is an interface for the factory pattern
type Champion interface {
	Move()
	Attack()
	GetName() string
}
