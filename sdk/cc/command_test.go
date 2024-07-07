package cc

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

type Command struct {
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Syntax      []string `json:"Syntax"`
	Examples    []string `json:"Examples"`
	Notes       string   `json:"Notes"`
}

type Commands struct {
	Debugging []Command
	Extension []Command
	Hwdbg     []Command
	Meta      []Command
}

func TestUnmarshalCommandJson(t *testing.T) {
	type AutoGenerated struct {
		Name        string   `json:"Name"`
		Description string   `json:"Description"`
		Syntax      []string `json:"Syntax"`
		Examples    []string `json:"Examples"`
		Notes       []any    `json:"Notes"`
		FullName    string   `json:"FullName"`
	}
	var generated []AutoGenerated
	mylog.Check(json.Unmarshal(stream.NewBuffer("sina2.json").Bytes(), &generated))
	// mylog.Struct(generated)
	g := stream.NewGeneratedFile()
	g.P("package sdk")
	g.P()
	for _, s := range generated {
		fullName := s.FullName

		fullName = strings.ReplaceAll(fullName, "DetectRdmsrExecution", "ReadMsr")
		fullName = strings.ReplaceAll(fullName, "DetectWrmsrExecution", "WriteMsr")

		fullName = strings.ReplaceAll(fullName, "-", "_")
		if fullName != "ContinueDebuggee" {
			fullName = stream.ToCamelUpper(fullName, false)
		}
		fullName = strings.ReplaceAll(fullName, "ContinueDebuggee", "ContinueDebuggee_")
		g.P("func ", fullName, "() {")
		g.P("InterpreterEx(", strconv.Quote(s.Name), ") ")
		g.P("}")
		g.P()
	}
	stream.WriteGoFile("../commandWrapper.go", g.Bytes())
}
