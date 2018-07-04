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

var _, _, basicMessages = scanProtos(".")
var _, _, coreMessages = scanProtos("core")

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

var unwrap = map[[2]string][]struct {
	Field     string
	Type      string
	SetPrefix string
}{
	{"dfproto_core", "EmptyMessage"}: {},
	{"dfproto_core", "IntMessage"}: {
		{
			Field:     "Value",
			Type:      "int32",
			SetPrefix: "&",
		},
	},
	{"dfproto_core", "IntListMessage"}: {
		{
			Field: "Value",
			Type:  "[]int32",
		},
	},
	{"dfproto_core", "StringMessage"}: {
		{
			Field:     "Value",
			Type:      "string",
			SetPrefix: "&",
		},
	},
	{"dfproto_core", "StringListMessage"}: {
		{
			Field: "Value",
			Type:  "[]string",
		},
	},
	{"remotefortressreader", "SingleBool"}: {
		{
			Field:     "Value",
			Type:      "bool",
			SetPrefix: "&",
		},
	},
	{"remotefortressreader", "TiletypeList"}: {
		{
			Field: "TiletypeList",
			Type:  "[]*Titletype",
		},
	},
	{"remotefortressreader", "BuildingList"}: {
		{
			Field: "BuildingList",
			Type:  "[]*BuildingDefinition",
		},
	},
	{"remotefortressreader", "MaterialList"}: {
		{
			Field: "MaterialList",
			Type:  "[]*MaterialDefinition",
		},
	},
	{"remotefortressreader", "UnitList"}: {
		{
			Field: "CreatureList",
			Type:  "[]*UnitDefinition",
		},
	},
	{"remotefortressreader", "PlantList"}: {
		{
			Field: "PlantList",
			Type:  "[]*PlantDef",
		},
	},
	{"remotefortressreader", "CreatureRawList"}: {
		{
			Field: "CreatureRaws",
			Type:  "[]*CreatureRaw",
		},
	},
	{"remotefortressreader", "ArmyList"}: {
		{
			Field: "Armies",
			Type:  "[]*Army",
		},
	},
	{"remotefortressreader", "PlantRawList"}: {
		{
			Field: "PlantRaws",
			Type:  "[]*PlantRaw",
		},
	},
	{"remotefortressreader", "Status"}: {
		{
			Field: "Reports",
			Type:  "[]*Report",
		},
	},
}

func makeWrapper(packageName, prefix string) {
	plugin, rpcDeclarations, pluginMessages := scanProtos(prefix)

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
		writeWrapper(f, plugin, decl[0], decl[1], decl[2], packageName, pluginMessages)
	}
}

func scanProtos(prefix string) (plugin string, declarations [][3]string, messages []string) {
	protos, err := filepath.Glob(filepath.Join(prefix, "*.proto"))
	if err != nil {
		panic(err)
	}

	for i, p := range protos {
		name, decls, msgs := scanProto(p)
		if i != 0 && plugin != name {
			panic("plugin name mismatch in " + prefix + ": " + plugin + " != " + name)
		}
		plugin = name
		declarations = append(declarations, decls...)
		messages = append(messages, msgs...)
	}

	return
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

func qualifiedName(plugin, method string) string {
	if plugin != "" {
		return plugin + "::" + method
	}
	return method
}

func writeWrapper(w io.Writer, plugin, method, input, output, packageName string, pluginMessages []string) {
	fmt.Fprintf(w, "\n// Call%s is a convenience wrapper around the %s RPC method.\n", method, qualifiedName(plugin, method))

	var inputDecl, outputDecl string
	if uw, ok := unwrap[[2]string{getPrefix2(input, packageName, pluginMessages), input}]; ok {
		for _, u := range uw {
			inputDecl += ", " + strings.ToLower(u.Field) + " " + u.Type
		}
	} else {
		inputDecl = ", request *" + getPrefix(input, pluginMessages) + input
	}

	if uw, ok := unwrap[[2]string{getPrefix2(output, packageName, pluginMessages), output}]; ok {
		for _, u := range uw {
			outputDecl += u.Type + ", "
		}
	} else {
		outputDecl = "*" + getPrefix(output, pluginMessages) + output + ", "
	}

	fmt.Fprintf(w, "func Call%s(client *dfhackrpc.Client%s) (%sdfhackrpc.CommandResult, error) {\n", method, inputDecl, outputDecl)

	requestVar := "request"
	if uw, ok := unwrap[[2]string{getPrefix2(input, packageName, pluginMessages), input}]; ok {
		requestVar = "&request"
		fmt.Fprint(w, "\tvar request ", getPrefix(input, pluginMessages), input, "\n")
		for _, u := range uw {
			fmt.Fprint(w, "\trequest.", u.Field, " = ", u.SetPrefix, strings.ToLower(u.Field), "\n")
		}
		fmt.Fprintln(w)
	}

	fmt.Fprintf(w, "\tvar response %s%s\n\n", getPrefix(output, pluginMessages), output)
	if uw, ok := unwrap[[2]string{getPrefix2(output, packageName, pluginMessages), output}]; ok && len(uw) == 0 {
		fmt.Fprint(w, "\treturn ")
	} else {
		fmt.Fprint(w, "\trv, err := ")
	}

	if plugin == "" {
		fmt.Fprintf(w, "client.Call(%q, %s, &response)\n", method, requestVar)
	} else {
		fmt.Fprintf(w, "client.CallPlugin(%q, %q, %s, &response)\n", plugin, method, requestVar)
	}

	if uw, ok := unwrap[[2]string{getPrefix2(output, packageName, pluginMessages), output}]; ok && len(uw) != 0 {
		fmt.Fprint(w, "\n\treturn ")
		for _, u := range uw {
			fmt.Fprint(w, "response.Get", u.Field, "(), ")
		}
		fmt.Fprintln(w, "rv, err")
	} else if !ok {
		fmt.Fprintln(w, "\n\treturn &response, rv, err")
	}

	fmt.Fprintln(w, "}")
}

func getPrefix(message string, pluginMessages []string) string {
	prefix := getPrefix2(message, "", pluginMessages)

	if prefix != "" {
		prefix += "."
	}

	return prefix
}

func getPrefix2(message, packageName string, pluginMessages []string) string {
	for _, m := range pluginMessages {
		if m == message {
			return packageName
		}
	}
	for _, m := range basicMessages {
		if m == message {
			return "dfproto"
		}
	}
	for _, m := range coreMessages {
		if m == message {
			return "dfproto_core"
		}
	}
	panic("unknown message: " + message)
}
