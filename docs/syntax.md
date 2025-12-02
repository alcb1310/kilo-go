# Add Syntax Highlighting for a language

Inside the `$XDG_CONFIG_HOME/kilo` directory, we will need to create a `highlight` directory and
create a new file for each of the languages you want support for. If no files exist in the specified
directory, then `Kilo` will not attempt to highlight any of the code in the editor.

The following fields will be required:

- `number` is a boolean field that will enable number highlighting.
- `string` is a boolean field that will enable string highlighting.
- `filetype` is a string field that will be used to identify the language.
- `extensions` is a list of strings that will be used to identify the file extensions.
- `keywords` is a list of strings that will be used to identify the keywords.
- `types` is a list of strings that will be used to identify the types.
- `slc` is a string that will be used to identify the single line comment.
- `mlcs` is a string that will be used to identify the multi line comment start.
- `mlce` is a string that will be used to identify the multi line comment end.

## Example

Following is an expample for the `Go` language:

```toml
number=true
string=true
filetype="GO"
extensions=[".go"]
keywords=[ "package", "import", "func", "type", "var", "const", "if", "else",
	"switch", "case", "default", "for", "range", "goto", "continue", "select", "return", "break",
]
types=[ "bool", "byte", "error", "float32", "float64", "int", "int16", "int32", "nil",
	"int64", "int8", "rune", "string", "uint", "uint16", "uint32", "uint64", "uint8",
]
slc="//"
mlcs="/*"
mlce="*/"
```
