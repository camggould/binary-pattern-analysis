package main

import (
	"fmt"

	"github.com/camggould/binary-pattern-analysis/tools/analysis"
	"github.com/camggould/binary-pattern-analysis/tools/utils"
)

func visualizeMonogramFrequencyAnalysis(ciphertext string) {
	analysis.PrintCharacterFrequencies(ciphertext)
}

func main() {
	path := "data/alpha.txt"
	ciphertext, err := utils.GetCipherText(path)

	if err != nil {
		fmt.Printf("Failed to get ciphertext from file %s. %w\n", path, err)
	}

	visualizeMonogramFrequencyAnalysis(ciphertext)
}