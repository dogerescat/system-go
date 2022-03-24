package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	//os.Stdout.Write([]byte("hey\n"))

	// conn, err := net.Dial("tcp", "ascii.jp:80")
	// if err != nil {
	// 	panic(err)
	// }
	// conn.Write([]byte("GET /HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	// io.Copy(os.Stdout, conn)

	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// file.Write([]byte("os file example\n"))
	// file.Close()

}
