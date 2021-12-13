package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users represents a user repository
type Users struct {
	db *sql.DB
}

//NewUserRespository creates a user repository
func NewUserRespository(db *sql.DB) *Users {
	return &Users{db}
}

//Create inserts user in the database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (name, nick, email, password) values(?,?,?,?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()
	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastInsertID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}
	return uint64(lastInsertID), nil
}

//Fetch fetches all users based on a filter
func (repository Users) Fetch(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	lines, error := repository.db.Query("select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?", nameOrNick, nameOrNick)

	if error != nil {
		return nil, error
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}

//FetchByID fetches a user from the database
func (repository Users) FetchByID(ID uint64) (models.User, error) {
	lines, error := repository.db.Query("select id, name, nick, email, createdAt from users where id = ?", ID)

	if error != nil {
		return models.User{}, error
	}

	defer lines.Close()

	var user models.User
	if lines.Next() {
		if error = lines.Scan(

			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return models.User{}, error
		}

	}
	return user, nil
}

// Update the information of the user
func (repository Users) Update(ID uint64, user models.User) error {
	statement, error := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		return error

	}
	return nil
}

// Delete users by ID
func (repository Users) Delete(ID uint64) error {
	statement, error := repository.db.Prepare("delete from users where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()
	if _, error = statement.Exec(ID); error != nil {
		return error
	}
	return nil
}

//FetchByEmail and returns the id and password with a hash
func (repository Users) FetchByEmail(email string) (models.User, error) {
	lines, error := repository.db.Query("select id, password from users where email = ?", email)

	if error != nil {
		return models.User{}, error
	}

	defer lines.Close()

	var user models.User
	if lines.Next() {
		if error = lines.Scan(
			&user.ID,
			&user.Password,
		); error != nil {
			return models.User{}, error
		}
	}
	return user, nil

}

//Follow allows an user to follow another
func (repository Users) Follow(userID, followerID uint64) error {
	statement, error := repository.db.Prepare("insert ignore into followers (user_id, follower_id) values(?,?)")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}
	return nil
}

//unfollow allows an user to unfollow another
func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, error := repository.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()
	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}
	return nil
}

//FetchFollowers from user
func (repository Users) FetchFollowers(userID uint64) ([]models.User, error) {
	lines, error := repository.db.Query(`select u.id, u.name, u.nick, u.email, u.createdAt from users u inner join followers s on u.id = s.follower_id where s.user_id = ?`, userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}

//FetchFollowing gets accounts the user follows
func (repository Users) FetchFollowing(userID uint64) ([]models.User, error) {
	lines, error := repository.db.Query(`select u.id, u.name, u.nick, u.email, u.createdAt from users u inner join followers s on u.id = s.user_id where s.follower_id = ?`, userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}

//FetchPassword from the user
func (repository Users) FetchPassword(userID uint64) (string, error) {
	line, error := repository.db.Query("select password from users where id = ?", userID)

	if error != nil {
		return "", error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error = line.Scan(&user.Password); error != nil {
			return "", error
		}
	}
	return user.Password, nil
}

//UpdatePassword from the user
func (repository Users) UpdatePassword(userID uint64, password string) error {
	statement, error := repository.db.Prepare("update users set password = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()
	if _, error = statement.Exec(password, userID); error != nil {
		return error
	}
	fmt.Printf(password)
	return nil
}
