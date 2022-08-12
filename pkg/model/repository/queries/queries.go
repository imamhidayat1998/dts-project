package queries

const (
	CreateUser = `insert into user (id,username,password) values (?,?,?)`
)
