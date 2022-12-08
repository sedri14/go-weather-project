module example.com/client

go 1.19

replace example.com/weatherData => ../weatherData

require example.com/weatherData v0.0.0-00010101000000-000000000000

require (
	example.com/source1 v0.0.0-00010101000000-000000000000 // indirect
	example.com/source2 v0.0.0-00010101000000-000000000000 // indirect
	example.com/source3 v0.0.0-00010101000000-000000000000 // indirect
)

replace example.com/source1 => ../source1

replace example.com/source2 => ../source2

replace example.com/source3 => ../source3
