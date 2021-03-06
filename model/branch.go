package model

import (
	"time"
)

type Paper struct {
	PaperID          string `gorm:"type:varchar(30);primary_key;" json:"paper_id"`
	Title            string `gorm:"type:varchar(400);not null;" json:"title"`
	PaperPublishYear string `gorm:"type:varchar(5)" json:"paper_publish_year"`
	ConferenceID     string `gorm:"type:varchar(10)" json:"conference_id"`
}

type Affiliation struct {
	AffiliationID   string `gorm:"type:varchar(30);primary_key;" json:"affiliation_id"`
	AffiliationName string `gorm:"type:varchar(150)" json:"affiliation_name"`
}

type Author struct {
	AuthorID   string `gorm:"type:varchar(30);primary_key;" json:"author_id"`
	AuthorName string `gorm:"type:varchar(100)" json:"author_name"`
}

type Conference struct {
	ConferenceID   string `gorm:"type:varchar(30);primary_key;" json:"conference_id"`
	ConferenceName string `gorm:"type:varchar(30)" json:"conference_name"`
}

type PaperAuthorAffiliation struct {
	PaperID        string `gorm:"type:varchar(30);primary_key;" json:"paper_id"`
	AuthorID       string `gorm:"type:varchar(30)" json:"author_id"`
	AffiliationID  string `gorm:"type:varchar(30)" json:"affiliation_id"`
	AuthorSequence string `gorm:"type:varchar(3);primary_key;" json:"author_sequence"`
}

type PaperReference struct {
	PaperID             string `gorm:"type:varchar(30);" json:"paper_id"`
	PaperTitle          string `gorm:"type:varchar(400);" json:"paper_title"`
	ReferenceID         string `gorm:"type:varchar(30);" json:"reference_id"`
	ReferencePaperTitle string `gorm:"type:varchar(400);" json:"reference_paper_title"`
}

type Connection struct {
	Author1ID     string `gorm:"type:varchar(30)" json:"author1_id"`
	Author1Name   string `gorm:"type:varchar(100)" json:"author1_name"`
	Author1HIndex int64  `json:"author1_h_index"`
	Author2ID     string `gorm:"type:varchar(30)" json:"author2_id"`
	Author2Name   string `gorm:"type:varchar(100)" json:"author2_name"`
	Author2HIndex int64  `json:"author2_h_index"`
	CoNum         int64  `json:"co_num"`
}

type Comment struct {
	//gorm.Model
	CommentID   uint64    `gorm:"primary_key;" json:"comment_id"`
	Username    string    `gorm:"size:15" json:"username"`
	AuthorName  string    `gorm:"size:30" json:"author_name"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	PaperID     string    `gorm:"size:30" json:"paper_id"`
	CommentTime time.Time `json:"comment_time"`
	Content     string    `gorm:"size:255" json:"content"`
	OnTop       bool      `gorm:"default:false" json:"on_top"`
	Like        uint64    `json:"like"`
	Dislike     uint64    `json:"dislike"`
}

type CommentLike struct {
	CommentID     uint64 `gorm:"primary_key;" json:"comment_id"`
	UserID        uint64 `gorm:"primary_key;" json:"user_id"`
	LikeOrDislike bool   `json:"like_or_dislike"`
}

type LinkType struct {
	RootID string `json:"rootID"`
	Nodes  []Node `json:"nodes"`
	Links  []Link `json:"links"`
}
type Node struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
type Link struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

// 用于AuthorConnectionGraph的Response struct
type A struct {
	Bs []B `json:"data"`
	Cs []C `json:"links"`
}

type B struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Value int64  `json:"value"`
}

type C struct {
	Source     string `json:"source"`
	SourceName string `json:"source_name"`
	Target     string `json:"target"`
	TargetName string `json:"target_name"`
	Num        int64  `json:"value"`
}

// 用于PaperReferenceGraph的Response struct
type M struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Ms   []M    `json:"children"`
}
