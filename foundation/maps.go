package main
import "fmt"

func main(){
	freezing := make(map[string]float32)
    freezing["celsius"] = 0.0
	freezing["farhenhite"] = 32.0
	freezing["kelvin"] = 273.2

	fmt.Println(freezing["celsius"])
    fmt.Println(len(freezing))

	frozen := map[string]float32{
		"celsius": 0.0,
		"farhenhite": 32.0,
		"kelvin": 273.2,
	}
	fmt.Println(frozen["kelvin"])

}