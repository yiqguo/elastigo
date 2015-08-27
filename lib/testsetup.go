package elastigo

import (
	"testing"
	"time"
)

/*
// elastigo Conn adapter to avoid a circular dependency
type conn interface {
	CreateIndex(name string) (interface{}, error)
	DeleteIndex(name string) (interface{}, error)

	Index(index string, _type string, id string, args map[string]interface{}, data interface{}) (interface{}, error)
}
*/

func newIndexWorker(c *Conn, t *testing.T) func(interface{}) {

	return func(d interface{}) {
		_, err := c.Index("oilers", "heyday", "", nil, d)
		if err != nil {
			t.Fatalf("Index failed: %s", err)
		}
	}
}

func PopulateTestDB(t *testing.T, c *Conn) {

	// it is not technically necessary to create an index here
	c.CreateIndex("oilers")

	idx := newIndexWorker(c, t)

	idx(`{"name": "Mark Messier",   "jersey": 11, "pos": "LW", "goals": 37, "PIM": 165, 
			"born": "19610118", "teams": ["EDM", "NYR", "VAN"]}`)
	idx(`{"name": "Wayne Gretzky",  "jersey": 99, "pos": "C",  "goals": 87,
			"born": "19610126", "teams": ["EDM", "NYR", "STL"]}`)
	idx(`{"name": "Paul Coffey",    "jersey": 7,  "pos": "D",  "goals": 40,
			"born": "19610601", "teams": ["EDM", "DET"]}`)
	idx(`{"name": "Jari Kurri",     "jersey": 17, "pos": "RW", "goals": 52,
			"born": "19600518", "teams": ["EDM", "VAN"]}`)
	idx(`{"name": "Glenn Anderson", "jersey": 9,  "pos": "RW", "goals": 54,
			"born": "19601002", "teams": ["EDM", "NYR", "TOR", "STL"]}`)
	idx(`{"name": "Ken Linseman",   "jersey": 13, "pos": "C",  "goals": 18,
			"born": "19580811", "teams": ["EDM", "TOR"]}`)
	idx(`{"name": "Pat Hughes",     "jersey": 16, "pos": "RW", "goals": 27,
			"born": "19550325", "teams": ["EDM", "MTL", "PIT"]}`)
	idx(`{"name": "Dave Hunter",    "jersey": 12, "pos": "LW", "goals": 22,
			"born": "19580101", "teams": ["EDM", "PIT"]}`)
	idx(`{"name": "Kevin Lowe",     "jersey": 4,  "pos": "D",  "goals": 4,
			"born": "19590415", "teams": ["EDM", "NYR"]}`)
	idx(`{"name": "Charlie Huddy",  "jersey": 22, "pos": "D",  "goals": 8,
			"born": "19590602", "teams": ["EDM", "BUF", "STL"]}`)
	idx(`{"name": "Randy Gregg",    "jersey": 21, "pos": "D",  "goals": 13,
			"born": "19560219", "teams": ["EDM", "VAN"]}`)
	idx(`{"name": "Dave Semenko",   "jersey": 27, "pos": "LW", "goals": 4, "PIM": 118,
			"born": "19570712", "teams": ["EDM"]}`)
	idx(`{"name": "Grant Fuhr",     "jersey": 31, "pos": "G",  "GAA": 3.91,
			"born": "19620928", "teams": ["EDM", "TOR", "BUF", "STL"]}`)
	idx(`{"name": "Andy Moog",      "jersey": 35, "pos": "G",  "GAA": 3.77,
			"born": "19600218", "teams": ["EDM", "BOS", "DAL", "MTL"]}`)

	// HACK to let the ES magic happen
	time.Sleep(time.Second)
}

func TearDownTestDB(c *Conn) {
	c.DeleteIndex("oilers")
}
