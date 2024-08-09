package helper

import (
	"io"
	"os"
	"fmt"
)

func CopyImage(bfsrc string, src string, file io.Reader) error {
	fmt.Println(file)
	// menghapus image sebelumnya
	if err := os.Remove("././static/src/images/" + bfsrc); err != nil {
		return err
	}
	
	// membuat image yang di upload kembali
	outfile, err := os.Create("././static/src/images/" + src)
	if err != nil {
		return err
	}
	defer outfile.Close()

	// mencopy
	if _, err := io.Copy(outfile, file); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return err
}