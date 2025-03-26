package repository

import (
	"employee-auth/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(users *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByRole(role string) (*models.User, error)
	FindByDob(dob string) (*models.User, error)
	UpdateUser(user *models.User) error
	FindByName(name string) (*models.User, error)
	FindByNIP(nip string) (*models.User, error)
	FindByIdentifier(identifier string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	if db == nil {
		panic("Database connection is nil")
	}
	return &userRepository{db: db}
}

// FindByIdentifier searches for a user by email, username, or NIP
func (repo *userRepository) FindByIdentifier(identifier string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ? OR username = ? OR nip = ?", identifier, identifier, identifier).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Create Users
func (r *userRepository) CreateUser(users *models.User) error {
	return r.db.Create(users).Error
}

// Find User By Username
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find User by Email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find User by Role
func (r *userRepository) FindByRole(role string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("role = ?", role).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find User by Date of Birth
func (r *userRepository) FindByDob(dob string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("dob = ?", dob).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update User
func (r *userRepository) UpdateUser(user *models.User) error {

	return r.db.Save(user).Error

}

// Find User by Name
func (r *userRepository) FindByName(name string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("full_name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

// Find User by NIP
func (r *userRepository) FindByNIP(nip string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("nip = ?", nip).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find User By UUID
func (r *userRepository) FindByNip(nip string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("nip = ?", nip).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
