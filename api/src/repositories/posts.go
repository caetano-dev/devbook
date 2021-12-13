package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts struct
type Posts struct {
	db *sql.DB
}

// NewPostRepository creates a post repository
func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

//Create inserts a new post in the database
func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, error := repository.db.Prepare(`insert into posts (title, content, author_id) values (?, ?, ?)`)
	if error != nil {
		return 0, error
	}
	defer statement.Close()
	result, error := statement.Exec(post.Title, post.Content, post.AuthorID)
	if error != nil {
		return 0, error
	}
	lastInsertedID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}
	return uint64(lastInsertedID), nil
}

//FetchByID fetches a post by its id
func (repository Posts) FetchByID(postID uint64) (models.Post, error) {
	lines, error := repository.db.Query(`select p.*, u.nick from posts p inner join users u on u.id = p.author_id where p.id = ?`, postID)
	if error != nil {
		return models.Post{}, error
	}
	defer lines.Close()
	var post models.Post
	if lines.Next() {
		if error = lines.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); error != nil {
			return models.Post{}, error
		}
	}
	return post, nil
}

//Fetch fetches all posts from followed users
func (repository Posts) Fetch(userID uint64) ([]models.Post, error) {
	lines, error := repository.db.Query(`select p.*, u.nick from posts p inner 
	join users u on u.id = p.author_id inner 
	join followers s on p.author_id = s.user_id 
	where u.id = ? or s.follower_id = ? order by 1 desc`, userID, userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()
	var posts []models.Post

	for lines.Next() {
		var post models.Post
		if error = lines.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); error != nil {
			return nil, error
		}
		posts = append(posts, post)
	}
	return posts, nil
}

//Update the post
func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, error := repository.db.Prepare(`update posts set title = ?, content = ? where id = ?`)
	if error != nil {
		return error
	}
	defer statement.Close()
	if _, error = statement.Exec(post.Title, post.Content, postID); error != nil {
		return error
	}
	return nil
}

//Delete the post
func (repository Posts) Delete(postID uint64) error {
	statement, error := repository.db.Prepare(`delete from posts where id = ?`)
	if error != nil {
		return error
	}
	defer statement.Close()
	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil
}

//FetchPostByUser fetches all posts from a user
func (repository Posts) FetchPostByUser(userID uint64) ([]models.Post, error) {
	lines, error := repository.db.Query(`select p.*, u.nick from posts p join users u on u.id = p.author_id where p.author_id = ?`, userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()
	var posts []models.Post

	for lines.Next() {
		var post models.Post
		if error = lines.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); error != nil {
			return nil, error
		}
		posts = append(posts, post)
	}
	return posts, nil
}

//Like the post
func (repository Posts) Like(postID uint64) error {
	statement, error := repository.db.Prepare(`update posts set likes = likes + 1 where id = ?`)
	if error != nil {
		return error
	}
	defer statement.Close()
	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil
}

//Dislike removes the like
func (repository Posts) Dislike(postID uint64) error {
	statement, error := repository.db.Prepare(`update posts set likes = CASE WHEN likes > 0 THEN likes - 1 ELSE likes END where id = ?`)
	if error != nil {
		return error
	}
	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil
}
