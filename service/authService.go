package service

import (
	"employee-auth/config"
	"employee-auth/models"
	"employee-auth/repository"
	"employee-auth/utils"
	"fmt"
	"time"
)

var userRepo repository.UserRepository

// RegisterUser registers a new user
func RegisterUser(nip, username, fullName, email, password, role string, dob time.Time) (*models.User, error) {
	if userRepo == nil {
		userRepo = repository.NewUserRepository(config.DB)
	}

	//
	// Validate user data
	userData := models.User{Username: username, Email: email, PasswordHash: password, Role: models.Role(role)}
	if validationErr := utils.ValidateStruct(userData, "Username", "Email", "PasswordHash"); validationErr != "" {
		return nil, fmt.Errorf(validationErr)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create user object
	user := &models.User{
		NIP:          nip,
		Username:     username,
		FullName:     fullName,
		Email:        email,
		PasswordHash: hashedPassword,
		Dob:          dob,
		Role:         models.Role(role),
	}

	// Save user to the database
	if err := userRepo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return user, nil
}

// LoginUser authenticates a user
func LoginUser(identifier, password string) (map[string]interface{}, error) {
	// Find user by email, username, or NIP
	user, err := userRepo.FindByIdentifier(identifier)
	if err != nil {
		return nil, fmt.Errorf("User not found")
	}

	// Validate password
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("Invalid Credentials")
	}

	// Generate JWT Token
	token, err := utils.GenerateJWT(user.NIP, user.FullName, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("error generating JWT: %v", err)
	}

	return map[string]interface{}{
		"token": token,
		"role":  user.Role,
	}, nil
}
