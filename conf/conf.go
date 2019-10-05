package conf

type app struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}
}

var App app

func Read() {
	App.Database.Host = "localhost"
	App.Database.User = "user"
	App.Database.Password = "password"
	App.Database.Port = "3306"
	App.Database.Dbname = "accountbook_dev"
}
