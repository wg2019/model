package main

import (
	"github.com/urfave/cli/v2"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// GetInputParam 获取参数
func GetInputParam(c *cli.Context) (param *InputStruct, err error) {
	param = new(InputStruct)
	ok := true
	param.User, ok = c.Value(User).(string)
	if !ok {
		return nil, NewError("userName invalid")
	}
	param.Password, ok = c.Value(Password).(string)
	if !ok {
		return nil, NewError("password invalid")
	}
	param.Host, ok = c.Value(Host).(string)
	if !ok {
		return nil, NewError("host invalid")
	}
	param.Port, ok = c.Value(Port).(int64)
	if !ok {
		return nil, NewError("port invalid")
	}
	param.Database, ok = c.Value(Database).(string)
	if !ok {
		return nil, NewError("database invalid")
	}
	param.TableName, ok = c.Value(TableName).(string)
	if !ok {
		return nil, NewError("tableNames invalid")
	}
	param.Package, ok = c.Value(Package).(string)
	if !ok {
		return nil, NewError("package invalid")
	}
	param.Output, ok = c.Value(Output).(string)
	if !ok {
		return nil, NewError("output invalid")
	}
	return param, nil
}

// FirstCharacter 首字母
func FirstCharacter(name string) string {
	return strings.ToLower(name)[:1]
}

// CamelizeStr 大驼峰
func CamelizeStr(s string, upperCase bool) string {
	if len(s) == 0 {
		return s
	}
	s = replaceInvalidChars(s)
	var result string
	words := strings.Split(s, "_")
	for i, word := range words {
		if upper := strings.ToUpper(word); commonInitialismMap[upper] {
			result += upper
			continue
		}
		if i > 0 || upperCase {
			result += camelizeWord(word)
		} else {
			result += word
		}
	}
	return result
}

func camelizeWord(word string) string {
	runes := []rune(word)
	for i, r := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}

func replaceInvalidChars(str string) string {
	str = strings.ReplaceAll(str, "-", "_")
	str = strings.ReplaceAll(str, " ", "_")
	return strings.ReplaceAll(str, ".", "_")
}

func FileName(str string) string {
	str = replaceInvalidChars(str)
	return strings.ReplaceAll(str, "_", "") + ".go"
}

var commonInitialismMap = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"ETA":   true,
	"GPU":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"OS":    true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
	"OAuth": true,
}

// SaveFile 保存文件
func SaveFile(dirPath, fileName string, text []byte) error {
	_ = os.MkdirAll(dirPath, 0777)
	file, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		return err
	}
	defer file.Close()
	p, err := format.Source(text)
	if err != nil {
		return err
	}
	_, err = file.Write(p)
	return err
}
