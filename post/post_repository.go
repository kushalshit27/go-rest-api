package post

import (
	"context"
	"log"

	"github.com/kushalshit27/go-rest-api/internal/database"
	"github.com/kushalshit27/go-rest-api/internal/models"
)

type PostRepository interface {
	Save(post *models.Post) (*int, error)
	FindAll() ([]models.Post, error)
	Get(id *int) (*models.Post, error)
	Update(id *int, post *models.Post) (*models.Post, error)
	Remove(id *int) (*int, error)
}

type repository struct {
	db *database.DB
}

var (
	ctx = context.Background()
)

func NewPostRepository(db *database.DB) PostRepository {
	return &repository{db}
}
func (r *repository) FindAll() ([]models.Post, error) {
	queryString := `SELECT 
		b.id,
		b.title,
		b.description,
		b.created_at, 
		b.status,
		u.name,
		u.email,
		u.role,
		u.created_at
	FROM 
		blogs as b 
	JOIN 
		users as u 
	ON 
	b.created_by = u.id;`

	rows, err := r.db.Query(ctx, queryString)
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	var results []models.Post
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.ID, &p.Title, &p.Description, &p.Created, &p.Status, &p.User.Name, &p.User.Email, &p.User.Role, &p.User.CreatedAt)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		results = append(results, p)
	}
	return results, nil
}

func (r *repository) Save(post *models.Post) (*int, error) {
	sqlStatement := `INSERT INTO blogs(title, description, created_at, updated_at, status, created_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var createdId *int
	err := r.db.QueryRow(ctx, sqlStatement, post.Title, post.Description, post.Created, post.Updated, post.Status, post.CreatedBy).Scan(&createdId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return createdId, nil
}

func (r *repository) Get(id *int) (*models.Post, error) {
	var post *models.Post = new(models.Post)
	query := `SELECT id,title,description,created_at, status FROM blogs WHERE id =$1`
	err := r.db.QueryRow(ctx, query, *id).Scan(&post.ID, &post.Title, &post.Description, &post.Created, &post.Status)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return post, nil
}

func (r *repository) Update(id *int, post *models.Post) (*models.Post, error) {
	sqlStatement := `UPDATE blogs SET title=$1, description=$2, status=$3 WHERE id =$4 RETURNING id`
	err := r.db.QueryRow(ctx, sqlStatement, post.Title, post.Description, post.Status, id).Scan(&post.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return post, nil
}
func (r *repository) Remove(id *int) (*int, error) {
	deletedId := 0
	sqlStatement := `DELETE FROM blogs WHERE id =$1 RETURNING id`
	err := r.db.QueryRow(ctx, sqlStatement, id).Scan(&deletedId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &deletedId, nil
}
