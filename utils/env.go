package utils

import (
	"os"
	"strings"
)

func keyValue(item string) (key, val string) {
	splits := strings.Split(item, "=")
	key = splits[0]
	val = splits[1]
	return
}

func Env() map[string]string {
	data := os.Environ()

	items := make(map[string]string)
	for _, item := range data {
		key, val := keyValue(item)
		items[key] = val
	}

	return items
}

func Roll(a map[string]string) []string {
	env := make([]string, 0)
	for k, v := range a {
		env = append(env, k+"="+v)
	}
	return env
}
