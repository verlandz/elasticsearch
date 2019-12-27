package generate

import (
	"github.com/verlandz/elasticsearch/scripts/utils"
)

const (
	MIN_WORDS     = 1
	MAX_WORDS     = 3
	MIN_LEN_WORDS = 3
	MAX_LEN_WORDS = 7
)

var visited map[string]bool

func init() {
	visited = map[string]bool{} //uniq
}

func GenName() string {
	res := ""
	n_words := utils.RandInt(MIN_WORDS, MAX_WORDS)

	for i := 0; i < n_words; i++ {
		if i != 0 {
			res += " "
		}

		len_words := utils.RandInt(MIN_LEN_WORDS, MAX_LEN_WORDS)

		for j := 0; j < len_words; j++ {
			if j == 0 {
				res += string(utils.RandInt('A', 'Z'))
			} else {
				res += string(utils.RandInt('a', 'z'))
			}
		}
	}

	return res
}
