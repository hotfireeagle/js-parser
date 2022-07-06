package scanner

import (
	"bufio"
	"errors"
	"io"
	"jsj/util"
	"os"
)

type Scanner struct {
	path        string
	file        *os.File
	reader      *bufio.Reader
	lineColumn  int
	lineRow     int
	lineContent string
	isLastLine  bool
}

func New(absolutePath string) *Scanner {
	fileObj, err := os.Open(absolutePath)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fileObj)

	s := &Scanner{
		path:   absolutePath,
		file:   fileObj,
		reader: reader,
	}

	s.readLine()

	return s
}

func (s *Scanner) CurrentChar() (byte, error) {
	if s.lineColumn >= len(s.lineContent) {
		if s.isLastLine {
			return 0, util.ErrEof
		} else {
			panic("line coulumn is out of range")
		}
	} else {
		return s.lineContent[s.lineColumn], nil
	}
}

func (s *Scanner) NextChar() (byte, error) {
	if s.lineColumn+1 >= len(s.lineContent) {
		if s.isLastLine {
			return 0, util.ErrEof
		} else {
			s.readLine()
			return s.CurrentChar()
		}
	} else {
		s.lineColumn += 1
		return s.lineContent[s.lineColumn], nil
	}
}

// 对文件进行按行读取
func (s *Scanner) readLine() {
	lc, readErr := s.reader.ReadString('\n')

	if readErr != nil {
		if errors.Is(readErr, io.EOF) {
			s.isLastLine = true
		} else {
			s.Close()
			panic(readErr)
		}
	}

	s.setLineContent(lc)
}

func (s *Scanner) Close() {
	s.file.Close()
}

func (s *Scanner) GetFilePath() string {
	return s.path
}

func (s *Scanner) GetFile() *os.File {
	return s.file
}

func (s *Scanner) GetReader() *bufio.Reader {
	return s.reader
}

func (s *Scanner) setLineContent(str string) {
	s.lineContent = str
	s.lineColumn = 0
	s.lineRow += 1
}
