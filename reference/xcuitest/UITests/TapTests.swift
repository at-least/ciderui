import XCTest

// Opens TAP_URL in Safari and takes, one after another, the steps in
// TAP_STEPS: an element's label, tapped, XCUITest's tap: a touch of the
// simulator's screen through its HID, as a finger's, not WebDriver's;
// press:<label>, the element held a second, a long press; type:<text>, the
// text typed by XCUITest's keyboard into what has the focus. After each it
// waits for the page's "reported N" (what the step left, sent to the probe's
// run.mjs: ../xcuitest.mjs).
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
        for (i, step) in steps.enumerated() {
            XCTAssertFalse(dismiss.exists, "a popover of Safari's is over the page")
            if step.hasPrefix("type:") {
                safari.typeText(String(step.dropFirst(5)))
            } else {
                let press = step.hasPrefix("press:")
                let label = press ? String(step.dropFirst(6)) : step
                let target = web.descendants(matching: .any)[label].firstMatch
                XCTAssertTrue(target.waitForExistence(timeout: 10), "no \(label) on the page")
                if press { target.press(forDuration: 1) } else { target.tap() }
            }
            XCTAssertTrue(web.staticTexts["reported \(i + 1)"].waitForExistence(timeout: 10), "no report after \(step)")
        }
    }
}
