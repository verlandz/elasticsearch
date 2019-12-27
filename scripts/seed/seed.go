package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/verlandz/elasticsearch/scripts/elastic"
	gen "github.com/verlandz/elasticsearch/scripts/generators"
	m "github.com/verlandz/elasticsearch/scripts/models"
	"github.com/verlandz/elasticsearch/scripts/utils"
)

const N_ROBOT = 200

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	path := "track.txt"
	utils.CreateFile(path)
	for i := 0; i < N_ROBOT; i++ {
		robot := m.Robot{
			Name:      gen.GenName(),
			Release:   gen.GenRelease(),
			Ability:   gen.GenAbility(),
			Attribute: gen.GenAttribute(),
		}
		elastic.Insert(i, robot)

		// (optional) tracking name
		b, _ := json.Marshal(robot.Name)
		utils.AppendFile(string(b), path)
	}
}
