package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

// default variables
const (
	TmplFileName = "tmpl.html"
	CSVFileName  = "data.csv"
	DistFileName = "dist.html"
)

type option struct {
	SkipLine int
	NewLine  bool
	BR       bool
	Space    bool
	Help     bool
}

func parseFlags() option {
	op := option{}
	flag.IntVar(&op.SkipLine, "S", 0, "シート(csv)の先頭からint行目までをスキップします。")
	flag.BoolVar(&op.NewLine, "n", false, "改行を取り除きます。")
	flag.BoolVar(&op.BR, "b", false, "改行を<br/>に置換します。")
	flag.BoolVar(&op.Space, "s", false, "全角スペースを取り除きます。")
	flag.BoolVar(&op.Help, "h", false, "使い方を表示します。")
	flag.Parse()
	return op
}

func checkFileExists(fn string) {
	if _, err := os.Stat(fn); err != nil {
		log.Fatal(fn + "が配置されていません。")
	}
}

func createTmplate(fn string) *template.Template {
	fbs, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(fn + "を開けませんでした。\n" + err.Error())
	}

	var hbs []byte
	hbs = append(hbs, strings.TrimSpace(string(fbs))...)
	hbs = append(hbs, "\n"...)

	tbs := []byte("{{ range . }}\n")
	tbs = append(tbs, string(hbs)...)
	tbs = append(tbs, "{{ end }}"...)

	t, err := template.New("template").Parse(string(tbs))
	if err != nil {
		log.Fatal(fn + "の解析に失敗しました。\n" + err.Error())
	}

	return t
}

func parseCSVFile(fn string, op option) []map[string]string {
	const ABC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	cf, err := os.Open(fn)
	if err != nil {
		log.Fatal(fn + "を開けませんでした。\n" + err.Error())
	}
	defer cf.Close()

	r := csv.NewReader(cf)
	r.FieldsPerRecord = -1
	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(fn + "の解析に失敗しました。\n" + err.Error())
	}

	var dataSet []map[string]string
	for idx, items := range record {
		if op.SkipLine-1 >= idx {
			continue
		}

		data := make(map[string]string)
		for key, item := range items {

			if op.Space {
				item = strings.Replace(item, "　", "", -1)
			}
			if op.NewLine {
				item = strings.Replace(item, "\n", "", -1)
			}
			if op.BR {
				item = strings.Replace(item, "\n", "<br/>", -1)
			}

			data[string(ABC[key])] = item
		}
		sn := "0" + strconv.Itoa(idx+1-op.SkipLine)
		data["SN"] = sn[len(sn)-2:]
		dataSet = append(dataSet, data)
	}

	return dataSet
}

func main() {
	op := parseFlags()

	if op.Help {
		fmt.Println(helpText)
		os.Exit(0)
	}

	checkFileExists(TmplFileName)
	checkFileExists(CSVFileName)

	w, err := os.Create(DistFileName)
	if err != nil {
		log.Fatal(DistFileName + "の作成に失敗しました。\n" + err.Error())
	}
	defer w.Close()

	d := parseCSVFile(CSVFileName, op)
	t := createTmplate(TmplFileName)

	if err = t.Execute(w, d); err != nil {
		log.Fatal(DistFileName + "への書き込みに失敗しました。\n" + err.Error())
	}

	fmt.Println(DistFileName + "を作成しました。")
}
