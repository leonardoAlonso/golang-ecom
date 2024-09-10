package user

import (
	"database/sql"
	"fmt"

	"github.com/leonardoAlonso/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db} // Use the repository pattern to use the database connection
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		// If the user is not found, return nil
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}

func (s *Store) CreateUser(u types.User) error {
	return nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func scanRowIntoUser(row *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
