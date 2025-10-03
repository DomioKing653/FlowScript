# FlowScript Syntax Highlighting (example)

This folder contains a minimal VS Code extension scaffold that provides TextMate-based syntax highlighting for the FlowScript language (files with the `.flow` extension).

How to test locally

1. Open this folder (`Extension`) in VS Code.
2. Press F5 to launch an Extension Development Host.
3. In the new window, open `examples/example.flow` and you should see highlighting.

Notes

- This uses a simple TextMate grammar (`syntaxes/flowscript.tmLanguage.json`). Tweak patterns to match your language.
- When you're ready to publish, add an icon and package metadata, then use `vsce` to package and publish.
