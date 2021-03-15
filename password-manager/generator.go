package main

import (
	"crypto/rand"
	"io"
	"math"
	"math/big"
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var (
	All = []string{LowerLetters, UpperLetters, Digits, Symbols}

	NoSymbols = []string{LowerLetters, UpperLetters, Digits}

	NumbersOnly = []string{Digits}
)

type config struct {
	flags  []string
	length int32
}

type PasswordGeneration struct {
	tags      []string
	config    config
	reader    io.Reader
	generated string
}

func getLimit(tag string, numLetters, others int32) int32 {
	if tag == LowerLetters || tag == UpperLetters {
		return numLetters
	}
	return others
}

func (p PasswordGeneration) genPass() (generatedPassword string, err error) {
	numLetters := int32(math.Round(float64(p.config.length) * 0.8))
	numDigits := int32(math.Round(float64(p.config.length) * 0.1))
	if numDigits+numDigits+numLetters > p.config.length {
		numLetters--
	}
	for _, tag := range p.tags {
		tagLimit := int(getLimit(tag, numLetters, numDigits))
		for i := 0; i < tagLimit; i++ {
			r, err := getRandom(p.reader, tag)
			if err != nil {
				return "", err
			}
			generatedPassword, err = appendRandom(p.reader, generatedPassword, r)
			if err != nil {
				return "", err
			}
		}
	}
	return
}

func appendRandom(reader io.Reader, base, append string) (result string, err error) {
	if base == "" {
		result = append
		return
	}
	n, err := rand.Int(reader, big.NewInt(int64(len(base)+1)))
	if err != nil {
		return
	}
	i := n.Int64()
	result = base[0:i] + append + base[i:]
	return
}

func getRandom(reader io.Reader, group string) (r string, err error) {
	rn, err := rand.Int(reader, big.NewInt(int64(len(group))))
	if err != nil {
		return
	}
	i := rn.Int64()
	r = string(group[i])
	return
}
