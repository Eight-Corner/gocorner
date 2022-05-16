package constants

const (
	host     = "localhost"
	port     = 8000
	database = "cornerdb"
	user     = "root"
	password = "1234"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
