package main

import (
	"fmt"
	"os"
	// "strings"
	// "bufio"
	"strconv"
	"bytes"
	"compress/zlib"
	"io"
	"reflect"
)

func extract(zr io.Reader) (io.Reader, error) {
	return zlib.NewReader(zr)
}

func cPress(fBuffer []byte) []uint8 {
	var buf bytes.Buffer
	zWriter := zlib.NewWriter(&buf)
	zWriter.Write(fBuffer)
	zWriter.Close()

	return buf.Bytes()
}

func cMeta(f *os.File) []byte {
	scf := new(os.File)
	*scf = *f
	scFinfo, err := scf.Stat()
	if err != nil {
		fmt.Println("NG")
	}
	meta := "blob "
	size := strconv.FormatInt(scFinfo.Size(), 10)
	meta += size
	return []byte(meta)
}

func main() {
	// f_path := "./B.py"
	f_path := "/mnt/c/Users/81701/Desktop/AtCoder/ABC/001/A_DFS_RE.py"
	f, err := os.Open(f_path)
	if err != nil {
		fmt.Println("NG")
	}
	defer f.Close()
	dc := make([]byte, 1024)
	(*f).Read(dc)

	fmt.Println(reflect.TypeOf(f))
	meta := cMeta(f)

	// fmt.Println(string(meta))

	meta = append(meta, 0)
	dc = append(meta, dc...)

	// fmt.Println(string(dc))

	Pressed := cPress(dc)
	fmt.Println(string(Pressed))

	/** ファイルに書き込み */
	wf, err := os.Create("./sample")
	if err != nil {
		fmt.Println("NG")
	}
	defer wf.Close()
	count, err := wf.Write(Pressed)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to write file")
	}

	fmt.Printf("write %d bytes\n", count)

	/** 下のコードは解凍用のコード*/
	// r, err := extract(&buf)
	// if err != nil {
	// 	fmt.Println("NG")
	// }
	// c := make([]byte, 1024)
	// r.Read(c)

	// fmt.Println(string(c))

}
