module example.com/source1

replace example.com/common => ../common

go 1.19

require (
	example.com/common v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.8.0
)

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
)
