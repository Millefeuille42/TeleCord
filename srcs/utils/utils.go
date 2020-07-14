package utils

import (
	"log"
	"sort"
	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Hang() {
	for {
		time.Sleep(1 * time.Second)
	}
}

func logError(err error) {
	if err != nil {
		log.Print(err)
	}
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func rankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv

	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}
