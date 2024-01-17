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

// Quiz represents a quiz with a title and a list of questions.
type Quiz struct {
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

// QuizData represents the overall structure of the provided quiz
type QuizData struct {
	Id    int    `json:"id"`
	Quiz  Quiz   `json:"quiz"`
	Topic string `json:"topic"`
}

// Topic represents a topic with a name and a list of quiz data
type Topic struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	QuizData []QuizData `json:"quiz_data"`
}

type QuizJSONMap map[string]interface{}

type QuizGame struct {
	Id        int         `json:"id"`
	Gid       int         `json:"game_id"`
	Data      QuizJSONMap `json:"data"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

func (q *QuizJSONMap) Value() (driver.Value, error) {
	return json.Marshal(q)
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (q *QuizJSONMap) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &q)
	case string:
		return json.Unmarshal([]byte(v), &q)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
