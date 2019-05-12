package functions

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/sha3"

	"../../p3/data"
)

//SignUp for sign up
func SignUp(password string, addr string) (string, error) {
	timestamp := time.Now().Unix()
	timestampStr := strconv.Itoa(int(timestamp))
	usernameHash := md5.Sum([]byte(timestampStr))
	// usernameHash := sha3.Sum256([]byte(timestampStr))
	username := hex.EncodeToString(usernameHash[:])
	tx := data.Transaction{}
	pwHash := sha3.Sum256([]byte(password))
	tx.Password = hex.EncodeToString(pwHash[:])
	tx.Balance = 10
	tx.Index = 1
	tx.IsFinished = true
	txJSON := tx.MarshalTx()
	sendMap := map[string]string{username: txJSON}
	sendJSON, err := json.Marshal(sendMap)
	// fmt.Println("json ", string(sendJSON))
	_, err = http.Post(addr+"/writeApi/2", "application/json", bytes.NewBuffer([]byte(string(sendJSON))))
	return username, err
}
