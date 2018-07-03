// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var basicMessages = findMessages(".")
var coreMessages = findMessages("core")

func main() {
	makeWrapper("dfproto", ".")

	plugins, err := filepath.Glob("plugins/*")
	if err != nil {
		panic(err)
	}

	for _, p := range plugins {
		makeWrapper(filepath.Base(p), p)
	}
}

func makeWrapper(packageName, prefix string) {
	protos, err := filepath.Glob(filepath.Join(prefix, "*.proto"))
	if err != nil {
		panic(err)
	}

	var plugin string
	var rpcDeclarations [][3]string
	var pluginMessages []string

	for i, p := range protos {
		name, decls, messages := scanProto(p)
		if i != 0 && plugin != name {
			panic("plugin name mismatch in " + packageName + ": " + plugin + " != " + name)
		}
		plugin = name
		rpcDeclarations = append(rpcDeclarations, decls...)
		pluginMessages = append(pluginMessages, messages...)
	}

	f, err := os.Create(filepath.Join(prefix, "wrapper.gen.go"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if plugin != "" {
		fmt.Fprintf(f, `// Package %s contains protocol buffer definitions for the %s DFHack plugin.
//
// See https://dfhack.readthedocs.io/en/latest/docs/Plugins.html#%s for more details.
`, packageName, plugin, packageName)
	}
	fmt.Fprintf(f, "package %s\n", packageName)

	writeImports(f, rpcDeclarations, pluginMessages)

	for _, decl := range rpcDeclarations {
		writeWrapper(f, plugin, decl[0], decl[1], decl[2], pluginMessages)
	}
}

func findMessages(prefix string) []string {
	var messages []string

	protos, err := filepath.Glob(filepath.Join(prefix, "*.proto"))
	if err != nil {
		panic(err)
	}

	for _, p := range protos {
		_, _, m := scanProto(p)
		messages = append(messages, m...)
	}

	return messages
}

func scanProto(name string) (plugin string, decls [][3]string, messages []string) {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	var hadPlugin bool

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "// Plugin: ") {
			if hadPlugin {
				panic("multiple plugin declarations in " + name)
			} else {
				hadPlugin = true
				plugin = line[len("// Plugin: "):]
			}
		} else if strings.HasPrefix(line, "// RPC ") {
			decl := strings.Split(line, " ")
			if len(decl) != 7 || decl[3] != ":" || decl[5] != "->" {
				panic("invalid RPC declaration: " + line)
			}
			decls = append(decls, [3]string{decl[2], decl[4], decl[6]})
		} else if strings.HasPrefix(line, "message ") {
			messages = append(messages, strings.TrimRight(strings.Split(line, " ")[1], "{"))
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return
}

func writeImports(w io.Writer, decls [][3]string, pluginMessages []string) {
	if len(decls) == 0 {
		return
	}

	paths := []string{"github.com/BenLubar/dfhackrpc"}
	var addedBasic, addedCore bool
	add := func(message string) {
		for _, m := range pluginMessages {
			if m == message {
				return
			}
		}
		for _, m := range basicMessages {
			if m == message {
				if !addedBasic {
					paths = append(paths, "github.com/BenLubar/dfhackrpc/dfproto")
					addedBasic = true
				}
				return
			}
		}
		for _, m := range coreMessages {
			if m == message {
				if !addedCore {
					paths = append(paths, "github.com/BenLubar/dfhackrpc/dfproto/core")
					addedCore = true
				}
				return
			}
		}
		panic("cannot find message named " + message)
	}

	for _, d := range decls {
		add(d[1])
		add(d[2])
	}

	sort.Strings(paths)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "import (")
	for _, p := range paths {
		fmt.Fprintf(w, "\t%q\n", p)
	}
	fmt.Fprintln(w, ")")
}

func writeWrapper(w io.Writer, plugin, method, input, output string, pluginMessages []string) {
	fmt.Fprintln(w)
	qualifiedName := method
	if plugin != "" {
		qualifiedName = plugin + "::" + method
	}
	fmt.Fprintf(w, "// Call%s is a convenience wrapper around the %s RPC method.\n", method, qualifiedName)
	var inputDecl, outputDecl string
	if input != "EmptyMessage" {
		inputDecl = ", request *" + getPrefix(input, pluginMessages) + input
	}
	if output != "EmptyMessage" {
		outputDecl = "*" + getPrefix(output, pluginMessages) + output + ", "
	}
	fmt.Fprintf(w, "func Call%s(client *dfhackrpc.Client%s) (%sdfhackrpc.CommandResult, error) {\n", method, inputDecl, outputDecl)
	requestVar := "request"
	if input == "EmptyMessage" {
		requestVar = "&request"
		fmt.Fprint(w, "\tvar request dfproto_core.EmptyMessage\n\n")
	}
	fmt.Fprintf(w, "\tvar response %s%s\n\n", getPrefix(output, pluginMessages), output)
	if output == "EmptyMessage" {
		fmt.Fprint(w, "\treturn ")
	} else {
		fmt.Fprint(w, "\trv, err := ")
	}
	if plugin == "" {
		fmt.Fprintf(w, "client.Call(%q, %s, &response)\n", method, requestVar)
	} else {
		fmt.Fprintf(w, "client.CallPlugin(%q, %q, %s, &response)\n", plugin, method, requestVar)
	}
	if output != "EmptyMessage" {
		fmt.Fprintln(w, "\n\treturn &response, rv, err")
	}
	fmt.Fprintln(w, "}")
}

func getPrefix(message string, pluginMessages []string) string {
	for _, m := range pluginMessages {
		if m == message {
			return ""
		}
	}
	for _, m := range basicMessages {
		if m == message {
			return "dfproto."
		}
	}
	for _, m := range coreMessages {
		if m == message {
			return "dfproto_core."
		}
	}
	panic("unknown message: " + message)
}
