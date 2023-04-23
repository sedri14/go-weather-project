## Weather Summary

This project is a weather program (written in Go) that scrapes weather data from three different weather sites and provides the average of the three as the result.

The program consists of several modules:

🌥️ a module for each scraped site that returns the data in a unified format

🌥️ a data module that calculates the data from all sources

🌥️ a client module that uses the data module to display the results



The program offers several functions that allow users to retrieve various types of weather information:

🌟 Will it rain? - receives a city and a number and returns the chance of rain in this city in the next x days.

🌟 Next rain day - returns the next day with a chance for rain over 50% in a specified city.

🌟 Average temp - gets a city and a number of days and returns the average temperature.

🌟 Temp array - gets a city and a number of days and returns the min and max temperature for every day from today to x.

🌟 Weather summary - returns the weather summary of today, including min and max temperature, humidity, wind, and chance of rain.
