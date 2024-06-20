package models

type TagName string

const (
	PlayQuizTag   TagName = "PlayQuizTag"
	PlayGameTag   TagName = "PlayGameTag"
	WonQuizTag    TagName = "WonQuizTag"
	WonGameTag    TagName = "WonGameTag"
	ConnectionTag TagName = "ConnectionTag"
	LevelTag      TagName = "LevelTag"
	AvatarTag     TagName = "AvatarTag"
)

type Tag struct {
	Id   int     `json:"id" gorm:"primary_key"`
	Name TagName `json:"name"`
}
