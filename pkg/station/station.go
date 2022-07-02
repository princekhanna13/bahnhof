package station

type Station struct {
	Name string
	ID   string
}

var (
	Home = Station{
		Name: "Daglfing",
		ID:   "de:09162:700",
	}
	Work = Station{
		Name: "Ostabhnhof",
		ID:   "de:09162:5",
	}
)
