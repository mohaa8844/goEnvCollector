package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	projectPath:="C:/go/src/student"
	envs:=[]string{}
	re:=regexp.MustCompile(`os.Getenv\(["]?(.*?)["]?\)`)

	err := filepath.Walk(projectPath,
		func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			content,err:=ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println(err)
			}
			items:=re.FindAllString(string(content),-1)
			for _,v :=range items {
				v=strings.TrimPrefix(v,"os.Getenv(")
				v=strings.TrimPrefix(v,	`"`)
				v=strings.TrimSuffix(v,	`)`)
				v=strings.TrimSuffix(v,	`"`)
				if !Contains(envs,v){
					envs=append(envs,v)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}


	fi, err := os.OpenFile("config.env", os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
	}
	for _,env :=range envs{
		_,err:=fi.Write([]byte(env+"\n"))
		if err!=nil {
			fmt.Println(err)
		}
	}
	fi.Close()
}
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}