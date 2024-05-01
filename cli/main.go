package cli

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"example/todizer"
)

var inputVar string
var outputVar string


func init() {
    flag.StringVar(&inputVar, "input", "", "input file to read todo list from [default: stdin]")
    flag.StringVar(&outputVar, "output", "", "output file to write your work to [default: stdout]")
}

func Execute() {
    var err error
    flag.Parse()

    inputVar, err = cleanFile(inputVar)
    if err != nil {
        log.Fatal(err)
    }

    outputVar, err = cleanFile(outputVar)
    if err != nil {
        log.Fatal(err)
    }
    
    var input io.ReadCloser
    var output io.ReadWriter

    output = bytes.NewBuffer(nil)

    input = os.Stdin
    if inputVar != "" {
        // Is it a good idea to keep a file descriptor
        // for a long period of time ?
        input, err = os.Open(inputVar)
        if err != nil {
            log.Fatal(err)
        }
        defer input.Close()
    } 

    if input == os.Stdin {
        fmt.Println("Enter your todos:")    
    }

    menu := todizer.New(input, output)
    if err := todizer.Execute(menu); err != nil {
        log.Fatal(err)
    }

    // flush data into the specified output
    //  the actual output 
    actualOutput := os.Stdout
    if outputVar != "" {
        actualOutput, err = os.Create(outputVar)
        if err != nil {
            log.Fatal(err)
        }
        defer actualOutput.Close()
    } 
    sc := bufio.NewScanner(output)
    for sc.Scan() {
        fmt.Fprintln(actualOutput, sc.Text())
    }
}


func cleanFile(path string) (string, error) {
    var err error
    if path != "" {

       path, err = expandUser(path)  
       if err != nil {
           return "", err
       }

       if !filepath.IsAbs(path) {
         return "", errors.New("")
       }

    }
    return path, nil
}

func expandUser(path string) (string, error) {
    if strings.HasPrefix(path, "~") {
        userHomeDir, err := os.UserHomeDir() 
        if err != nil {
            return "", err
        }
        path = strings.TrimPrefix(path, "~")
        return filepath.Join(userHomeDir, path), nil
    }
    return path, nil
}
