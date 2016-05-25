type Model interface {
	// New returns a new Model (use an underlying pointer)
	New() Model
	// Get returns a Model corresponding to a string
	Get(id string) (Model, error)
	// Delete removes the model from the store
	Delete() (Model, error)
	// Save saves the model to a store
	Save() (Model, error)
	// GetList: Provide a map[string][]iterface{} of filters, an offset and a limit,
	// return a list of Models or error
	// filters correspond to options from GetFilterOptions
	GetList(filters map[string][]interface{}, offset uint64, limit uint64) ([]Model, error)
	// Bulk save a slice of Models
	// Provided in case your store has an optimized batch insert.
	SaveMany(models []Model) ([]Model, []error)
	// GetFilterOptions: Provide filtering to be used with GetList,
	// should be a mapping of key to type.
	// Note: uses the convertStringToType found in conversion.go
	GetFilterOptions() map[string]string
	// Validate returns a model or an error specifying why the Model in invalid.
	Validate() (Model, error)
}
