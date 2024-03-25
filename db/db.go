package db

import "fmt"

type BabyStruct struct {
	// TODO
}

func (b *BabyStruct) ListWords() []string {
	return []string{}
}
func (b *BabyStruct) WordInfo(word string) (*WordInfo, error) {
	return nil, fmt.Errorf("TODO")
}
func (b *BabyStruct) AddWord(word string) (*WordInfo, error) {
	return nil, fmt.Errorf("TODO")
}

func GetBaby(uid string) (*BabyStruct, error) {
	return nil, fmt.Errorf("TODO")
}

type WordInfo struct {
	// TODO
}
