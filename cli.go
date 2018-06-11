package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct{}

func (cli *CLI) createBlockchain(address string) {
	fmt.Println(address)
	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Done!")
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" printchain - print all the blocks of the blockchain")
	fmt.Println(" createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		fmt.Println("Args < 2")
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	switch os.Args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic("Run - error on printchain", err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) printChain() {

	// TODO: fix this
	bc := NewBlockchain("")
	defer bc.db.Close()

	bci := bc.Interator()

	for {
		block := bci.Next()
		fmt.Println("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			return
		}
	}
}
