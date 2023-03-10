# 🍒 Cranberry

Cranberry is a small Go program that helps you find console logging statements in your TypeScript project.

This has been tested in the following frameworks:

- Angular 
- React
- Svelte
- Vue
- Solid.js

## 🌟 Features

Cranberry can:

- Find `console.log()`, `console.table()`, `console.warn()`, `console.info()`, and `console.debug()` statements in your TypeScript code.
- Display the file path, line number, and line contents of each console statement it finds.
- Count the number of console statements found and display a summary at the end.
- [EXPERIMENTAL] Remove the offending lines in your project directory.

## 🚀 Usage

```bash
cranberry [options]
```

### 📜 Options

- `-d`, `--dir <path>`: Directory to examine.
- `-h`, `--help`: Display help message.
- `-r`, `--remove`: Remove offending lines.

## How it works

Cranberry looks for `.ts` (TypeScript) files recursively in the directory you specify (or the current directory by default) and checks each file for the following console methods:

- `console.log()`
- `console.table()`
- `console.warn()`
- `console.info()`
- `console.debug()`

If it finds any of these methods, but not `console.error()`, it will print the offending line(s) with the file name and line number.

## Installation

To use Cranberry, you can either download the binary from the [releases page](https://github.com/VentGrey/Cranberry/releases) or use it with Go installed on your machine.

1. Clone the repositoy: `git clone https://github.com/VentGrey/Cranberry`
2. Build the binary: `cd Cranberry && go build -o cranberry`
3. Run the binary: `./cranberry`

## Example

```bash
$ cranberry -d src/
console.log() in src/index.ts, line 10: console.log('Hello, world!');
console.table() in src/index.ts, line 20: console.table([{name: 'Alice', age: 30}, {name: 'Bob', age: 40}]);
We found 2 incidents!
```

## Credits

Cranberry was created by [VentGrey & La Esquina Gris](https://ventgrey.github.io/)

## 📝 License

Cranberry is licensed under the GPL-3+ License. See LICENSE for more information.
