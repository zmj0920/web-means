package data

import (
	"time"
)

// 文章， 与表 threads 对应
type Thread struct {
	Id        int64     // `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	Uuid      string    // `uuid` varchar(64) NOT NULL,
	Topic     string    // `topic` text,
	UserId    int64     // `user_id` bigint(20) unsigned DEFAULT NULL,
	CreatedAt time.Time // `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
}

// 从数据库获取所有帖子，并返回
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id, uuid, topic, user_id, created_at from threads order by created_at desc") // 降序
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	_ = rows.Close()
	return
}


