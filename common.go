/*
 Auther : F.W
 Create time  2019/7/25
*/
package utils

import (
	"encoding/gob"
	"os"
	"path"
)

func DumpsToFile(dir string, object interface{}) error {

	file, err := os.Create(path.Join(dir, "db.dump"))
	if err == nil {
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(object)
	}
	file.Close()
	return err
}

func LoadsToObject(dir string, object interface{}) error {
	file, err := os.Open(path.Join(dir, "db.dump"))
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&object)
	}
	file.Close()
	return err

}
