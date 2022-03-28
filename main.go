package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello world\n"))
// }
func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
	//os.Stdout.Write([]byte("hey\n"))

	// conn, err := net.Dial("tcp", "ascii.jp:80")
	// if err != nil {
	// 	panic(err)
	// }
	// conn.Write([]byte("GET /HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	// io.Copy(os.Stdout, conn)

	// file, err := os.Create("old.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// file.Write([]byte("os file example\n"))
	// file.Close()

	// file, err := os.Open("old.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// file2, _ := os.Create("copy.txt")
	// io.Copy(file2, file)

	// file, _ := os.Create("rand.txt")
	// defer file.Close()
	// io.CopyN(file, rand.Reader, 1024)

	file, _ := os.Create("sample.zip")
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	a, _ := zipWriter.Create("a.txt")
	io.Copy(a, strings.NewReader("一つ目のファイルです"))
	b, _ := zipWriter.Create("b.txt")
	io.Copy(b, strings.NewReader("二つ目のファイルです"))

}
