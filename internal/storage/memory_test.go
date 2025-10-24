package storage

import (
	"awesomeProject/internal/models"
	"testing"
)

func TestMemoryStorage_Create(t *testing.T) {
	storage := NewMemory()
	user := models.NewUser("John Doe", "john@example.com", 30)

	// Test successful creation
	err := storage.Create(user)
	if err != nil {
		t.Fatalf("Create() failed: %v", err)
	}

	// Test duplicate user creation
	err = storage.Create(user)
	if err == nil {
		t.Error("Create() should fail for duplicate user")
	}

	// Test nil user
	err = storage.Create(nil)
	if err == nil {
		t.Error("Create() should fail for nil user")
	}
}

func TestMemoryStorage_GetByID(t *testing.T) {
	storage := NewMemory()
	user := models.NewUser("John Doe", "john@example.com", 30)

	// Test getting non-existent user
	_, err := storage.GetById("non-existent")
	if err == nil {
		t.Error("GetByID() should fail for non-existent user")
	}

	// Test get by id
	storage.Create(user)
	retrieved, err := storage.GetById(user.ID)
	if err != nil {
		t.Fatalf("GetByID() failed: %v", err)
	}

	if retrieved.ID != user.ID {
		t.Errorf("GetByID() returned wrong user ID: got %s, want %s", retrieved.ID, user.ID)
	}
	if retrieved.Name != user.Name {
		t.Errorf("GetByID() returned wrong name: got %s, want %s", retrieved.Name, user.Name)
	}
}

func TestMemoryStorage_GetByEmail(t *testing.T) {
	storage := NewMemory()
	user := models.NewUser("John Doe", "john@example.com", 30)

	// Test getting non-existent email
	_, err := storage.GetByEmail("nonexistent@example.com")
	if err == nil {
		t.Error("GetByEmail() should fail for non-existent email")
	}

	// Test getting existing email
	storage.Create(user)
	retrieved, err := storage.GetByEmail(user.Email)
	if err != nil {
		t.Fatalf("GetByEmail() failed: %v", err)
	}

	if retrieved.Email != user.Email {
		t.Errorf("GetByEmail() returned wrong email: got %s, want %s", retrieved.Email, user.Email)
	}
}

func TestMemoryStorage_GetAll(t *testing.T) {
	storage := NewMemory()

	// Test empty storage
	users, err := storage.GetAll()
	if err != nil {
		t.Fatalf("GetAll() failed: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("GetAll() should return empty slice for empty storage, got %d users", len(users))
	}

	// Test with multiple users
	user1 := models.NewUser("John Doe", "john@example.com", 30)
	user2 := models.NewUser("Jane Smith", "jane@example.com", 25)

	storage.Create(user1)
	storage.Create(user2)

	users, err = storage.GetAll()
	if err != nil {
		t.Fatalf("GetAll() failed: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("GetAll() should return 2 users, got %d", len(users))
	}
}

func TestMemoryStorage_Update(t *testing.T) {
	storage := NewMemory()
	user := models.NewUser("John Doe", "john@example.com", 30)
	storage.Create(user)

	// Test updating non-existent user
	nonExistentUser := models.NewUser("None", "none@example.com", 0)
	err := storage.Update("non-existent", nonExistentUser)
	if err == nil {
		t.Error("Update() should fail for non-existent user")
	}

	// Test successful update
	updatedUser := &models.User{
		Name:  "John Smith",
		Email: "johnsmith@example.com",
		Age:   31,
	}
	err = storage.Update(user.ID, updatedUser)
	if err != nil {
		t.Fatalf("Update() failed: %v", err)
	}

	// Verify update
	retrieved, err := storage.GetById(user.ID)
	if err != nil {
		t.Fatalf("GetByID() after update failed: %v", err)
	}
	if retrieved.Name != "John Smith" {
		t.Errorf("Update() didn't change name: got %s, want John Smith", retrieved.Name)
	}
	if retrieved.Email != "johnsmith@example.com" {
		t.Errorf("Update() didn't change email: got %s, want johnsmith@example.com", retrieved.Email)
	}
}

func TestMemoryStorage_Delete(t *testing.T) {
	storage := NewMemory()
	user := models.NewUser("John Doe", "john@example.com", 30)

	// Test deleting non-existent user
	err := storage.Delete("non-existent")
	if err == nil {
		t.Error("Delete() should fail for non-existent user")
	}

	// Test successful delete
	storage.Create(user)
	err = storage.Delete(user.ID)
	if err != nil {
		t.Fatalf("Delete() failed: %v", err)
	}

	// Verify deletion
	if storage.Exists(user.ID) {
		t.Error("User should not exist after deletion")
	}
}

func TestMemoryStorage_Exists(t *testing.T) {
	storage := NewMemory()
	user := models.NewUser("John Doe", "john@example.com", 30)

	// Test non-existent user
	if storage.Exists("non-existent") {
		t.Error("Exists() should return false for non-existent user")
	}

	// Test existing user
	storage.Create(user)
	if !storage.Exists(user.ID) {
		t.Error("Exists() should return true for existing user")
	}
}

func TestMemoryStorage_Count(t *testing.T) {
	storage := NewMemory()

	// Test empty storage
	if count := storage.Count(); count != 0 {
		t.Errorf("Count() should return 0 for empty storage, got %d", count)
	}

	// two users to test count
	user1 := models.NewUser("John Doe", "john@example.com", 30)
	user2 := models.NewUser("Jane Smith", "jane@example.com", 25)

	storage.Create(user1)
	if count := storage.Count(); count != 1 {
		t.Errorf("Count() should return 1, got %d", count)
	}

	storage.Create(user2)
	if count := storage.Count(); count != 2 {
		t.Errorf("Count() should return 2, got %d", count)
	}
}
