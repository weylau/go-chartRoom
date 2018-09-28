package model

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"fmt"
)


var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func(this *UserDao) InitUserDao(pool *redis.Pool) (userDao *UserDao) {
	return &UserDao{
		pool:pool,
	}
}

func (this *UserDao) getUserByName(conn redis.Conn, user_name string) (user *User, err error) {
	//查询用户是否存在
	ret, err := redis.String(conn.Do("HGet", "users", user_name))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		fmt.Println("redis error:", err)
		return user, err
	}
	fmt.Println("redis ret:", ret)

	user = &User{}
	err = json.Unmarshal([]byte(ret), user)
	if err != nil {
		return user, err
	}
	return
}

func (this *UserDao) createUser(conn redis.Conn, user *User) (err error) {
	userStr, err := json.Marshal(user)
	if err != nil {
		return err
	}
	//查询用户是否存在
	_, err = redis.String(conn.Do("HGet", "users", user.UserName))
	if err == nil {
		err = ERROR_USER_EXISTS
		return err
	}
	_, err = conn.Do("HSet", "users", user.UserName, string(userStr))
	if err != nil {
		fmt.Println("redis error:", err)
		return err
	}
	return
}


func (this *UserDao) Login(user_name string, user_pwd string) (user *User, err error) {
	conn := this.pool.Get()
	user, err = this.getUserByName(conn, user_name)
	if err != nil {
		fmt.Println("user.getUserById error:", err)
		return user, err
	}
	if user.UserPwd != user_pwd {
		err = ERROR_USER_PWD
		return
	}
	return
}


func (this *UserDao) Register(user_id int, user_name string, user_pwd string) (user *User, err error) {
	conn := this.pool.Get()
	user = &User{
		UserId:user_id,
		UserName:user_name,
		UserPwd:user_pwd,
	}
	err = this.createUser(conn, user)
	if err != nil {
		fmt.Println("user.createUser error:", err)
		return user, err
	}
	return
}
