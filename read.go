package main

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dywoq/genshinapi/character"
	"github.com/dywoq/genshinapi/statistics"
)

// readCharacters reads all found characters in data/characters folder,
// and returns a slice of them. If there are any encountered errors,
// the function returns them instead of nil.
func readCharacters() ([]*character.Character, error) {
	data := []*character.Character{}
	err := filepath.WalkDir("./data/characters", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".json") {
			ch, err := processCharacterFile(path)
			if err != nil {
				return err
			}
			data = append(data, ch)
		}
		return nil
	})

	if err != nil {
		return []*character.Character{}, err
	}
	return data, nil
}

func processCharacterFile(file string) (*character.Character, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return &character.Character{}, nil
	}
	c := &character.Character{}
	err = json.Unmarshal(f, &c)
	if err != nil {
		return &character.Character{}, nil
	}
	if !character.Correct(c) {
		return &character.Character{}, statistics.ErrCheckFailed
	}
	return c, nil
}
