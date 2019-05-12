package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../p3/data"
	"./functions"
)

var ADDR string

// var balance = -1

//IfLogin if the user login
// var IfLogin = false
var IfLogin = false

var UserName string

//TX transaction
var TX = data.Transaction{}

func init() {
	TX.Balance = 10
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	args1 := os.Args[1]
	ADDR = "http://localhost:" + args1
	fmt.Println("this is the node", ADDR)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "sum": //test case
		aStr := arrCommandStr[1]
		a, _ := strconv.ParseInt(aStr, 10, 64)
		bStr := arrCommandStr[2]
		b, _ := strconv.ParseInt(bStr, 10, 64)
		fmt.Fprintln(os.Stdout, a+b)
		return nil

	case "signup":
		password := arrCommandStr[1]
		username, err := functions.SignUp(password, ADDR)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, username)
		return nil
	case "login":
		UserName = arrCommandStr[1]
		password := arrCommandStr[2]
		err := functions.Login(UserName, password, &IfLogin, &TX, ADDR)
		// fmt.Println("balance2: ", TX.Balance)
		// fmt.Println("index", TX.Index)
		// fmt.Println("is login", IfLogin)

		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, "Login success. ")
		return nil
	case "lucky":
		if !IfLogin {
			return errors.New("You need Login first! ")
		}
		if len(arrCommandStr) < 6 {
			return errors.New("You should give five numbers from 1-10")
		}
		userNum := []int{}
		for i := 1; i <= 5; i++ {
			n, _ := strconv.Atoi(arrCommandStr[i])
			if n < 1 || n > 10 {
				return errors.New("You should give five numbers from 1-10")
			}
			userNum = append(userNum, n)
		}
		err := functions.Lucky(userNum, ADDR, &TX, UserName)
		return err

	case "balance":
		if !IfLogin {
			return errors.New("You need Login first! ")
		}
		balance := TX.Balance
		fmt.Fprintln(os.Stdout, "Balance: ", balance)
		return nil

	case "logout":

	}

	return errors.New("Command not found. ")
}
