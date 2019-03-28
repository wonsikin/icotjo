package parser

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// Parser parser the file and output de result file
func Parser(input, output string, unsorted bool) (err error) {
	file, err := ReadFile(input)
	if err != nil {
		return err
	}
	defer file.Close()

	header, body, err := ReadContent(file)
	if err != nil {
		fmt.Printf("fail when reading input file: %v\n", err)
		return err
	}

	// sort content by key
	if !unsorted {
		err = sortInputFile(file.Name(), header, body)
		if err != nil {
			fmt.Printf("fail when sorting input file: %v\n", err)
			return err
		}
	}

	var mapList = make(map[string]map[string]string)
	for i := 1; i < len(header); i++ {
		var tempMap = make(map[string]string)
		for j := 0; j < len(body); j++ {
			tempMap[body[j][0]] = body[j][i]
		}
		mapList[header[i]] = tempMap
	}

	for k, v := range mapList {
		data, err := jsonMarshal(v, true)
		if err != nil {
			return err
		}

		var out bytes.Buffer
		json.Indent(&out, data, "", "\t")

		outputFile, err := os.Create(path.Join("./", k+".json"))
		if err != nil {
			return err
		}

		out.WriteTo(outputFile)
	}
	return nil
}

// ReadFile read file by input
func ReadFile(input string) (file *os.File, err error) {
	goPaths := filepath.SplitList("GOPATH")
	if len(goPaths) == 0 {
		return nil, errors.New("GOPATH environment variable is not set or empty")
	}

	goRoot := runtime.GOROOT()
	if goRoot == "" {
		return nil, errors.New("GOROOT environment variable is not set or empty")
	}

	absPath, err := filepath.Abs(input)
	if err != nil {
		return nil, err
	}

	file, err = os.Open(absPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// ReadContent read the content of the csv file , handler func(string)
func ReadContent(file *os.File) (header []string, body [][]string, err error) {
	r := csv.NewReader(file)
	// 逐行读取
	records, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	header = records[0]
	body = records[1:]

	return header, body, nil
}

func jsonMarshal(v interface{}, safeEncoding bool) (data []byte, err error) {
	data, err = json.Marshal(v)

	if safeEncoding {
		data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)
		data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
		data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
	}

	return data, err
}
