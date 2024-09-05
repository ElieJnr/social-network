package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
)

type PostService struct {
	db *sql.DB
}

func NewPostService() *PostService {
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

// fonction d'insertion des posts
func (p *PostService) InsertPost(postValue models.CheckResult, w http.ResponseWriter, r *http.Request, origin string) error {
	user, err := utils.CurrentUser(w, r)
	if err != nil {
		fmt.Println("error getting current user")
		return err
	}
	postID, err := utils.GenerateUuid()
	if err != nil {
		fmt.Println("error generating post ID")
		return err
	}

	tx, err := p.db.Begin()
	if err != nil {
		fmt.Println("error in transaction")
		return err
	}

	defer tx.Rollback()

	var status string
	if origin == "group" {
		status = origin
	} else {
		status = postValue.Status
	}

	_, err = tx.Exec("INSERT INTO Posts (id, userId, content, imageUrl, statut) VALUES (?, ?, ?, ?, ?)",
		postID, user.UserId, postValue.Content, postValue.PhotoURL, status)
	if err != nil {
		fmt.Println("error inserting post")
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO UserPost (postId, userId, statut) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("error preparing statement")
		return err
	}
	defer stmt.Close()

	if origin == "group" {
		// Pour les groupes, on insère simplement lid du group
		_, err = stmt.Exec(postID, postValue.GroupId, status)
		if err != nil {
			fmt.Println("error executing statement for group")
			return err
		}
	} else {
		// Logique existante pour les posts
		if status == "public" || status == "private" {
			_, err = stmt.Exec(postID, user.UserId, status)
			if err != nil {
				fmt.Println("error executing statement")
				return err
			}
		} else if status == "almost-private" {
			_, err = stmt.Exec(postID, user.UserId, status)
			if err != nil {
				fmt.Println("error executing statement")
				return err
			}

			for _, allowedUserID := range postValue.AllowedUsers {
				_, err = stmt.Exec(postID, allowedUserID, status)
				if err != nil {
					fmt.Println("error executing statement")
					return err
				}
			}
		}
	}

	fmt.Println("post/group inserted successfully")
	return tx.Commit()
}

// se charge de récuperer les posts et de le partager au service de post
func (p *PostService) GetAllPosts(w http.ResponseWriter, r *http.Request) ([]models.Posts, error) {
	query := `SELECT id, userId, content, imageUrl, statut, createDate FROM Posts ORDER BY createDate DESC`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %w", err)
	}
	defer rows.Close()

	var allPosts []models.Posts
	for rows.Next() {
		var post models.Posts
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Content, &post.Image_url, &post.Post_status, &post.Creation_date); err != nil {
			fmt.Println("err: rowsScan", err)
			continue
		}

		if err := postDetails(p.db, &post, w, r, ""); err != nil {
			fmt.Println("err: postDetails", err)
			continue
		}

		allPosts = append(allPosts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over posts: %w", err)
	}

	fmt.Println("all posts retrieved successfully")

	return allPosts, nil
}

// se charge de récuperer les posts d'un groupe et de le partager au service de post
func (p *PostService) GetGroupPosts(w http.ResponseWriter, r *http.Request, groupId string) ([]models.Posts, error) {
	query := `
    SELECT p.id, p.userId, p.content, p.imageUrl, p.statut, p.createDate
    FROM Posts p
    INNER JOIN UserPost up ON p.id = up.postId
    WHERE up.userId = ? AND up.statut = 'group'
    ORDER BY p.createDate DESC`
	rows, err := p.db.Query(query, groupId)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %w", err)
	}
	defer rows.Close()

	var allPosts []models.Posts
	for rows.Next() {
		var post models.Posts
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Content, &post.Image_url, &post.Post_status, &post.Creation_date); err != nil {
			fmt.Println("err: rowsScan", err)
			continue
		}

		if err := postDetails(p.db, &post, w, r, groupId); err != nil {
			fmt.Println("err: postDetails", err)
			continue
		}

		allPosts = append(allPosts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over posts: %w", err)
	}

	fmt.Println("all posts/group retrieved successfully")

	return allPosts, nil
}

func (p *PostService) GetPostsById(w http.ResponseWriter, r *http.Request, userId string) ([]models.Posts, error) {
	ownPosts, err := utils.GetPostsById(p.db, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get own posts: %w", err)
	}

	for i := range ownPosts {
		if err := postDetails(p.db, &ownPosts[i], w, r, ""); err != nil {
			return nil, fmt.Errorf("failed to get post details: %w", err)
		}
	}

	return ownPosts, nil
}

// se charge de récuperer les informations d'un post ie author, nbr de like, qui peut voir le post, nbr de commentaires...
func postDetails(db *sql.DB, post *models.Posts, w http.ResponseWriter, r *http.Request, groupId string) error {
	currentUser, err := utils.CurrentUser(w, r)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	author, err := utils.GetAuthor(db, post.UserID)
	if err != nil {
		return fmt.Errorf("failed to get post author: %w", err)
	}

	nbrComment, err := utils.GetNbrComment(db, post.PostID)
	if err != nil {
		return fmt.Errorf("failed to get number of comments: %w", err)
	}

	nbrLike, err := utils.GetNbrLike(db, post.PostID)
	if err != nil {
		return fmt.Errorf("failed to get number of likes: %w", err)
	}

	likeStatus, dislikeStatus, err := utils.GetLikeDislikeStatus(db, post.PostID, currentUser.UserId)
	if err != nil {
		return fmt.Errorf("failed to get like/dislike status: %w", err)
	}

	isVisible, err := utils.CheckVisibility(db, post.UserID, post.PostID, currentUser.UserId, post.Post_status, groupId)
	if err != nil {
		return fmt.Errorf("failed to get like/dislike status: %w", err)
	}

	isFollowing, err := utils.IsFollowing(db, post.UserID, currentUser.UserId)
	if err != nil {
		return fmt.Errorf("failed to check if user is following: %w", err)
	}

	post.Author = author
	post.Formated_date = utils.FormatTimeAgo(post.Creation_date)
	post.Can_see = isVisible
	post.Like_nbr = nbrLike
	post.Comments_nbr = nbrComment
	post.Like_status = likeStatus
	post.IsFollower = isFollowing
	post.HasImage = post.Image_url != ""
	post.Dislike_status = dislikeStatus

	return nil
}
