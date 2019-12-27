package generate

import (
	m "github.com/verlandz/elasticsearch/scripts/models"
	"github.com/verlandz/elasticsearch/scripts/utils"
)

// Data Source:
// https://www.eurogamer.net/articles/2018-12-21-pokemon-go-type-chart-effectiveness-weaknesses

var listAttribute = []string{
	"Normal",
	"Fighting",
	"Flying",
	"Poison",
	"Ground",
	"Rock",
	"Bug",
	"Ghost",
	"Steel",
	"Fire",
	"Water",
	"Grass",
	"Electric",
	"Psychic",
	"Ice",
	"Dragon",
	"Fairy",
	"Dark",
}

const (
	MIN_ATTR, MAX_ATTR   = 1, 4
	MIN_POWER, MAX_POWER = 1, 100
)

var lenListAttribute int

func init() {
	lenListAttribute = len(listAttribute)
}

func GenAttribute() []m.Attribute {
	res := []m.Attribute{}
	n_attr := utils.RandInt(MIN_ATTR, MAX_ATTR)
	visited := map[string]bool{}

	for i := 0; i < n_attr; i++ {
		var _name string
		for {
			_name = listAttribute[utils.RandInt(1, lenListAttribute-1)]
			if !visited[_name] {
				visited[_name] = true
				break
			}
		}

		temp := m.Attribute{
			Name:  _name,
			Power: utils.RandInt(MIN_POWER, MAX_POWER),
		}
		res = append(res, temp)
	}

	return res
}
