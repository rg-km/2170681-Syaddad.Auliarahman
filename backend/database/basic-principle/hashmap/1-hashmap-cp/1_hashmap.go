package main

import "errors"

type HashMap struct {
	m map[int]string
}

func NewHashMap() *HashMap {
	return &HashMap{
		m: make(map[int]string),
	}
}

func (h *HashMap) Get(key int) (string, error) {
	val, ok := h.m[key]
	if !ok {
		return "", errors.New("not found")
	}
	return val, nil
}

func (h *HashMap) Put(key int, value string) error {
	if _, ok := h.m[key]; ok {
		return errors.New("key exists in the hashmap")
	}
	h.m[key] = value
	return nil
}

func (h *HashMap) GetRange(from, to int) ([]string, error) {
	if from > to {
		return nil, errors.New("from > to")
	} else if from < 0 {
		return nil, errors.New("from < 0")
	} else if to > len(h.m) {
		return nil, errors.New("to > len(h.m)")
	} else if to < 0 {
		return nil, errors.New("to < 0")
	} else if from == to {
		return []string{h.m[from]}, nil
	}
	// var res []string
	// for i := from; i <= to; i++ {
	// 	res = append(res, h.m[i])
	// }
	return nil, nil // TODO: replace this
}
