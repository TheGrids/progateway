package models

import (
	"github.com/google/uuid"
)

type Test struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title     string
	Questions []Question // Связь с таблицей вопросов
}

type Question struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Text    string
	TestID  uuid.UUID // Внешний ключ для связи с тестом
	Answers []Answer  // Связь с таблицей ответов
}

type Answer struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Text       string
	IsCorrect  bool
	QuestionID uuid.UUID // Внешний ключ для связи с вопросом
}

type UserTestResult struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID  uuid.UUID // Идентификатор пользователя
	TestID  uuid.UUID // Идентификатор теста
	Score   int
	Answers []UserAnswer `gorm:"foreignkey:UserTestResultID"` // Связь с таблицей ответов пользователя
}

type UserAnswer struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserTestResultID uuid.UUID // Идентификатор результата теста
	QuestionID       uuid.UUID // Идентификатор вопроса
	AnswerID         uuid.UUID // Идентификатор выбранного пользователем ответа
}
