package data

import "encoding/json"

type Transaction struct {
	Password   string `json:"password"`
	Balance    int    `json:"balance"`
	Index      int    `json:"index"`
	IsFinished bool   `json:"isfinished"`
}

func UnmarshalTx(txJSON string) Transaction {
	tx := Transaction{}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(txJSON), &dat); err != nil {
		panic(err)
	}
	tx.Password = dat["password"].(string)
	tx.Balance = int(dat["balance"].(float64))
	tx.Index = int(dat["index"].(float64))
	tx.IsFinished = dat["isfinished"].(bool)
	return tx
}

func (tx *Transaction) MarshalTx() string {
	cacheContent := make(map[string]interface{})
	cacheContent["password"] = tx.Password
	cacheContent["balance"] = tx.Balance
	cacheContent["index"] = tx.Index
	cacheContent["isfinished"] = tx.IsFinished
	str, err := json.Marshal(cacheContent)
	if err != nil {
		panic(err)
	}
	return string(str)

}
