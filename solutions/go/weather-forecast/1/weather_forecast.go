//Package weather forecasts the weather.
package weather

var (
    //CurrentCondition is the current weather condition.
	CurrentCondition string
    //CurrentLocation is the current location being forecast.
	CurrentLocation  string
)

//Forecast function prints the forecast for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
