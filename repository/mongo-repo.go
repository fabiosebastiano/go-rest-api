package repository

// NewMongoRepository create a new repo for MONGO DB
func NewMongoRepository() PostRepository {
	return &repo{}
}
