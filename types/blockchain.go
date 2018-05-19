package types

// Transactions keeps a sequence of Blocks
type Transactions struct {
	Blocks []*Block
}

// AddBlock saves provided data as a block in the transactions
func (bc *Transactions) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Newtransactions creates a new transactions with genesis Block
func Newtransactions() *Transactions {
	return &Transactions{[]*Block{NewGenesisBlock()}}
}
