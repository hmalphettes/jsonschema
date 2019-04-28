package jsonschema

import "strings"

type keyValue struct {
	key   string
	value string
}

// scanTags returns a list of key=values. Skips the flags as long as there are at the beginning
func scanTagKVs(line string) []*keyValue {
	if line == "" {
		return nil
	}
	currentKV := &keyValue{}
	keyValues := []*keyValue{
		currentKV,
	}
	toks := strings.Split(line, "=")
	for i, tok := range toks {
		if currentKV.key == "" {
			// Skip the flags that are not key=value: flag1,flag2,key=value
			n := strings.Split(tok, ",")
			currentKV.key = n[len(n)-1]
			continue
		}
		if i+1 == len(toks) {
			currentKV.value = tok
			continue
		}
		j := strings.LastIndex(tok, ",")
		if j == -1 {
			currentKV.value = tok
			continue
		}
		if j > 0 {
			// left of the last comma is the value
			currentKV.value = tok[0:j]
		}
		// right of the last comma is the next key
		currentKV = &keyValue{
			key: tok[j+1 : len(tok)],
		}
		keyValues = append(keyValues, currentKV)
	}
	return keyValues
}
