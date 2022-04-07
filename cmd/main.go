package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstName()
		fmt.Printf("ご予約者一覧: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("チケットは全て完売しました。")
		}
	} else {
		if !isValidName {
			fmt.Println("ファーストネーム・ラストネームのどちらか、あるいは両方が短すぎます。")
		}
		if !isValidEmail {
			fmt.Println("メールアドレスに＠がありません。＠を含めて入力してください。")
		}
		if !isValidTicketNumber {
			fmt.Println("入力されたチケットの枚数が無効です。")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("%v予約アプリにようこそ！\n", conferenceName)
	fmt.Printf("合計で%v枚のチケットがあり、%v枚はまだ予約可能です。\n", conferenceTickets, remainingTickets)
	fmt.Println("チケットはこちらからお求めください。")
	fmt.Printf("\n")
}

func getFirstName() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("ファーストネームを入力してください : ")
	fmt.Scan(&firstName)

	fmt.Println("ラストネームを入力してください : ")
	fmt.Scan(&lastName)

	fmt.Println("メールアドレスを入力してください : ")
	fmt.Scan(&email)

	fmt.Println("チケットの枚数を入力してください : ")
	fmt.Scan(&userTickets)
	fmt.Printf("\n")

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("ご予約者 %v\n", bookings)

	fmt.Printf("%v・%v 様、%v枚のご予約ありがとうございます。\n", firstName, lastName, userTickets)
	fmt.Printf("%vに、確認メールを送信します。ご確認ください。\n", email)
	fmt.Printf("チケットは残り%v枚となりました。\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v・%v", userTickets, firstName, lastName)
	fmt.Println("###################################")
	fmt.Printf("チケットを送信しました:\n %v \n宛先メールアドレス : %v\n", ticket, email)
	fmt.Println("###################################")
	wg.Done()
}
