package work2

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

/*
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
link:https://juejin.cn/post/6844904177714331661

注意：无论是 Wrap ， WithMessage 还是 WithStack ，当传入的 err 参数为 nil 时， 都会返回nil， 这意味着我们在调用此方法之前无需作 nil 判断，保持了代码简洁
*/

func GetSql() error {
	return errors.Wrap(sql.ErrNoRows, "GetSql failed")
}

/*

 */
func GetSql2() error {
	return errors.WithStack(sql.ErrNoRows)
}

func Call() error {
	return errors.WithMessage(GetSql(), "bar failed")
}

func main() {
	err := Call()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}
