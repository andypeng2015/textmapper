package gen

import (
	"embed"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/inspirer/textmapper/grammar"
)

var languages = map[string]*language{
	"go": {
		SharedDefs: builtin(`go_shared`),
		CachedDefs: builtin(`go_cached`),
		Lexer: []file{
			{"token/token.go", builtin(`go_token`)},
			{"lexer_tables.go", builtin(`go_lexer_tables`)},
			{"lexer.go", builtin(`go_lexer`)},
		},
		Parser: []file{
			{"parser.go", builtin(`go_parser`)},
			{"parser_tables.go", builtin(`go_parser_tables`)},
		},
		Stream: []file{
			{"stream.go", builtin(`go_stream`)},
		},
		Types: []file{
			{"listener.go", builtin(`go_listener`)},
		},
		Selector: []file{
			{"selector/selector.go", builtin(`go_selector`)},
		},
		AST: []file{
			{"ast/tree.go", builtin(`go_ast_tree`)},
			{"ast/parse.go", builtin(`go_ast_parse`)},
		},
		TypedAST: []file{
			{"ast/ast.go", builtin(`go_ast`)},
			{"ast/factory.go", builtin(`go_ast_factory`)},
		},
	},

	"cc": {
		SharedDefs: builtin(`cc_shared`),
		CachedDefs: builtin(`cc_cached`),
		Lexer: []file{
			{"token.h", builtin(`cc_token_h`)},
			{"lexer.h", builtin(`cc_lexer_h`)},
			{"lexer.cc", builtin(`cc_lexer_cc`)},
		},
		Parser: []file{
			{"parser.h", builtin(`cc_parser_h`)},
			{"parser.cc", builtin(`cc_parser_cc`)},
		},
	},

	"ts": {
		SharedDefs: builtin(`ts_shared`),
		CachedDefs: builtin(`ts_cached`),
		Lexer: []file{
			{"token.ts", builtin(`ts_token`)},
			{"lexer_tables.ts", builtin(`ts_lexer_tables`)},
			{"lexer.ts", builtin(`ts_lexer`)},
		},
		Types: []file{
			{"common.ts", builtin("ts_common")},
			{"listener.ts", builtin("ts_listener")},
		},
		Selector: []file{
			{"selector.ts", builtin(`ts_selector`)},
		},
		Parser: []file{
			{"parser.ts", builtin("ts_parser")},
			{"parser_tables.ts", builtin("ts_parser_tables")},
		},
		Stream: []file{
			{"stream.ts", builtin(`ts_stream`)},
		},
		AST: []file{
			{"tree.ts", builtin("ts_tree")},
			{"builder.ts", builtin("ts_builder")},
		},
	},
}

type file struct {
	name     string
	template string
}

type language struct {
	Lexer    []file
	Parser   []file
	Stream   []file
	Types    []file
	Selector []file
	AST      []file
	TypedAST []file
	Bison    []file

	SharedDefs string
	CachedDefs string
}

func (l *language) templates(g *grammar.Grammar) []file {
	var ret []file
	if g.Lexer.Tables != nil {
		ret = append(ret, l.Lexer...)
	} else {
		// Take the token file only.
		ret = append(ret, l.Lexer[0])
		if g.Options.FlexMode {
			ret = append(ret, file{name: "token_codes.inc", template: tokenCodesTpl})
		}
	}
	if g.Parser.Tables != nil {
		ret = append(ret, l.Parser...)
		if g.Options.TokenStream {
			ret = append(ret, l.Stream...)
		}
	}
	if g.Parser.Types != nil {
		ret = append(ret, l.Types...)
		if g.Options.GenSelector || g.Options.EventFields {
			ret = append(ret, l.Selector...)
		}
		if g.Options.EventAST {
			ret = append(ret, l.AST...)
		}
		if g.Options.EventFields && g.Options.EventAST {
			ret = append(ret, l.TypedAST...)
		}
	}
	if g.Options.WriteBison {
		ret = append(ret, file{name: g.Name + ".y", template: bisonTpl})
	}
	return ret
}

var bisonTpl = builtin(`bison`)
var tokenCodesTpl = builtin(`cc_token_codes_inc`)

//go:embed templates/*
var fs embed.FS

func builtin(name string) string {
	name = "templates/" + name + ".go.tmpl"
	content, err := fs.ReadFile(name)
	if err != nil {
		panic(fmt.Sprintf("cannot read %v: %v", name, err))
	}
	return patchTemplates(string(content))
}

// patchTemplates changes the syntax of Go templates and makes ` -}}` consume
// a single newline character only (instead of all the following whitespace).
func patchTemplates(tmpl string) string {
	const seq = " -}}\n"

	strs := strings.SplitAfter(tmpl, seq)
	if len(strs) == 1 {
		// Fast path. Nothing to patch.
		return tmpl
	}

	var ret strings.Builder
	ret.WriteString(strs[0])

	for _, s := range strs[1:] {
		if r, _ := utf8.DecodeRuneInString(s); unicode.IsSpace(r) {
			// Insert a stopper to keep the line numbers intact for error messages.
			ret.WriteString("{{/**/}}")
		}
		ret.WriteString(s)
	}
	return ret.String()
}
