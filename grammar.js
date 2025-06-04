/**
 * @file Syntax Highlighting for the Aalgola Programming Language
 * @author Aalgolab GmbH <info@aalgolab.de>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "aalgola",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
