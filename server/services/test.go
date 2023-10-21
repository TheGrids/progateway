package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"progateway.ru/m/v2/models"
)

func AddTest(con *gin.Context) {
	test1 := models.Test{
		Title: "Java",
		Questions: []models.Question{
			{
				Text: "Что такое Java Virtual Machine (JVM)?",
				Answers: []models.Answer{
					{Text: "Язык программирования", IsCorrect: false},
					{Text: "Система сборки мусора", IsCorrect: false},
					{Text: "Виртуальная машина для выполнения Java-программ", IsCorrect: true},
					{Text: "Среда разработки", IsCorrect: false},
				},
			},
			{
				Text: "Какой модификатор доступа используется для классов, методов и полей, к которым можно обратиться из любого места?",
				Answers: []models.Answer{
					{Text: "public", IsCorrect: true},
					{Text: "private", IsCorrect: false},
					{Text: "protected", IsCorrect: false},
					{Text: "default", IsCorrect: false},
				},
			},
			{
				Text: "Как создать экземпляр класса в Java?",
				Answers: []models.Answer{
					{Text: "new Object();", IsCorrect: true},
					{Text: "create Object();", IsCorrect: false},
					{Text: "initialize Object();", IsCorrect: false},
					{Text: "object.create();", IsCorrect: false},
				},
			},
			{
				Text: "Как объявить переменную в Java, которая не может быть изменена после присвоения значения?",
				Answers: []models.Answer{
					{Text: "final", IsCorrect: true},
					{Text: "const", IsCorrect: false},
					{Text: "static", IsCorrect: false},
					{Text: "immutable", IsCorrect: false},
				},
			},
			{
				Text: "Каким образом вы можете выполнить переход к следующей итерации в цикле for в Java?",
				Answers: []models.Answer{
					{Text: "continue;", IsCorrect: true},
					{Text: "skip;", IsCorrect: false},
					{Text: "next;", IsCorrect: false},
					{Text: "break;", IsCorrect: false},
				},
			},
			{
				Text: "Что такое JavaBeans в контексте Java?",
				Answers: []models.Answer{
					{Text: "Место для хранения кофе", IsCorrect: false},
					{Text: "Библиотека для обработки звука", IsCorrect: false},
					{Text: "Классы Java, предоставляющие сеттеры и геттеры для доступа к свойствам", IsCorrect: true},
					{Text: "Классы для манипуляции строками", IsCorrect: false},
				},
			},
			{
				Text: "Какой интерфейс в Java используется для обработки событий кнопок?",
				Answers: []models.Answer{
					{Text: "ActionListener", IsCorrect: true},
					{Text: "ButtonListener", IsCorrect: false},
					{Text: "ClickListener", IsCorrect: false},
					{Text: "EventHandler", IsCorrect: false},
				},
			},
			{
				Text: "Каким образом вы можете создать множество наследников от одного класса в Java?",
				Answers: []models.Answer{
					{Text: "Используя наследование", IsCorrect: true},
					{Text: "Используя интерфейсы", IsCorrect: false},
					{Text: "Используя агрегацию", IsCorrect: false},
					{Text: "Это невозможно в Java", IsCorrect: false},
				},
			},
			{
				Text: "Какие исключения в Java являются проверяемыми (checked) и должны быть обработаны или объявлены?",
				Answers: []models.Answer{
					{Text: "RuntimeException", IsCorrect: false},
					{Text: "IOException", IsCorrect: true},
					{Text: "NullPointerException", IsCorrect: false},
					{Text: "ArrayIndexOutOfBoundsException", IsCorrect: true},
				},
			},
			{
				Text: "Какой метод используется для запуска приложения Java (точка входа)?",
				Answers: []models.Answer{
					{Text: "run()", IsCorrect: false},
					{Text: "start()", IsCorrect: true},
					{Text: "main()", IsCorrect: false},
					{Text: "execute()", IsCorrect: false},
				},
			},
		},
	}

	test2 := models.Test{
		Title: "Тест по Golang",
		Questions: []models.Question{
			{
				Text: "Что такое Golang (Go)?",
				Answers: []models.Answer{
					{Text: "Язык программирования", IsCorrect: true},
					{Text: "Город в США", IsCorrect: false},
					{Text: "Имя популярной песни", IsCorrect: false},
					{Text: "Программа для редактирования изображений", IsCorrect: false},
				},
			},
			{
				Text: "Как объявить переменную в Golang?",
				Answers: []models.Answer{
					{Text: "var x int", IsCorrect: true},
					{Text: "let x = 5", IsCorrect: false},
					{Text: "declare x as int", IsCorrect: false},
					{Text: "new Variable(x)", IsCorrect: false},
				},
			},
			{
				Text: "Что такое срез (slice) в Golang?",
				Answers: []models.Answer{
					{Text: "Время года", IsCorrect: false},
					{Text: "Динамический массив", IsCorrect: true},
					{Text: "Тип данных", IsCorrect: false},
					{Text: "Структура данных для хранения времени", IsCorrect: false},
				},
			},
			{
				Text: "Как объявить функцию в Golang?",
				Answers: []models.Answer{
					{Text: "function myFunction() { }", IsCorrect: false},
					{Text: "def myFunction() { }", IsCorrect: false},
					{Text: "func myFunction() { }", IsCorrect: true},
					{Text: "method myFunction() { }", IsCorrect: false},
				},
			},
			{
				Text: "Какой оператор используется для цикла в Golang?",
				Answers: []models.Answer{
					{Text: "for", IsCorrect: true},
					{Text: "loop", IsCorrect: false},
					{Text: "while", IsCorrect: false},
					{Text: "repeat", IsCorrect: false},
				},
			},
			{
				Text: "Каким образом вы можете создать структуру (struct) в Golang?",
				Answers: []models.Answer{
					{Text: "new Struct()", IsCorrect: false},
					{Text: "type MyStruct struct { }", IsCorrect: true},
					{Text: "create Struct { }", IsCorrect: false},
					{Text: "declare struct MyStruct { }", IsCorrect: false},
				},
			},
			{
				Text: "Как объявить карту (map) в Golang?",
				Answers: []models.Answer{
					{Text: "map{}", IsCorrect: false},
					{Text: "make(map)", IsCorrect: false},
					{Text: "map[KeyType]ValueType{}", IsCorrect: true},
					{Text: "createMap()", IsCorrect: false},
				},
			},
			{
				Text: "Что такое 'defer' в Golang?",
				Answers: []models.Answer{
					{Text: "Отложенная функция", IsCorrect: true},
					{Text: "Ошибочное ключевое слово", IsCorrect: false},
					{Text: "Условный оператор", IsCorrect: false},
					{Text: "Тип данных", IsCorrect: false},
				},
			},
			{
				Text: "Как создать срез (slice) в Golang?",
				Answers: []models.Answer{
					{Text: "new Slice()", IsCorrect: false},
					{Text: "slice := []int{}", IsCorrect: true},
					{Text: "createSlice()", IsCorrect: false},
					{Text: "var slice []int", IsCorrect: false},
				},
			},
			{
				Text: "Как объявить константу в Golang?",
				Answers: []models.Answer{
					{Text: "const x = 5", IsCorrect: true},
					{Text: "let x = 5", IsCorrect: false},
					{Text: "var x = 5", IsCorrect: false},
					{Text: "define x as 5", IsCorrect: false},
				},
			},
		},
	}

	// Сохранение теста и связанных данных в базе данных
	models.DB.Create(&test1)
	models.DB.Create(&test2)
	con.JSON(http.StatusOK, gin.H{})
}

func Test(c *gin.Context) {
	_, ok := CheckToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Недостаточно прав"})
		return
	}
	testID := c.Param("id")

	var test models.Test
	if err := models.DB.Preload("Questions.Answers").Where("id = ?", testID).First(&test).Error; err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	// убираем IsCorrect из ответа
	var strippedTest models.Test
	strippedTest.ID = test.ID
	strippedTest.Title = test.Title
	strippedTest.Questions = make([]models.Question, len(test.Questions))
	for i, q := range test.Questions {
		strippedTest.Questions[i].ID = q.ID
		strippedTest.Questions[i].Text = q.Text
		strippedTest.Questions[i].TestID = q.TestID
		strippedTest.Questions[i].Answers = make([]models.Answer, len(q.Answers))
		for j, a := range q.Answers {
			strippedTest.Questions[i].Answers[j].ID = a.ID
			strippedTest.Questions[i].Answers[j].Text = a.Text
			// Пропускаем поле IsCorrect
		}
	}

	c.JSON(200, strippedTest)
}

func GetAllTest(c *gin.Context) {
	_, ok := CheckToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Недостаточно прав"})
		return
	}
	var test []models.Test

	models.DB.Find(&test)
	c.JSON(200, test)
}

func SubmitTest(c *gin.Context) {
	id, ok := CheckToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Недостаточно прав"})
		return
	}
	testID := c.Param("testID")

	var test models.Test
	if err := models.DB.Preload("Questions.Answers").Where("id = ?", testID).First(&test).Error; err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	var userAnswers []models.UserAnswer
	if err := c.ShouldBindJSON(&userAnswers); err != nil {
		c.JSON(400, gin.H{"msg": "Неверный формат ответов пользователя"})
		return
	}

	// Подсчет правильных ответов
	correctAnswers := 0
	for _, ua := range userAnswers {
		for _, q := range test.Questions {
			if q.ID == ua.QuestionID {
				for _, a := range q.Answers {
					if a.ID == ua.AnswerID && a.IsCorrect {
						correctAnswers++
					}
				}
			}
		}
	}

	// Запишите результат пользователя в базу данных
	userTestResult := models.UserTestResult{
		UserID: id,
		TestID: test.ID,
		Score:  correctAnswers,
	}

	models.DB.Create(&userTestResult)

	// Связываем ответы пользователя с результатом теста
	for _, ua := range userAnswers {
		ua.UserTestResultID = userTestResult.ID
		models.DB.Create(&ua)
	}

	c.JSON(200, gin.H{"correctAnswers": correctAnswers})
}
