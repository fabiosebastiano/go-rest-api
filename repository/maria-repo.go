package repository

// NewMariaRepository create a new repo for MARIA DB
func NewMariaRepository() PostRepository {
	return &repo{}
}
