package database

import "gorm.io/gorm"

// Struct to store block data in db
type Block struct {
	gorm.Model
	BlockNumber  string `json:"blockNumber"`
	BlockHash    string `json:"blockHash"`
	NumberOfTxns int64	`json:"numberOfTxns"`
	Miner 		 string `json:"miner"`
}
