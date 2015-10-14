package tree

import (
	"bytes"
	"encoding/gob"
	"github.com/armon/go-radix"
	"github.com/suzuken/dummy"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestBuildTree(t *testing.T) {
	r := strings.NewReader(strings.Join([]string{
		"hoge	1",
		"fuga	2",
		"kuke	3",
	}, "\n"))

	tree, err := BuildTree(r)
	if err != nil {
		t.Fatalf("fail to build tree: %s", err)
	}
	v, ok := tree.Get("hoge")
	if !ok {
		t.Fatalf("fail to fetch item by hoge")
	}
	if "1" != v.(string) {
		t.Fatalf("cannot fetch 1 by hoge.")
	}
}

func TestExportAndLoadTree(t *testing.T) {
	r := strings.NewReader(strings.Join([]string{
		"hoge	1",
		"fuga	2",
		"kuke	3",
	}, "\n"))

	tree, err := BuildTree(r)
	if err != nil {
		t.Fatalf("fail to build tree: %s", err)
	}

	path := "test1.gob"
	if err := ExportTreeToGobFile(tree, path); err != nil {
		t.Fatalf("fail to export tree %s", err)
	}
	loadedTree, err := LoadTreeFromGobFile(path)
	if err != nil {
		t.Fatalf("fail to load gob file. %s", err)
	}
	if !reflect.DeepEqual(tree, loadedTree) {
		t.Fatal("loaded tree doesn't equal original tree.")
	}
	if err := os.Remove(path); err != nil {
		t.Fatalf("fail to remove temporary gob file: %s", err)
	}
}

// genTreeTSV generate tsv for building tree.
func genTreeTSV(length, keyLen, valueLen int) io.Reader {
	g := dummy.NewGenerator()
	var buf bytes.Buffer
	for i := 0; i < length; i++ {
		buf.WriteString(g.String(keyLen))
		buf.WriteString("\t")
		buf.WriteString(g.String(valueLen))
		buf.WriteString("\n")
	}
	return &buf
}

func benchmarkBuildTree(length int, b *testing.B) {
	b.ReportAllocs()

	r := genTreeTSV(length, 10, 3)
	for i := 0; i < b.N; i++ {
		BuildTree(r)
	}
}

func BenchmarkBuildTree10000(b *testing.B)   { benchmarkBuildTree(10000, b) }
func BenchmarkBuildTree100000(b *testing.B)  { benchmarkBuildTree(100000, b) }
func BenchmarkBuildTree1000000(b *testing.B) { benchmarkBuildTree(1000000, b) }

// genTree generate radix.Tree which has specified length
func genTree(length, keyLen, valueLen int) *radix.Tree {
	g := dummy.NewGenerator()
	t := radix.New()
	for i := 0; i < length; i++ {
		t.Insert(g.String(keyLen), g.String(valueLen))
	}
	return t
}

func benchmarkBuildTreeFromGob(length int, b *testing.B) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(genTree(length, 10, 3).ToMap()); err != nil {
		b.Fatalf("encode err: %s", err)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		BuildTreeFromGob(&buf)
	}
}

func BenchmarkBuildTreeFromGob10000(b *testing.B)   { benchmarkBuildTreeFromGob(10000, b) }
func BenchmarkBuildTreeFromGob100000(b *testing.B)  { benchmarkBuildTreeFromGob(100000, b) }
func BenchmarkBuildTreeFromGob1000000(b *testing.B) { benchmarkBuildTreeFromGob(1000000, b) }
