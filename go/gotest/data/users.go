package data

type Users struct {
	Id int
	Name string 
	Email string
	Password string
	CreatedAt time.Time
}

func Userquery( pageNum int, pageSize int) (users[] Users, int,int,int)  {
	    count, err :=Db.Query("select count(*) from users") 
		if err != nil {
			return
		}
		// total = count
		rows, err := Db.Query("select id, name, email, password, created_at from users order by created_at desc limit ?,?",(pageNum-1)*pageSize,pageSize) // 降序
        if err != nil {
			return
		}
		for  row.Next(){
			user := Users{}
			if err = row.Scan(&user.Id, &user.Name,&user.Email,&user.Password,&user.CreatedAt)
          users = append(users,user)
		}
		_ = rows.Close()
		return users,pageNum,pageSize,count
}