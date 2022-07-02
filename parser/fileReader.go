// 读取文件
// 必须提供一个取出第一个字符的方法

package parser

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type FileReader struct {
	// 文件读取器
	reader *bufio.Reader

	file *os.File

	// 要读取的文件的绝对路径
	filePath string

	// 当前读取的文件行内容
	lineRaw string

	// 当前读取的行号
	lineNum int

	// 当前读取行的列索引
	lineColumn int

	// 是否已经读取到文件末尾
	hasEOF bool
}

func FileReaderConstructor(fp string) *FileReader {
	fileObj, openErr := os.Open(fp)

	if openErr != nil {
		panic(openErr)
	}

	fr := &FileReader{
		filePath: fp,
		file:     fileObj,
		reader:   bufio.NewReader(fileObj),
	}

	// 文件一打开即开始读取一行
	fr.readline()
	return fr
}

func (fr *FileReader) readline() {
	lineContent, readErr := fr.reader.ReadString('\n')

	if readErr != nil {
		if errors.Is(readErr, io.EOF) {
			fr.hasEOF = true
		} else {
			fr.CloseFile()
			panic(readErr)
		}
	}

	fr.lineRaw = lineContent
}

func (fr *FileReader) GetFilePath() string {
	return fr.filePath
}

func (fr *FileReader) GetLineRaw() string {
	return fr.lineRaw
}

func (fr *FileReader) GetLineNum() int {
	return fr.lineNum
}

func (fr *FileReader) GetLineColumn() int {
	return fr.lineColumn
}

func (fr *FileReader) GetHasEOF() bool {
	return fr.hasEOF
}

func (fr *FileReader) CloseFile() {
	fr.file.Close()
}
