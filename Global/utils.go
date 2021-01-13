package Global

import "os"

func IsExistFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
