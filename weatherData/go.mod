module example.com/weatherData

go 1.19

replace example.com/source1 => ../source1

replace example.com/source2 => ../source2

replace example.com/source3 => ../source3

require (
	example.com/source1 v0.0.0-00010101000000-000000000000
	example.com/source2 v0.0.0-00010101000000-000000000000
	example.com/source3 v0.0.0-00010101000000-000000000000
)
