package tree

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"github.com/armon/go-radix"
	"io"
	"os"
	"strings"
)

// BuildTree creates radix tree from provided io.Reader stream.
//
// r expected as 2 column tab separated format, for example (k, v).
func BuildTree(r io.Reader) (*radix.Tree, error) {
	tree := radix.New()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b := scanner.Text()
		kv := strings.Split(b, "\t")
		// ignoring invalid line.
		if len(kv) != 2 {
			continue
		}
		tree.Insert(kv[0], kv[1])
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tree, nil
}

// BuildTreeFromFile creates tree from given path.
func BuildTreeFromFile(path string) (*radix.Tree, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	return BuildTree(fp)
}

// LoadTreeFromGobFile creates tree by gob file.
func LoadTreeFromGobFile(path string) (*radix.Tree, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	// buffering mapped radix.tree
	var b map[string]interface{}
	dec := gob.NewDecoder(fp)
	if err := dec.Decode(&b); err != nil {
		return nil, err
	}
	return radix.NewFromMap(b), nil
}

// ExportTreeToGobFile exports tree as gob file.
func ExportTreeToGobFile(t *radix.Tree, path string) error {
	// export as gob file
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(t.ToMap()); err != nil {
		return err
	}

	fp, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	defer fp.Close()
	if err != nil {
		return err
	}
	if _, err := fp.Write(b.Bytes()); err != nil {
		return err
	}
	return nil
}
