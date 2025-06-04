import XCTest
import SwiftTreeSitter
import TreeSitterAalgola

final class TreeSitterAalgolaTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_aalgola())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Aalgola grammar")
    }
}
