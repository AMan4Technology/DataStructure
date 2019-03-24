package example

import "DataStructure/internal"

func KOf(key string) internal.K {
    return k(key)
}

type k string

func (k k) Key() string {
    return string(k)
}
