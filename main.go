package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/**
* 根据.vscode/settings.json -> file.exclude
* 来创建.hide文件夹并生成隐藏文件的软连接
 */

func trimComment(data []byte) []byte {
	text := string(data)
	lines := strings.Split(text, "\n")
	newData := make([]string, 0, len(lines))

	for i := 0; i < len(lines); {
		line := lines[i]
		// 消除 /**/
		if start := strings.Index(line, "/*"); start != -1 {
			var t = line
			for i < len(lines) {
				end := strings.Index(t[start:], "*/")
				if end != -1 {
					newData = append(newData, t[:start]+t[start+end+2:])
					i++
					break
				} else {
					i++
					line = lines[i]
					t += line
				}
			}
			continue
		}
		// 消除 //
		if idx1 := strings.Index(line, "//"); idx1 != -1 {
			newData = append(newData, line[:idx1])
		} else {
			newData = append(newData, line)
		}
		i++
	}
	return []byte(strings.Join(newData, "\n"))
}

func fmtData(data []byte) []byte {

	text := string(trimComment(data))
	start := 0
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case ',':
			start = i
		case ' ', '\n', '\r':
		case '}':
			if start != 0 {
				i -= 1
				text = text[:start] + text[start+1:]
			}
		default:
			start = 0
		}
	}
	return []byte(text)
}

func getSettings() (map[string]any, string, error) {
	for cur, _ := os.Getwd(); cur != "/"; cur = filepath.Dir(cur) {
		path := cur + "/.vscode/settings.json"
		_, err := os.Stat(path)
		if err == nil {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				return nil, cur, err
			}
			// 定义转换存储的对象
			var jsonData map[string]any

			e := json.Unmarshal(fmtData(data), &jsonData)

			if e != nil {
				fmt.Println(e)
				return nil, cur, err
			}
			excludes, ok := jsonData["files.exclude"].(map[string]any)
			if ok {
				return excludes, cur, nil
			}
		}
	}
	return nil, "", nil
}

// 获取当前文件夹下的文件
func getCurrentDir() ([]string, error) {
	path, _ := os.Getwd()
	dirs, err := ioutil.ReadDir(path)
	filenames := make([]string, 0, len(dirs))
	if err != nil {
		return filenames, err
	}

	for _, d := range dirs {
		filenames = append(filenames, path+"/"+d.Name())
	}

	return filenames, nil
}

// 转换设置，* -> .*?
func transferSettings(settings map[string]any) []string {
	regs := make([]string, 0, len(settings))
	for k, v := range settings {
		b, _ := v.(bool)
		if b {
			regs = append(regs, strings.ReplaceAll(k, "*", ".*?"))
		}
	}
	return regs
}

// 获取匹配的路径
func getMathPath(filenames, regs []string) []string {
	mathPath := make([]string, 0, len(filenames))
	for _, filename := range filenames {
		for _, reg := range regs {
			match, _ := regexp.MatchString(strings.ReplaceAll(reg, "*", ".*"), filename)
			if match {
				mathPath = append(mathPath, filename)
			}
		}
	}
	return mathPath
}

// 创建软连接
func createSymLink(filenames []string) error {
	cur, _ := os.Getwd()
	_, err := os.Stat(cur + "/.hide")
	if err != nil {
		// 创建文件夹
		e := os.Mkdir(cur+"/.hide", 0755)
		if e != nil {
			return e
		}
	}
	// 便利创建软连接
	for _, path := range filenames {
		idx := strings.LastIndex(path, "/")
		name := path[idx+1:]
		linkPath := cur + "/.hide/" + name
		if _, err := os.Stat(linkPath); err != nil {
			e := os.Symlink(path, linkPath)
			fmt.Println("success: " + linkPath)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

func main() {

	// 获取设置文件
	settings, _, err := getSettings()
	if err != nil {
		fmt.Println(err)
		return
	}
	regs := transferSettings(settings)

	filenames, _ := getCurrentDir()
	mathPath := getMathPath(filenames, regs)
	err = createSymLink(mathPath)
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println("success!")
	}
}
