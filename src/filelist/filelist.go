/*
#!/usr/bin/env gorun
# _author_ lbbxsxlz@gmail.com
*/

package filelist

import (
	"errors"
	"io/ioutil"
	"os"
)
/*
func DeleteFileAndDir(path string) (error) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {
			return err
		}
		if f.IsDir() {
			fmt.Println("fileName:", f.Name())
			os.RemoveAll(path + "/" + f.Name())
		} else {
			fmt.Println("fileName:", f.Name())
			os.Remove(path + "/" + f.Name())
		}
		return nil
	})

	if err != nil {
		return errors.New("filepath.Walk() fail")
	}

	return nil
}
*/

func DeleteFileAndDir(path string) (error){
	var err error
	files, err := ioutil.ReadDir(path)
	if err != nil {
		errors.New("ReadDir fail")
	}

	for _, fi := range files {
		if fi.IsDir() {
			os.RemoveAll(path + "/" + fi.Name())
		} else {
			os.Remove(path + "/" + fi.Name())
		}
	}

	return nil
}