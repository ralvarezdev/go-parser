# go-parser

Go source parsing helpers: parse a Go file into an AST, traverse it, write it back, and rewrite struct `json:"..."` tags programmatically. Requires Go 1.23.4 (per `go.mod`).

## Installation

```bash
go get github.com/ralvarezdev/go-parser
```

## API

**Package `goparser`**

- **`Parser`** — `ParseFile(filePath)` (parses with comments), `WriteFile(filePath, fileSet, node)`, `TraverseAST(node, fn)`.
- **`NewDefaultParser()`** — `go/parser` + `go/printer` implementation.
- **Errors** — `ErrNilParser`, `ErrNilFileSet`, `ErrNilASTNode`.

**Package `tags/json`**

- **`Writer`** — `OverwriteTags(filePath, structsTagsMapper)`, `HideStructsTags(filePath, structsFieldsMapper)`, `HideFilesTags(filesStructsFieldsMapper)`.
- **`NewDefaultWriter(parser, logger)`** — default implementation; the `*slog.Logger` is optional.
- **Mapper types** — `FieldsTagsMapper` (field to tag), `StructsTagsMapper` (struct to fields), `StructsFieldsMapper`, `FilesStructsFieldsMapper`.

## Usage

```go
parser := goparser.NewDefaultParser()
writer, err := gotagsjson.NewDefaultWriter(parser, nil)

err = writer.OverwriteTags("model.go", gotagsjson.StructsTagsMapper{
    "User": gotagsjson.FieldsTagsMapper{
        "Email": `json:"email,omitempty"`,
    },
})
```

The expected tag string format is defined in `tags/json/writer.go`; existing tags are matched with the pattern `json:"[^"]*"`.

## Development

```bash
go build ./...
```

There are no tests.

## License

GNU General Public License v3.0 (see [LICENSE](LICENSE)).
