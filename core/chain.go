package core

import (
	"errors"

	"github.com/tanayshandilya/blockchain/core/block"
)

type BlockChain struct {
	Version     string         `json:"version"`
	LastUpdated string         `json:"lastUpdated"`
	Ledger      []*block.Block `json:"ledger"`
}

func (c *BlockChain) Initialize() error {
	genesis := new(block.Block)
	genesis.CreateGenesis()
	c.Version = ChainVersion
	c.LastUpdated = "1996-05-31 00:00:00.0000000 +0000 UTC"
	c.fill(genesis)
	return nil
}

func (c *BlockChain) fill(blocks ...*block.Block) {
	c.Ledger = blocks
}

func (c *BlockChain) GetLatestBlock() *block.Block {
	return c.Ledger[len(c.Ledger)-1]
}

func (c *BlockChain) GetGenesisBlock() *block.Block {
	return c.Ledger[0]
}

func (c *BlockChain) GetBlock(hash string) (*block.Block, error) {
	for _, b := range c.Ledger {
		if b.Hash == hash {
			return b, nil
		}
	}
	return new(block.Block), errors.New("block not found")
}

func (c *BlockChain) AddBlock(block *block.Block) error {
	if len(block.Consensus) != 0 {
		return errors.New("Block contains consensus data. Hence cannot be added as a new block.")
	}
	c.Ledger = append(c.Ledger, block)
	return nil
}
