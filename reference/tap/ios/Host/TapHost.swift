import SwiftUI

// The UI test's host: XCUITest needs an app of its own; the test drives Safari.
@main
struct TapHost: App {
    var body: some Scene {
        WindowGroup { Text("ciderui tap") }
    }
}
