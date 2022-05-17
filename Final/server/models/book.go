package models

type Book struct {
	tableName struct{} `pg:"books"`
	BookID    int      `pg:"book_id,pk"`
	Author    string   `pg:"author"`
	Title     string   `pg:"title"`
	Rating    int      `pg:"rating"`
}

type User struct {
	tableName struct{} `pg:"users"`
	UserID    int      `pg:"user_id,pk"`
	Login     string   `pg:"login"`
	Password  string   `pg:"password"`
}

type Admin struct {
	tableName struct{} `pg:"admins"`
	AdminID   int      `pg:"admin_id,pk"`
	Login     string   `pg:"login"`
	Password  string   `pg:"password"`
}

type UserBook struct {
	tableName struct{} `pg:"users_books"`
	ID        int      `pg:"id,pk"`
	UserID    int      `pg:"user_id"`
	BookID    int      `pg:"book_id"`
	Title     string   `pg:"title"`
	Status    string   `pg:"status"`
	Star      bool     `pg:"star"`
}

type BookComplaint struct {
	tableName struct{} `pg:"books_complaint"`
	ID        int      `pg:"id,pk"`
	BookID    int      `pg:"book_id"`
	Title     string   `pg:"title"`
	Complaint string   `pg:"complaint"`
}
