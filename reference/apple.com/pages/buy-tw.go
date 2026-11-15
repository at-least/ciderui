package pages

import "github.com/at-least/ciderui/internal/pagespec"

func buyTw() pagespec.Spec {
	s := buy()
	s.URL = "https://www.apple.com/tw/shop/buy-iphone/iphone-18-pro"
	return s
}
