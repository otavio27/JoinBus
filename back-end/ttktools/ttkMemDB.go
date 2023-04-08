package ttktools

// Memdb ...
type Memdb struct {
	Memdb map[string][]byte
}

// MemDB returns a MemDB instance
func MemDB() *Memdb {
	t := Memdb{}
	t.Memdb = make(map[string][]byte)
	return &t
}

// Write a key in the memory database
func (t *Memdb) Write(key string, value []byte) {
	t.Memdb[key] = value
}

// Read a key from the memory database
func (t *Memdb) Read(key string) []byte {
	return t.Memdb[key]
}

// Delete a key from memory database
func (t *Memdb) Delete(key string) {
	delete(t.Memdb, key)
}

func (t *Memdb) List() []string {
	ret := make([]string, 0)
	for k := range t.Memdb {
		ret = append(ret, k)
	}
	return ret
}
