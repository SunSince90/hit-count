# HitCount

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/SunSince90/hit-count/Test)
![GitHub top language](https://img.shields.io/github/languages/top/SunSince90/hit-count)
![GitHub](https://img.shields.io/github/license/SunSince90/hit-count)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/SunSince90/hit-count)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/SunSince90/hit-count)

A simple Golang package that counts occurrences of a certain event.

## Usage

```golang
import (
    counters "github.com/SunSince90/hit-count"
)

func main() {
    key404 = "404"
    key500 = "500"

    manager := counters.Manager("error-counters")

    // Create counters with target 10.
    // When that counter reaches 10, do something about this.
    manager.Add(key404, counters.NewIntCounter(0, 10))
    manager.Add(key500, counters.NewIntCounter(0, 10))

    for {
        // do something...
        resp, err := something()
        if err == Err404 {
            if manager.Get(key404).Increase().Hit() {
                // Got error 404 for 10 times **in a row**:
                // Do something about this
            }
        }
        if err == Err500 {
            // Need error 404 to happen for ten times in a row.
            // This time it didn't happen, so we reset the couter.
            manager.Get(key404).Reset()

            if manager.Get(key500).Increase().Hit() {
                // Got error 500 for 10 times:
                // Do something about this.
                // NOTE: We don't need Err500 to happen 10 times in a row,
                // that's why we don't reset it, like we did with key404
            }
        }
    }

    // Reset all counters...
    manager.Reset()
}
```
