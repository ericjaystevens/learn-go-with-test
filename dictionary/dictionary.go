package dictionary

import "errors"

type GoDict map[string]string

func (g GoDict) Search(key string) (string, error) {
	definition, wasFound := g[key]

	if !wasFound {
		return "", errors.New("could not find the word you were looking for")
	} else {
		return definition, nil
	}
}
