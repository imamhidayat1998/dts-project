package queries

const (
	CreateUser         = `insert into user (id,fullname,password) values (?,?,?)`
	QueryUpdateUser    = `update set fullname=?,password=? from user where id=? `
	SelectIdUser       = `select id from user where fullname = ? and password = ?`
	QueryInsertArtikel = `insert into artikel(artikel_id,judul,author,description,created_at) values(?,?,?,?,?)`
)
