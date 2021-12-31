package core

import (
	"errors"
)

type BlockChain struct {
	Version     string   `json:"version"`
	LastUpdated string   `json:"lastUpdated"`
	Ledger      []*Block `json:"ledger"`
}

func (c *BlockChain) Initialize() error {
	genesis := new(Block)
	genesis.createGenesis()
	c.Version = ChainVersion
	c.LastUpdated = "1996-05-31 00:00:00.0000000 +0000 UTC"
	c.fill(genesis)
	return nil
}

func (c *BlockChain) fill(blocks ...*Block) {
	c.Ledger = blocks
}

func (c *BlockChain) GetLatestBlock() *Block {
	return c.Ledger[len(c.Ledger)-1]
}

func (c *BlockChain) GetGenesisBlock() *Block {
	return c.Ledger[0]
}

func (c *BlockChain) GetBlock(hash string) (*Block, error) {
	for _, b := range c.Ledger {
		if b.Hash == hash {
			return b, nil
		}
	}
	return new(Block), errors.New("block not found")
}
