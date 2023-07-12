type UserRepository struct{
	db *Database
}
// hubungan sama DB

func NewUserRepository(db *Database) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

// contoh user ID 
func (ur *UserRepository) GetUserById(id string, user *models.User) error{
	query := "bla bla bla"
	err := r.db.query("blabla bla")
	if err != nill{
		return err
	}

	return nill
}
''
// contoh user ID 
func (ur *UserRepository) GetUserById(id string, user *models.User) error{
	query := "bla bla bla"
	err := r.db.query("blabla bla")
	if err != nill{
		return err
	}

	return nill
}