package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Topic represents a topic with a name and a list of quiz data
type Topic struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	QuizData []Quiz `json:"quiz_data"`
}

type Question struct {
	Question   string   `json:"question"`
	GoodAnswer string   `json:"good_answer"`
	Answers    []string `json:"answers"`
}

type Quiz struct {
	Id        int        `json:"id"`
	Topic     string     `json:"topic"`
	Questions []Question `json:"questions"`
	Title     string     `json:"title"`
}

type NextQuiz struct {
	Id    int    `json:"id"`
	Topic string `json:"topic"`
	Title string `json:"title"`
}

type Quizzes struct {
	Id      int    `json:"id"`
	Quizzes []Quiz `json:"quizzes"`
	Name    string `json:"name"`
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

func (q *Quiz) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, q)
	}
	return err
}

func (q *Quiz) ToJSON() string {
	a, err := json.Marshal(q)
	if err != nil {
		panic(err)
	}
	return string(a)
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
