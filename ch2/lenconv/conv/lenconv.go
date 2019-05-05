package lenconv

import "fmt"

type Feet float64
type Meters float64

type Pounds float64
type Kilograms float64

func (f Feet) String() string      { return fmt.Sprintf("%fft", f) }
func (m Meters) String() string    { return fmt.Sprintf("%fm", m) }
func (p Pounds) String() string    { return fmt.Sprintf("%flb", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%fkg", k) }

func FeetToMeters(feet Feet) Meters { return Meters(feet / 3.2808) }

func PoundsToKilograms(pounds Pounds) Kilograms { return Kilograms(pounds / 2.2046) }
