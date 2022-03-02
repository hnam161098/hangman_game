package main

import (
	"fmt"
	"math/rand"
	"strings"
	"test/constant"
	"time"
)

func generateRandomIndex(in []string) int {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(len(in))
	return result
}

func checkExistAnwser(ans string, list []string) bool {
	for _, v := range list {
		if ans == v {
			return true
		}
	}
	return false
}

func getKeyword() string {
	keyword_list := []string{"python", "golang", "java", "ruby", "nodejs"}
	keyword := keyword_list[generateRandomIndex(keyword_list)]
	return keyword
}

func main() {
	var anwser string
	var life int = 5

	guess := getKeyword()

	list_anwsers := make([]string, 0)
	word_anwsers := make([]string, 0)

	for i := 0; i <= len(guess)-1; i++ {
		word_anwsers = append(word_anwsers, "-")
		list_anwsers = append(list_anwsers, "-")
	}

	for {
		fmt.Println("Từ cần đoán: ", word_anwsers)
		fmt.Println("Máu còn lại: ", life)
		fmt.Print("Câu trả lời của bạn: ")
		fmt.Scan(&anwser)

		fmt.Println("=====================")

		if anwser == "quit" {
			fmt.Println(constant.Thanks)
			break
		}

		if checkExistAnwser(anwser, word_anwsers) {
			fmt.Println(constant.Exist)
			life -= 1
		}

		if checkExistAnwser(anwser, strings.Split(guess, "")) {
			for i := range guess {
				if anwser == string(guess[i]) {
					list_anwsers[i] = anwser
					word_anwsers[i] = anwser
				}
			}
		} else if len(anwser) > 1 {
			fmt.Println(constant.Err)
			life -= 1
		} else {
			fmt.Println(constant.Wrong)
			life -= 1
		}

		if strings.Join(word_anwsers, "") == guess {
			fmt.Println("đáp án hoàn chỉnh là: ", guess)
			fmt.Println(constant.Win)
			break
		}

		if life == 0 {
			fmt.Println(constant.Lose)
			break
		}

	}

}
