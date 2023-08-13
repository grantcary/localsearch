package main

import (
	"bufio"
	"flag"
	"fmt"
	"localsearch/find"
	"localsearch/store"
	"os"
	user "os/user"
	"strings"
)

func main() {
	userdir, _ := user.Current()
	basedir := userdir.HomeDir + "/Desktop/"
	dir := flag.String("dir", basedir, "Specify directory to search")

	flag.Parse()

	filenames, filedirs := store.Scrape(dir)
	fmt.Println("* File System Loaded *")
	bigram := store.Bigram(filenames)
	unigram := store.Unigram(filenames)

	for k, v := range bigram {
		fmt.Println(k, v)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Search: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		found := find.Search(text, unigram, bigram)
		for _, val := range found {
			fmt.Println(filedirs[val])
		}
	}

}
