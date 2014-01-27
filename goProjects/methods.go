package main
import "fmt"

//declare time zone type, declare constants
type TZ int
const (
	HOUR TZ = 60*60; UTC TZ = 0*HOUR; EST TZ = -5*HOUR;
)

//declare timeZones map of strings to hold timezone strings
var timeZones = map[string]TZ { "UTC": UTC, "EST": EST }

func (tz TZ) String() string {	//Method on TZ (not ptr)
	for name, zone := range timeZones {
		if tz == zone { return name }
	}
	return fmt.Sprintf("%+d:%02d", tz/3600, (tz%3600)/60);
}

func main() {
	fmt.Println(EST);	// Print* knows about method String()
	fmt.Println(5*HOUR/2);
}
// Output (two lines) EST +2:30
