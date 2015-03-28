package jsongz

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
)

//combine json marshalling and gzip compression to write gzipped json files
type Writer struct {
	Filename  string
	file      *os.File
	bufwriter *bufio.Writer
	writer    *gzip.Writer
}

func NewWriter(filename string) (*Writer, error) {
	curFile, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	bufwriter := bufio.NewWriter(curFile)
	writer := gzip.NewWriter(bufwriter)

	return &Writer{
		Filename:  filename,
		file:      curFile,
		bufwriter: bufwriter,
		writer:    writer,
	}, nil
}

func (w *Writer) Encode(v interface{}) error {
	enc := json.NewEncoder(w.writer)
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	//push writes to file and close everything
	err = w.writer.Close()
	if err != nil {
		return err
	}
	err = w.bufwriter.Flush()
	if err != nil {
		return err
	}

	w.file.Close()
	return nil
}

//combine json marshalling and gzip compression to write gzipped json files
type Reader struct {
	Filename  string
	file      *os.File
	bufreader *bufio.Reader
	reader    *gzip.Reader
}

func NewReader(filename string) (*Reader, error) {
	curFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	bufreader := bufio.NewReader(curFile)
	reader, err := gzip.NewReader(bufreader)
	if err != nil {
		return nil, err
	}

	return &Reader{
		Filename:  filename,
		file:      curFile,
		bufreader: bufreader,
		reader:    reader,
	}, nil

}

func (r *Reader) Decode(v interface{}) error {
	/*
		gzReader, err := gzip.NewReader()
		if err != nil {
			return err
		}
	*/
	dec := json.NewDecoder(r.reader)
	err := dec.Decode(v)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func ReadFile(filename string, v interface{}) error {
	gzreader, err := NewReader(filename)
	if err != nil {
		return err
	}
	err = gzreader.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(filename string, v interface{}) error {
	gzwriter, err := NewWriter(filename)
	if err != nil {
		return err
	}

	err = gzwriter.Encode(v)
	if err != nil {
		return err
	}
	return nil
}
