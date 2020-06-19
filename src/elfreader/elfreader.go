/*
#!/usr/bin/env gorun
# _author_ lbbxsxlz@gmail.com
*/

package elfreader

import (
	"debug/elf"
	"errors"
	"io/ioutil"
	"log"
)

func CreateUuidFile(progfile, sectionname, filename string)(error) {
	ef, err := elf.Open(progfile)
	if err != nil {
		log.Printf("open elf fail \n")
		return errors.New("open elf fail \n")
	}
	defer ef.Close()

	section := ef.Section(sectionname)
	if section == nil {
		log.Printf("get uuid section fail \n")
		return errors.New("get uuid section fail \n")
	}

	data, err:= section.Data()
	if err != nil {
		log.Printf("get uuid data fail \n")
		return errors.New("get get uuid data fail \n")
	}

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		log.Printf("WriteFile data fail \n")
		return errors.New("WriteFile data fail \n")
	}

	ef.Close()
	return nil
}
