package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"slices"
)

type JsonFile struct {
	UserJArray []User `json:"users"`
}

type User struct {
	Email              string `json:"email"`
	NumOfMealsToSelect int    `json:"numOfmealsToSelect"`
	MealJArray         []Meal `json:"meals"`
}

type Meal struct {
	Name              string   `json:"name"`
	IngrediantsJArray []string `json:"ingrediants"`
}

func main() {
	userJsonFile, errReadingJson := readJsonFile()
	if errReadingJson != nil {
		fmt.Println("Error gathering file from disk: ", errReadingJson)
		return
	}

    userArray := userJsonFile.UserJArray;

    for i := 0; i < len(userArray); i++ {
        meal1, err := selectMeals(userArray[i])
        if err != nil {
            fmt.Println("Was unable to succesfully select meal for users", err)
        }
        fmt.Println(meal1)
    }

}

func readJsonFile() (JsonFile, error) {
	contents, err := os.ReadFile("./resources/userList.json")
	if err != nil {
		return JsonFile{}, err
	}

	data := JsonFile{}

	err = json.Unmarshal(contents, &data)

	if err != nil {
		fmt.Println(err)
	}

	return data, nil

}

func selectMeals(usersData User) ([]Meal, error)  {
    numOfUsersMeal := len(usersData.MealJArray)
	numOfmealsToSelect := usersData.NumOfMealsToSelect

	randomMealsToBeSelected, err := generateUniqueRandomIntegers(numOfUsersMeal, numOfmealsToSelect)
	if err != nil {
        return []Meal{}, err
	}


    for i := 0; i < rand

    //get the meal objects
   return []Meal{}, nil
}

func generateUniqueRandomIntegers(numberRange int, amountToGenerate int) ([]int, error) {
	if amountToGenerate > numberRange {
		return []int{}, errors.New("Amount asked to generate, is higher than the number of user meals added")
	}
	uniqueInts := []int{}

	i := 0
    // I might have an off by one error in this loop, we'll see
	for i < amountToGenerate {
		currentInt := rand.Intn(numberRange)
		if !slices.Contains(uniqueInts, currentInt) {
			uniqueInts = append(uniqueInts, currentInt)
			i++
		}
	}

	return uniqueInts, nil
}

func sendEmail() {
	from := ""
	password := ""

	to := []string{""}

	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("This is a test email message")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully")
}
