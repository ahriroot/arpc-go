package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// {'version': '1.0', 'package': [{'language': 'python', 'path': 'api'}], 'procedures': [{'name': 'GetUserV1', 'request': 'RequestV1', 'response': 'ResponseV1'}], 'param': {'RequestV1': [{'name': 'user_id', 'type': 'integer', 'index': 1}], 'ResponseV1': [{'name': 'user_id', 'type': 'integer', 'index': 1}, {'name': 'username', 'type': 'string', 'index': 2}]}}

type Package struct {
	Language string `json:"language"`
	Path     string `json:"path"`
}

type Procedures struct {
	Name     string `json:"name"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Param struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Index int    `json:"index"`
}

type ArpcMeta struct {
	Version    string             `json:"version"`
	Package    []Package          `json:"package"`
	Procedures []Procedures       `json:"procedures"`
	Param      map[string][]Param `json:"param"`
}

func CompileArpc(path string) (*ArpcMeta, error) {
	var err error

	var arpc_meta ArpcMeta = ArpcMeta{}

	// 当前解析到的行号
	var line_num int = 0
	// 正在解析 package
	var handle_package bool = false
	// // 正在解析 procedure
	// var handle_procedures bool = false
	// // 正在解析 param
	// var handle_param bool = false
	// // 当前解析 param 名
	// var param_name string = ""

	var file *os.File
	file, err = os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("打开文件出错：", err)
	}
	defer file.Close()

	var content []byte
	var buf []byte = make([]byte, 128)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}
		content = append(content, buf[:n]...)
	}
	var lines = strings.Split(string(content), "\n")

	for _, line := range lines {
		line_num++
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if handle_package {
			if strings.HasPrefix(line, "}") {
				handle_package = false
				continue
			} else if strings.HasPrefix(line, "{") {
				continue
			} else {
				// 正则匹配 language: path
				var reg = regexp.MustCompile(`^(.*):\s*(.*)`)
				var match = reg.FindStringSubmatch(line)
				fmt.Println(match)
				if len(match) == 3 {
					fmt.Println(match[0], match[1], match[2])
					var p = Package{
						Language: match[1],
						Path:     match[2],
					}
					arpc_meta.Package = append(arpc_meta.Package, p)
				}
			}
		} else {
			if strings.HasPrefix(line, "package") {
				if arpc_meta.Package != nil {
					return nil, errors.New(fmt.Sprintf("File [%s]\n\tline [%d]: Repeated package area", path, line_num))
				}
				handle_package = true
				arpc_meta.Package = make([]Package, 0)
				continue
			}

		}
	}

	return &arpc_meta, nil
}

func Compile(path string) {
	var res, _ = CompileArpc(path)
	fmt.Printf("%#v \n", res)
}
