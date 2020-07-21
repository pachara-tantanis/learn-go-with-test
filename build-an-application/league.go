package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing League, %v", err)
	}

	return league, err
}
