package functions

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/sha3"

	"../../p3/data"
)

//Lucky for getting lottery numbers
func Lucky(userNum []int, addr string, tx *data.Transaction, username string) error {
	// usernameHash := sha3.Sum256([]byte("657a6a6cb8f654ff424b96da5079a48c"))
	usernameHash := sha3.Sum256([]byte(username))
	usernameUint := binary.BigEndian.Uint64(usernameHash[:])
	rand.Seed(int64(usernameUint) + time.Now().UnixNano())
	balance := tx.Balance
	if balance < 2 {
		return errors.New("You have not enough money! ")
	}
	// fmt.Println("now balance: ", balance)
	balance -= 2
	luckyArr := generateLuckyNumber()
	rewardNum := 0
	for i, luckyNum := range luckyArr {
		if userNum[i] == luckyNum {
			rewardNum++
		}
	}
	switch rewardNum {
	case 1:
		balance += 2
	case 2:
		balance += 5
	case 3:
		balance += 20
	case 4:
		balance += 100
	case 5:
		balance += 500
	}
	fmt.Fprintln(os.Stdout, "Lottery numbers: ", luckyArr)
	fmt.Fprintln(os.Stdout, "You have ", rewardNum, " right numbers. ")
	tx.Balance = balance
	index := tx.Index
	tx.Index = index + 1
	txJSON := tx.MarshalTx()
	sendMap := map[string]string{username: txJSON}
	sendJSON, err := json.Marshal(sendMap)
	// fmt.Println("json ", string(sendJSON))
	_, err = http.Post(addr+"/writeApi/2", "application/json", bytes.NewBuffer([]byte(string(sendJSON))))
	return err
}

func generateLuckyNumber() []int {
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(10) + 1
	}
	return arr
}
