package functions

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"

	"../../p3/data"
	"golang.org/x/crypto/sha3"
)

//Login for login
func Login(username string, password string, ifLogin *bool, tx *data.Transaction, addr string) error {
	if *ifLogin {
		return errors.New("You already login. ")
	}
	resp, err := http.Get(addr + "/read/" + username)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return errors.New("Cannot find this user, please try later. ")
	} else if resp.StatusCode == 200 {
		transaction := data.Transaction{}
		json.NewDecoder(resp.Body).Decode(&transaction)
		pwHash := sha3.Sum256([]byte(password))
		tryPwHash := hex.EncodeToString(pwHash[:])
		if tryPwHash == transaction.Password {
			if transaction.IsFinished {
				*ifLogin = true
				tx.Balance = transaction.Balance
				tx.Index = transaction.Index
				tx.Password = transaction.Password
				tx.IsFinished = transaction.IsFinished
				// fmt.Println("balance111:", transaction.Balance)
				return nil
			}
			return errors.New("There is unfinished transaction, please try later. ")
		}
		return errors.New("Incorrect password. ")
	}
	return errors.New("Something wrong, please exit. ")
}
