package entity

import (
	"io"
	"time"

	"gopkg.in/guregu/null.v4"
)

type Post struct {
	id          int
	authorID    int
	author      *User
	title       string
	content     io.Reader
	publishedAt null.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func NewPost() *Post {
	return &Post{}
}

func (p *Post) WithID(id int) *Post {
	if p == nil {
		return nil
	}
	p.id = id
	return p
}

func (p *Post) WithAuthorID(authorID int) *Post {
	if p == nil {
		return nil
	}
	p.authorID = authorID
	return p
}

func (p *Post) WithAuthor(author *User) *Post {
	if p == nil {
		return nil
	}
	p.author = author
	return p
}

func (p *Post) WithTitle(title string) *Post {
	if p == nil {
		return nil
	}
	p.title = title
	return p
}

// io.Readerを使い回すとバグらないか心配
func (p *Post) WithContent(content io.Reader) *Post {
	if p == nil {
		return nil
	}
	p.content = content
	return p
}

func (p *Post) WithPublishedAt(publishedAt null.Time) *Post {
	if p == nil {
		return nil
	}
	p.publishedAt = publishedAt
	return p
}

func (p *Post) WithCreatedAt(createdAt time.Time) *Post {
	if p == nil {
		return nil
	}
	p.createdAt = createdAt
	return p
}

func (p *Post) WithUpdatedAt(updatedAt time.Time) *Post {
	if p == nil {
		return nil
	}
	p.updatedAt = updatedAt
	return p
}

func (p *Post) GetID() int {
	if p == nil {
		return 0
	}
	return p.id
}

func (p *Post) GetAuthorID() int {
	if p == nil {
		return 0
	}
	return p.authorID
}

func (p *Post) GetAuthor() *User {
	if p == nil {
		return nil
	}
	return p.author
}

func (p *Post) GetTitle() string {
	if p == nil {
		return ""
	}
	return p.title
}

func (p *Post) GetContent() io.Reader {
	if p == nil {
		return nil
	}
	return p.content
}

func (p *Post) GetPublishedAt() null.Time {
	if p == nil {
		return null.Time{}
	}
	return p.publishedAt
}

func (p *Post) GetCreatedAt() time.Time {
	if p == nil {
		return time.Time{}
	}
	return p.createdAt
}

func (p *Post) GetUpdatedAt() time.Time {
	if p == nil {
		return time.Time{}
	}
	return p.updatedAt
}
