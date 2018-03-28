/*
 * Filename: gitcredential.go
 * Author: Chris Drexler <ckolumbus@ac-drexler.de>
 *
 * Copyright (c) 2018 Chris Drexler <ckolumbus@ac-drexler.de>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 *
 */

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	urllib "net/url"
	"os"
)

var (
	version        = "0.1.0"
	version_commit = "undefined"
)

////////////////////////////////////////////////////////////////
/// adapted from https://github.com/jprichardson/readline-go
/// MIT License : Copyright (c) 2013 JP Richardson
////////////////////////////////////////////////////////////////
func ReadLine(reader io.Reader) string {
	buf := bufio.NewReader(reader)
	line, err := buf.ReadBytes('\n')
	for err == nil {
		line = bytes.TrimRight(line, "\n")
		if len(line) > 0 {
			if line[len(line)-1] == 13 { //'\r'
				line = bytes.TrimRight(line, "\r")
			}
			return (string(line))
		}
		line, err = buf.ReadBytes('\n')
	}

	if len(line) > 0 {
		return (string(line))
	}
	return ""
}

////////////////////////////////////////////////////////////////

var Config struct {
	Silent bool
	Help   bool
}

// note, that variables are pointers

func init() {
	flag.BoolVar(&Config.Silent, "s", false, "silent operation, not console output/help")
	flag.BoolVar(&Config.Help, "h", false, "print help")
}

func printTitle(out io.Writer) {
	fmt.Fprint(out, "Construct GIT credential entry accroding to: <scheme>://<user>:<pwd>@<host>/[<path>]\n")
}
func printHelp(out io.Writer) {
	fmt.Fprint(out, "\n")
	printTitle(out)
	fmt.Fprint(out, "\n")
	fmt.Fprint(out, "   scheme : https,http (default=https)\n")
	fmt.Fprint(out, "   user   : user name \n")
	fmt.Fprint(out, "   pwd    : password\n")
	fmt.Fprint(out, "   host   : hostname (including port if needed)\n")
	fmt.Fprint(out, "   path   : path to repo (default=\"\")\n")
	fmt.Fprint(out, "\n")
	fmt.Fprint(out, "Example : gitcred >> credentials.git\n")
	fmt.Fprint(out, "\n")
	fmt.Fprintf(out, "Version : %s, Commit : %s\n", version, version_commit)
	fmt.Fprint(out, "Source  : https://github.com/ckolumbus/gitcred\n")
	fmt.Fprint(out, "Copyright (c) 2018 Chris Drexler <ckolumbus@ac-drexler.de>\n")
	fmt.Fprint(out, "This software is released under the MIT License. (https://opensource.org/licenses/MIT)\n")
	fmt.Fprint(out, "\n")
}

func main() {
	flag.Parse()

	// print help
	if Config.Help {
		printHelp(os.Stdout)
		os.Exit(0)
	}

	// either write to stderr
	// or to nowhere in silent mode
	out := os.Stderr
	if Config.Silent {
		out = nil
	}

	printTitle(out)
	fmt.Fprint(out, "scheme: ")
	scheme := ReadLine(os.Stdin)
	fmt.Fprint(out, "user  : ")
	user := ReadLine(os.Stdin)
	fmt.Fprint(out, "pwd   : ")
	pwd := ReadLine(os.Stdin)
	fmt.Fprint(out, "host  : ")
	host := ReadLine(os.Stdin)
	fmt.Fprint(out, "path  : ")
	path := ReadLine(os.Stdin)
	if scheme == "" {
		scheme = "https"
	}

	userinfo := urllib.UserPassword(user, pwd)
	resultUrl := &urllib.URL{Scheme: scheme, User: userinfo, Host: host, Path: path}
	fmt.Println(resultUrl.String())
}
