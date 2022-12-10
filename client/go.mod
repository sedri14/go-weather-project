module example.com/client

go 1.19

replace example.com/weatherData => ../weatherData

require example.com/weatherData v0.0.0-00010101000000-000000000000

require (
	example.com/common v0.0.0-00010101000000-000000000000 // indirect
	example.com/source1 v0.0.0-00010101000000-000000000000 // indirect
	example.com/source2 v0.0.0-00010101000000-000000000000 // indirect
	example.com/source3 v0.0.0-00010101000000-000000000000 // indirect
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
)

replace example.com/source1 => ../source1

replace example.com/source2 => ../source2

replace example.com/source3 => ../source3

replace example.com/common => ../common
