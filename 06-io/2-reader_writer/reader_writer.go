package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./reader_writer.go")
	if err != nil {
		log.Fatal(err)
	}
	b, err := get(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	f, err = os.Create("./reader_writer_copy.go")
	if err != nil {
		log.Fatal(err)
	}
	err = store(f, b)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}

func store(w io.Writer, b []byte) error {
	// Здесь должна быть некотороая логика.
	// После обработки данных мы их записываем туда,
	// куда хочет вызывающий код.
	_, err := w.Write(b)
	return err
}

func get(r io.Reader) ([]byte, error) {
	var b []byte
	var buf = make([]byte, 10)
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return nil, err
		}
		b = append(b, buf...)
	}
	// Здесь какая-то логика.
	return b, nil
}
