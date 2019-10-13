package example

import "github.com/AMan4Technology/DataStructure/internal"

func KOf(key string) internal.K {
    return k(key)
}

type k string

func (k k) Key() string {
    return string(k)
}
