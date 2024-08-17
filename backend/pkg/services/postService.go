package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"time"
	
	"github.com/google/uuid"
)

type PostService struct {
	db *sql.DB
}

func NewPostservice() *PostService {
	dbs := sqlite.GlobalDB
	
	return &PostService{
		db: dbs.GetDB(),
	}
}
func (p *PostService) GetDB() *sql.DB {
	return p.db
}

func (p *PostService) SetDB(db *sql.DB) {
	p.db = db
}

func (p *PostService) CreatePost(post models.Post) error {
	_, err := p.db.Exec(`INSERT INTO posts (userId, title, content, imageP, creatdate, statut) VALUES (?, ?, ?, ?, ?, ?)`, post.UserId, post.Title, post.Content, post.Image, post.CreatedAt, post.Statut)
	return err
}

func (p *PostService) GetAllPosts() ([]models.Post, error) {
	//rows, err := p.db.Query(`SELECT id, title, content, image, user_id, created_at FROM posts ORDER BY created_at DESC`)
	rows, err := p.db.Query("SELECT id, userId, title, content, imageP, creatdate, statut  FROM posts ORDER BY creatdate DESC`")
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		var (
			id        int
			userId    uuid.UUID
			title     string
			content   string
			image     []byte
			creatdate time.Time
			statut    string
		)
		err = rows.Scan(&id,&userId, &title, &content, &image, &creatdate, statut)

		if err != nil {
			fmt.Println("Error", err)
			return nil, err
		}
		post.Id = id
		post.UserId = userId
		post.Title = title
		post.Content = content
		post.Image = image
		post.CreatedAt = creatdate
		post.Statut = statut
		posts = append(posts, post)
	}
	return posts, nil
}
