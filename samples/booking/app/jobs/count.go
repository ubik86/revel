package jobs

import (
	"fmt"
	"github.com/ubik86/revel"
	"github.com/ubik86/revel/modules/jobs/app/jobs"
	"github.com/ubik86/revel/samples/booking/app/controllers"
	"github.com/ubik86/revel/samples/booking/app/models"
)

// Periodically count the bookings in the database.
type BookingCounter struct{}

func (c BookingCounter) Run() {
	bookings, err := controllers.Dbm.Select(models.Booking{},
			`select * from Booking`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d bookings.\n", len(bookings))
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("@every 1m", BookingCounter{})
	})
}
