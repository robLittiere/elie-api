package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Question struct {
	Question     string   `json:"question"`
	GoodAnswer   string   `json:"good_answer"`
	WrongAnswers []string `json:"wrong_answers"`
}

// Topic represents a topic with a name and a list of quiz data
type Topic struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	QuizData []Quiz `json:"quiz_data"`
}

type QuizGameJSONMap map[string]interface{}

type QuizGame struct {
	Id        int             `json:"id"`
	Gid       int             `json:"game_id"`
	Game      Game            `json:"game" gorm:"foreignkey:Gid"`
	Data      QuizGameJSONMap `json:"data" gorm:"type:jsonb"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

func (q *QuizGameJSONMap) Value() (driver.Value, error) {
	return json.Marshal(q)
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (q *QuizGameJSONMap) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &q)
	case string:
		return json.Unmarshal([]byte(v), &q)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
