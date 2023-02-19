package jq

import (
	"github.com/itchyny/gojq"
)

func RunJq(query string, input any) (any, error) {
	q, err := gojq.Parse(query)
	if err != nil {
		return nil, err
	}
	iter := q.Run(input)

	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return nil, err
		}

		return v, nil
	}
	return "", nil
}
