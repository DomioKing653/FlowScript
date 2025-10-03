
# FlowScript

FlowScript is a lightweight scripting language and interpreter written in Go. It features a Pratt parser, flexible syntax, and is designed for easy embedding and extension.

## Features

- Pratt parser for fast and flexible expression parsing
- Variable declarations (`let`, `const`)
- Arithmetic and logical expressions
- Statement-based syntax
- Extensible AST and parser

## Getting Started

### Prerequisites

- Go 1.20 or newer

### Build from Source

Clone the repository and build:

```bash
git clone https://github.com/DomioKing653/FlowScript.git
cd FlowScript
go build -o Flow.exe ./src/main.go
```

### Running FlowScript Code

Write your script in `examples/main.flw` or any `.flw` file.

Run:

```bash
./Flow.exe examples/main.flw
```

### Example

```js
let x = 5 + 2 * 3;
const y = x - 4;
```

## Contributing

Contributions are welcome! Feel free to open issues or pull requests. Please be respectful and constructive.

## License

MIT License. See LICENSE.md for details.
