package learnfile

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src ,err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err : %v.\n",srcName, err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName,os.O_WRONLY|os.O_CREATE,0644)
	if err != nil {
		fmt.Printf("open %s failed, err : %v.\n",dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func FileCopy(){
	_, err := CopyFile("./learnfile/dst.txt","./learnfile/src.txt")
	if err != nil {
		fmt.Printf("open file failed, err : %v.\n", err)
		return
	}
	fmt.Println("copy done!")
}