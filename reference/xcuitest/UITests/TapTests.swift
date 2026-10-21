import XCTest

// Opens TAP_URL in Safari and taps, one after another, the elements named in
// TAP_STEPS (their labels), each XCUITest's tap: a touch of the simulator's
// screen through its HID, as a finger's, not WebDriver's. After each it waits
// for the page's "reported N" (what the tap left, sent to the probe's run.mjs:
// ../xcuitest.mjs).
// Safari on a new simulator puts a tip over the page ("View Bookmarks, Share
// Menu, and Open Tabs"), a popover whose dismiss region covers the window and
// takes the next touch: the page gets its touch events and no click. So the
// tip is closed first, and no popover of Safari's may be up when a tap lands.
final class TapTests: XCTestCase {
    func testTaps() throws {
        continueAfterFailure = false
        let env = ProcessInfo.processInfo.environment
        let url = try XCTUnwrap(URL(string: try XCTUnwrap(env["TAP_URL"])))
        let steps = try XCTUnwrap(env["TAP_STEPS"]).split(separator: ",").map(String.init)
        let safari = XCUIApplication(bundleIdentifier: "com.apple.mobilesafari")
        safari.launch()
        safari.open(url)
        let web = safari.webViews.firstMatch
        XCTAssertTrue(web.staticTexts["ready"].waitForExistence(timeout: 60), "the page did not show")
        let dismiss = safari.otherElements["PopoverDismissRegion"]
        if dismiss.waitForExistence(timeout: 5) {
            safari.popovers.firstMatch.buttons["xmark.circle.fill"].tap()
            XCTAssertTrue(dismiss.waitForNonExistence(timeout: 5), "Safari's tip did not close")
            print("ciderui-tap: Safari's tip closed")
        }
        for (i, label) in steps.enumerated() {
            let target = web.descendants(matching: .any)[label].firstMatch
            XCTAssertTrue(target.waitForExistence(timeout: 10), "no \(label) on the page")
            XCTAssertFalse(dismiss.exists, "a popover of Safari's is over the page")
            target.tap()
            XCTAssertTrue(web.staticTexts["reported \(i + 1)"].waitForExistence(timeout: 10), "no report after \(label)")
        }
    }
}
