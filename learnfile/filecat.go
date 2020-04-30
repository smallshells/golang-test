package learnfile

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func cat(r *bufio.Reader){
	for {
		buf,err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		fmt.Fprintf(os.Stdout,"%s",buf)
	}
}

func FileCat(){
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout,"reading from %s failed, err:%v\n",flag.Arg(i),err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
