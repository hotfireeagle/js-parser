package utils

import "errors"

var EOL byte = 10 // ASCII对应换行符，ASCII中的0对应文件的最后一行，如果存在最后一行的话

var EOL_ERR = errors.New("EOL")

var EOF_ERR = errors.New("EOF")
