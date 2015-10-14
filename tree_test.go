package tree

import (
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

func BenchmarkBuildTree(b *testing.B) {
	b.ReportAllocs()

	r := strings.NewReader(strings.Join([]string{
		"hoge	1",
		"fuga	2",
		"kuke	3",
	}, "\n"))

	for i := 0; i < b.N; i++ {
		BuildTree(r)
	}
}
