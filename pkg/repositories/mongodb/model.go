package mongodb

type DB struct {
	DBName        string
	Protocol      string
	IP            string
	Port          string
	User          string
	Password      string
	AuthSource    string
	AuthMechanism string
	Collections   string
}
