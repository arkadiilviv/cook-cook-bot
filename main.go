package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var apiUrl string = "https://www.themealdb.com/api/json/v1/1/"

func main() {
	bot, err := botapi.NewBotAPI("5124106193:AAHNyBpcg7OiBaDyUFm2jCB9zE7MrjYQlRE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := botapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	var meals MealStruct
	var isSelect bool = false
	port := os.Getenv("PORT")
	router := gin.New()
	router.Static("/static", "static")
	router.Run(":" + port)
	for update := range updates {
		if update.Message != nil { // If we got a message
			var meal Meal
			messageText := "Error"
			requestText := update.Message.Text
			searchResult := ""

			if isSelect {
				for _, item := range meals.Meals {
					if item.StrMeal == requestText {
						meal = item
						break
					}
				}
				if len(meal.StrMeal) > 0 {
					messageText = GetMealRecipe(meal)
				} else {
					messageText = "Not found"
				}
				isSelect = false
			} else {
				if requestText == "/random" {
					searchResult = SearchRandom()
					isSelect = false
				} else {
					searchResult = Search(requestText)
				}

				json.Unmarshal([]byte(searchResult), &meals)

				if meals.Meals == nil {
					messageText = "No results"
				} else if len(meals.Meals) > 1 {
					messageText = "Select one meal \n" + GetMealListStr(meals)
					isSelect = true
				} else {
					meal = meals.Meals[0]
					messageText = GetMealRecipe(meal)
				}
			}

			msg := botapi.NewMessage(update.Message.Chat.ID, messageText)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func Search(name string) string {
	searchUrl := apiUrl + "search.php?s=" + name
	res, err := http.Get(searchUrl)
	if err != nil {
		log.Panic(err.Error())
	}
	resText, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Panic(err.Error())
	}
	return string(resText)
}

func SearchRandom() string {
	searchUrl := apiUrl + "random.php"
	res, err := http.Get(searchUrl)
	if err != nil {
		log.Panic(err.Error())
	}
	resText, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Panic(err.Error())
	}
	return string(resText)
}

func GetMealListStr(meals MealStruct) string {
	var result string = ""
	for index, meal := range meals.Meals {
		result += strconv.Itoa(index+1) + ". " + meal.StrMeal + "\n"
	}
	return result
}

func GetMealRecipe(meal Meal) string {
	var result string
	result += fmt.Sprintf("%s\n", meal.StrMeal)
	result += "_________________________________\n"
	result += fmt.Sprintf("%s\n", meal.StrCategory)
	result += "_________________________________\n"
	result += fmt.Sprintf("%s\n", meal.StrInstructions)
	result += "_________________________________\n"
	result += GetMealLine(meal.StrIngredient1, meal.StrMeasure1)
	result += GetMealLine(meal.StrIngredient2, meal.StrMeasure2)
	result += GetMealLine(meal.StrIngredient3, meal.StrMeasure3)
	result += GetMealLine(meal.StrIngredient4, meal.StrMeasure4)
	result += GetMealLine(meal.StrIngredient5, meal.StrMeasure5)
	result += GetMealLine(meal.StrIngredient6, meal.StrMeasure6)
	result += GetMealLine(meal.StrIngredient7, meal.StrMeasure7)
	result += GetMealLine(meal.StrIngredient8, meal.StrMeasure8)
	result += GetMealLine(meal.StrIngredient9, meal.StrMeasure9)
	result += GetMealLine(meal.StrIngredient10, meal.StrMeasure10)
	result += GetMealLine(meal.StrIngredient11, meal.StrMeasure11)
	result += GetMealLine(meal.StrIngredient12, meal.StrMeasure12)
	result += GetMealLine(meal.StrIngredient13, meal.StrMeasure13)
	result += GetMealLine(meal.StrIngredient14, meal.StrMeasure14)
	result += GetMealLine(meal.StrIngredient15, meal.StrMeasure15)
	result += GetMealLine(meal.StrIngredient16, meal.StrMeasure16)
	result += GetMealLine(meal.StrIngredient17, meal.StrMeasure17)
	result += GetMealLine(meal.StrIngredient18, meal.StrMeasure18)
	result += GetMealLine(meal.StrIngredient19, meal.StrMeasure19)
	result += GetMealLine(meal.StrIngredient20, meal.StrMeasure20)

	result += "Link: " + meal.StrYoutube

	return result
}

func GetMealLine(ingr string, meas string) string {
	if len(ingr) == 0 {
		return ""
	}
	return fmt.Sprintf("%s: %s\n", ingr, meas)
}

type MealStruct struct {
	Meals []Meal
}

type Meal struct {
	StrMeal         string
	StrCategory     string
	StrArea         string
	StrInstructions string
	StrMealThumb    string
	StrYoutube      string
	StrIngredient1  string
	StrIngredient2  string
	StrIngredient3  string
	StrIngredient4  string
	StrIngredient5  string
	StrIngredient6  string
	StrIngredient7  string
	StrIngredient8  string
	StrIngredient9  string
	StrIngredient10 string
	StrIngredient11 string
	StrIngredient12 string
	StrIngredient13 string
	StrIngredient14 string
	StrIngredient15 string
	StrIngredient16 string
	StrIngredient17 string
	StrIngredient18 string
	StrIngredient19 string
	StrIngredient20 string
	StrMeasure1     string
	StrMeasure2     string
	StrMeasure3     string
	StrMeasure4     string
	StrMeasure5     string
	StrMeasure6     string
	StrMeasure7     string
	StrMeasure8     string
	StrMeasure9     string
	StrMeasure10    string
	StrMeasure11    string
	StrMeasure12    string
	StrMeasure13    string
	StrMeasure14    string
	StrMeasure15    string
	StrMeasure16    string
	StrMeasure17    string
	StrMeasure18    string
	StrMeasure19    string
	StrMeasure20    string
	StrSource       string
}
