package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lopesmarcello/money-transfer/internal/utils/validator"
)

func EncodeJSON[T any](w http.ResponseWriter, r *http.Request, status int, data T) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to code json %w", err)
	}
	return nil
}

func DecodeValidJSON[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("error while decoding JSON:\n%w", err)
	}

	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, fmt.Errorf("problems decoding valid JSON:\n%T\n%d problems", problems, len(problems))
	}

	return data, nil, nil
}

func DecodeJSON[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("error while decoding JSON:\n%w", err)
	}

	return data, nil
}

func JSONmsg(args ...string) map[string]string {
	msg := make(map[string]string)

	for index, value := range args {
		// i0 = key
		// i0 = value
		// ie: "message", "something was successful", "foo", "bar"...
		// starts on he second position of the array
		if index%2 != 0 {
			key := args[index-1]
			msg[key] = value
		}
	}

	return msg
}
