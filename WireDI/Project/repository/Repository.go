package repository

type UserRepo struct{} // Structs representing the repository DB Layer

// Constructor for UserRepo which returns a pointer to UserRepo
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUser() string {
	return "User from DB"
}

// Provides GetUser() which returns "User from DB"
