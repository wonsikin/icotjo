package parser

import (
	"encoding/csv"
	"os"
	"sort"
)

// ByKey implements sort.Interface for [][]string based on
// the Age field.
type ByKey [][]string

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i][0] < a[j][0] }

func sortInputFile(fileName string, header []string, body [][]string) error {
	// remove file at the first
	err := os.Remove(fileName)
	if err != nil {
		return err
	}

	// sort body
	sort.Sort(ByKey(body))
	content := make([][]string, 0)
	content = append(content, header)
	content = append(content, body...)

	// create a same name file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	w.WriteAll(content)

	return nil
}
