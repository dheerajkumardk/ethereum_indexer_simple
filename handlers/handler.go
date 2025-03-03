package handlers

import (
	"example.com/indexer-wallet-txns/database"
	"github.com/gofiber/fiber/v2"
)

func GetBlockByBlockNumber(c *fiber.Ctx) error {
	// get the block number from params
	blockNumber := c.Params("blockNumber")

	db := database.BlockDB
	if db == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database is not initialized",
		})
	}

	var block database.Block
	// search for blockNumber in db
	db.Find(&block, "block_number = ?", blockNumber)
	if block.BlockHash == "" {
		c.Status(500).Send([]byte("No block found with given BlockNumber"))
	}
	return c.JSON(block)
}

func GetBlockByBlockHash(c *fiber.Ctx) error {
	// get the blockHash from params
	blockHash := c.Params("blockHash")
	db := database.BlockDB
	if db == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database is not initialized",
		})
	}

	var block database.Block
	// search for blockHash in db
	db.Find(&block, "block_hash = ?", blockHash)
	if block.BlockHash == "" {
		c.Status(500).Send([]byte("No block found with given BlockHash"))
	}
	return c.JSON(block)
}

func GetAllBlocks(c *fiber.Ctx) error {
	db := database.BlockDB
	if db == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database is not initialized",
		})
	}

	var blocks []database.Block
	db.Find(&blocks)
	return c.JSON(blocks)
}

func AddBlock(c *fiber.Ctx) error {
	db := database.BlockDB
	if db == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database is not initialized",
		})
	}

	block := new(database.Block)
	if err := c.BodyParser(block); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	result := db.Create(&block)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create block",
		})
	}
	return c.JSON(block)
}
