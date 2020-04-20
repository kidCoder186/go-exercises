package protein

import (
	"errors"
	"strings"
)

var (
	ErrStop        = errors.New("invalid codon stop")
	ErrInvalidBase = errors.New("invalid base")
	c2pMap         = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine", "UUC": "Phenylalanine",
		"UUA": "Leucine", "UUG": "Leucine",
		"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
		"UAU": "Tyrosine", "UAC": "Tyrosine",
		"UGU": "Cysteine", "UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
	}
)

// FromCodon returns name of protein of corresponding codon
func FromCodon(s string) (res string, err error) {
	res = c2pMap[s]
	if strings.Compare(res, "STOP") == 0 {
		res, err = "", ErrStop
	} else if strings.Compare(res, "") == 0 {
		res, err = "", ErrInvalidBase
	}
	return
}

// FromRNA returns list of protein of a given RNA
func FromRNA(s string) ([]string, error) {
	res := []string{}
	existProteinMap := map[string]bool{}
	for i := 0; i < len(s); i += 3 {
		if i+3 > len(s) {
			return res, ErrInvalidBase
		}
		protein := c2pMap[string(s[i:i+3])]
		if strings.Compare(protein, "") == 0 {
			return res, ErrInvalidBase
		} else if strings.Compare(protein, "STOP") == 0 {
			return res, nil
		}
		if !existProteinMap[protein] {
			existProteinMap[protein] = true
			res = append(res, protein)
		}
	}
	return res, nil
}
