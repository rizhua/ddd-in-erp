package converter

import "rizhua.com/infrastructure/persistence/po"

type User struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Birthday string `json:"birthday"`
	Gender   int8   `json:"gender"`
	Status   int8   `json:"status"`
	LastTime string `json:"lastTime"`
	CreateAt string `json:"createAt"`
}

func (t *User) Info(in po.User) (info User) {
	info = User{
		ID:       in.ID,
		Nickname: in.Nickname,
		Email:    in.Email,
		Mobile:   in.Mobile,
		Birthday: in.Birthday.Format("2006-01-02"),
		Gender:   in.Gender,
		Status:   in.Status,
		LastTime: in.LastTime.Format("2006-01-02 15:04:05"),
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *User) List(in []po.User) (list []User) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
