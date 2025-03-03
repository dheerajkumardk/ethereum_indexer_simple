package database

import "gorm.io/gorm"

// Struct to store block data in db
type Block struct {
	gorm.Model
	BlockNumber  int64 `json:"blockNumber"    gorm:"column:block_number"`
	BlockHash    string `json:"blockHash"     gorm:"column:block_hash"`
	NumberOfTxns int64	`json:"numberOfTxns"  gorm:"column:number_of_txns"`
	Miner 		 string `json:"miner" 		  gorm:"column:miner"`
}
