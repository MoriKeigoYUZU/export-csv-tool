package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type InputData struct {
	CSV string `json:"csv"`
}

func main() {
	// 標準入力からJSONデータを読み取る
	var inputBuffer bytes.Buffer
	_, err := io.Copy(&inputBuffer, os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read input:", err)
		os.Exit(1)
	}

	// JSONデータを構造体にデコード
	var input InputData
	err = json.Unmarshal(inputBuffer.Bytes(), &input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse JSON:", err)
		os.Exit(1)
	}

	// Base64デコード
	csvData, err := base64.StdEncoding.DecodeString(input.CSV)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to decode Base64:", err)
		os.Exit(1)
	}

	// CSVデータをファイルに書き出し
	outputFile := "output.csv"
	err = os.WriteFile(outputFile, csvData, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write CSV file:", err)
		os.Exit(1)
	}

	fmt.Printf("CSV file successfully saved as %s\n", outputFile)
}
