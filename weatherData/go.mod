module example.com/weatherData

go 1.19

replace example.com/common => ../common

replace example.com/source1 => ../source1

replace example.com/source2 => ../source2

replace example.com/source3 => ../source3

require (
	example.com/common v0.0.0-00010101000000-000000000000
	example.com/source1 v0.0.0-00010101000000-000000000000
)

require (
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
)
