package common

import (
	"github.com/weili71/go-filex"
	"os"
	"path/filepath"
)

var CurrentPath = filex.NewFile(os.Args[0]).ParentFile().Pathname
var UploadSavePath = filepath.Join(CurrentPath, "upload")
var UploadDir = "/upload"
