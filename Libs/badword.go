package Libs

import (
	"log"

	"github.com/importcjj/sensitive"
)

type LBadWords struct {
	Filter *sensitive.Filter
}

func NewLBadWords() *LBadWords {
	return &LBadWords{}
}

func (bd *LBadWords) Init(dictfile string) {
	bd.Filter = sensitive.New()
	err := bd.Filter.LoadWordDict(dictfile)
	if err != nil {
		log.Println("Load Baddict Failed", err.Error())
	} else {
		log.Println("Load Baddict Success")
	}
}
