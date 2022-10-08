package post

type Post struct {
	PostID   int    `db:"post_id" json:"post_id"`
	PostRaw  string `db:"post_raw" json:"post_raw"`
	DomainID int    `db:"domain_id" json:"domain_id"`
}

type PostImage struct {
	PostImageID        int `db:"post_image_id" json:"post_image_id"`
	PostImage          int `db:"post_image" json:"post_image"`
	PostImageThumbnail int `db:"post_image_thumbnail" json:"post_image_thumbnail"`
	PostID             int `db:"post_id" json:"post_id"`
}

type PostList struct {
	Data []Post `db:"data"`
}

func (p *Post) Save() error {
	return nil
}

func (p *Post) Delete() error {
	return nil
}
