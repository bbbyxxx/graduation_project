package dal

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func DelKey(conn redis.Conn, key string) error {
	var (
		err error
	)
	if conn == nil {
		conn, err = GetConn()
		if err != nil {
			return err
		}
	}
	defer conn.Close()
	_, err = conn.Do("Del", key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SetEx(conn redis.Conn, key string, value interface{}, expireTime int64) error {
	var (
		err error
	)
	if conn == nil {
		conn, err = GetConn()
		if err != nil {
			return err
		}
	}
	defer conn.Close()
	_, err = conn.Do("SetEx", key, expireTime, value)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Get(conn redis.Conn, key string) (interface{}, error) {
	var (
		err       error
		resInt    int
		resString string
	)
	if conn == nil {
		conn, err = GetConn()
		if err != nil {
			return nil, err
		}
	}
	resInt, err = redis.Int(conn.Do("Get", key))
	if err != nil {
		fmt.Println("get Int failed,", err)
		resString, err = redis.String(conn.Do("Get", key))
		if err != nil {
			fmt.Println("get String failed,", err)
			return nil, err
		}
	}

	fmt.Println(resInt)
	fmt.Println(resString)

	if resInt == 0 {
		return resString, nil
	} else {
		return resInt, nil
	}
}
