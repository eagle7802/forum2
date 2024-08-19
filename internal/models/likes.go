package models

import (
	"database/sql"
)

type UserLikeData struct {
	ID      int  `json:"postID"`
	Likes   int  `json:"likeCount"`
	IsLiked bool `json:"isLiked"`
}

type UserDislikeData struct {
	ID         int  `json:"postID"`
	Dislikes   int  `json:"dislikeCount"`
	IsDisliked bool `json:"isDisliked"`
}

type LikeModel struct {
	DB *sql.DB
}

func (LM *LikeModel) LikeInsert(likeData UserLikeData, likedByUserID int) error {
	var count int
	err := LM.DB.QueryRow("SELECT COUNT(*) FROM likes WHERE postid = ? AND likedby = ?", likeData.ID, likedByUserID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err := LM.DB.Exec(`DELETE FROM likes WHERE likedby = ? AND postid = ?`, likedByUserID, likeData.ID)
		if err != nil {
			return err
		}
		return nil
	} else {

		_, err := LM.DB.Exec(`INSERT INTO likes (postid, likedby) VALUES (?, ?)`, likeData.ID, likedByUserID)
		if err != nil {
			return err
		}

			LM.RemoveDislike(likeData.ID, likedByUserID)
		
	}

	return nil
}


func (LM *LikeModel) DislikeInsert(dislikeData UserDislikeData, dislikedByUserID int) error {
	var count int
	err := LM.DB.QueryRow("SELECT COUNT(*) FROM dislikes WHERE postid = ? AND dislikedby = ?", dislikeData.ID, dislikedByUserID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err := LM.DB.Exec(`DELETE FROM dislikes WHERE dislikedby = ? AND postid = ?`, dislikedByUserID, dislikeData.ID)
		if err != nil {
			return err
		}
		return nil
	} else {

		_, err := LM.DB.Exec(`INSERT INTO dislikes (postid, dislikedby) VALUES (?, ?)`, dislikeData.ID, dislikedByUserID)
		if err != nil {
			return err
		}

		LM.RemoveLike(dislikeData.ID, dislikedByUserID)
		
	}

	return nil
}



func (LM *LikeModel) RemoveLike(postid int, likedByUserID int) error {
	stmt := `DELETE FROM likes WHERE likedby = ? AND postid = ?`
	_, err := LM.DB.Exec(stmt, likedByUserID, postid)
	if err != nil {
		return err
	}
	var likes int
	stmt2 := `SELECT likes FROM posts WHERE id = ?`
	row := LM.DB.QueryRow(stmt2, postid)
	row.Scan(&likes)
	likes--

	stmt3 := `UPDATE posts SET likes = ? WHERE ? = id`

	_, err = LM.DB.Exec(stmt3, likes, postid)
	if err != nil {
		return err
	}
	return nil
}

func (LM *LikeModel) RemoveDislike(postid int, likedBy int) error {
	stmt := `DELETE FROM dislikes WHERE dislikedby = ? AND postid = ?`
	_, err := LM.DB.Exec(stmt, likedBy, postid)
	if err != nil {
		return err
	}
	var dislikes int
	stmt2 := `SELECT dislikes FROM posts WHERE id = ?`
	row := LM.DB.QueryRow(stmt2, postid)
	row.Scan(&dislikes)
	dislikes--

	stmt3 := `UPDATE posts SET dislikes = ? WHERE ? = id`

	_, err = LM.DB.Exec(stmt3, dislikes, postid)
	if err != nil {
		return err
	}
	return nil
}

func (LM *LikeModel) IsLikedByUser(user int, postid int) bool {
	stmt := `SELECT EXISTS (SELECT * FROM likes WHERE likedby = ? AND postid = ?)`

	var exists bool

	err := LM.DB.QueryRow(stmt, user, postid).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (LM *LikeModel) IsDislikedByUser(user int, postid int) bool {
	stmt := `SELECT EXISTS (SELECT * FROM dislikes WHERE dislikedby = ? AND postid = ?)`

	var exists bool

	err := LM.DB.QueryRow(stmt, user, postid).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (LM *LikeModel) UpdateLikesCount() error {
	stmt := `UPDATE posts
	SET likes_count = (
		SELECT COUNT(*)
		FROM likes
		WHERE likes.postid = posts.postID
	); 
	
	UPDATE posts
	SET dislikes_count = (
		SELECT COUNT(*)
		FROM dislikes
		WHERE dislikes.postid = posts.postID
	);`

	_, err := LM.DB.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}
