# gtranslate 

Google Translate API for unlimited and free translations 📢.

# Install

    go get github.com/ystyle/gtranslate

# Use

```go
// set google domain (if you need)
gtranslate.SetGoogleDomain("https://translate.google.cn")

// translate
gtranslate.Translate("I'm alive", language.English, language.Spanish)
gtranslate.TranslateWithParams("I'm alive", gtranslate.TranslateWithParams{From: "en", To: "es"})

```




# Example

```go
package main

import (
	"fmt"
	"github.com/ystyle/gtranslate"
	"time"
)

func main() {
	text := "Hello World"
	start := time.Now()
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   "zh",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | zh: %s | time: %s\n", text, translated, time.Now().Sub(start))
	// en: Hello World | zh: 你好，世界 | time: 547.2785ms
}
```
