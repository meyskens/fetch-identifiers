package main

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"gopkg.in/src-d/go-billy.v4"
)

var repos = map[string]string{
	//"go-git.coment": "https://github.com/src-d/go-git",
	"freeCodeCamp.coment": "https://github.com/freeCodeCamp/freeCodeCamp",
	"vue.coment":          "https://github.com/vuejs/vue",
	"springboot.coment":   "https://github.com/spring-projects/spring-boot",
	"moby.coment":         "https://github.com/moby/moby",
}

func main() {
	if len(os.Args) > 2 { // ./fetch-comments https://github.com/moby/moby moby.coment
		// fetch a specific repo
		getCommentsOnRepo(os.Args[1], os.Args[2])
		return
	}
	wg := sync.WaitGroup{}
	for name, url := range repos {
		log.Println("Adding", url)
		wg.Add(1)
		go func(url, name string) {
			getCommentsOnRepo(url, name)
			wg.Done()
		}(url, name)
	}
	wg.Wait()
}

func getCommentsOnRepo(url, name string) {
	log.Println("Getting", url)
	outFile, _ := os.Create(name)
	defer outFile.Close()
	repoFiles, fs, err := getFiles(url)
	if err != nil {
		panic(err)
	}

	log.Println("Parsing files in", url)

	for _, file := range repoFiles.GoFiles {
		log.Println("Looking at", file)
		comments, err := listIdentifiers("go", file, fs)
		if err == nil {
			writeComments(outFile, comments)
		} else {
			log.Println(err)
		}
	}

	for _, file := range repoFiles.JavaFiles {
		log.Println("Looking at", file)
		comments, err := listIdentifiers("java", file, fs)
		if err == nil {
			writeComments(outFile, comments)
		} else {
			log.Println(err)
		}
	}

	for _, file := range repoFiles.JavascriptFiles {
		log.Println("Looking at", file)
		comments, err := listIdentifiers("javascript", file, fs)
		if err == nil {
			writeComments(outFile, comments)
		} else {
			log.Println(err)
		}
	}

	for _, file := range repoFiles.PythonFiles {
		log.Println("Looking at", file)
		comments, err := listIdentifiers("python", file, fs)
		if err == nil {
			writeComments(outFile, comments)
		} else {
			log.Println(err)
		}
	}

	for _, file := range repoFiles.PHPFiles {
		log.Println("Looking at", file)
		comments, err := listIdentifiers("php", file, fs)
		if err == nil {
			writeComments(outFile, comments)
		} else {
			log.Println(err)
		}
	}

	for _, file := range repoFiles.RubyFiles {
		log.Println("Looking at", file)
		comments, err := listIdentifiers("ruby", file, fs)
		if err == nil {
			writeComments(outFile, comments)
		} else {
			log.Println(err)
		}
	}

}

func listIdentifiers(lang, fileName string, fs billy.Filesystem) ([]string, error) {
	file, err := fs.Open(fileName)
	if err != nil {
		return nil, err
	}
	content, _ := ioutil.ReadAll(file)
	comments, err := fetchIdentifiers(lang, string(content))
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func writeComments(file *os.File, comments []string) {
	for _, comment := range comments {
		_, err := file.Write([]byte(comment + "\n"))
		if err != nil {
			log.Println(err)
		}
	}
}
