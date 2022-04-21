package storage

// LinkStorage is an interface that describes a simple storage for links.
// It can be anything, from a file to a database.
type LinkStorage interface {
	// Init used to initialize the storage (connect to the database, or something else)
	Init() error

	// Store used to store the link
	Store(string, string) error
	// Delete used to delete the link by id
	Delete(int) error
	// Get used to retrieve a link from storage using a filter
	Get(interface{}) (DatabaseEntry, error)

	// Contains returns whether the URL exists inside the storage
	Contains(string) bool
}
