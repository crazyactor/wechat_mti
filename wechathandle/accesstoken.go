package wechathandle

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type AccessToken struct {
	AppId      string
	AppSercret string
	TmpName    string
	LckName    string
}

//移除锁定文件
func (this *AccessToken) unlock() error {
	return os.Remove(this.LckName)
}

//建立锁定文件
func (this *AccessToken) lock() error {
	path := path.Dir(this.LckName)
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	if !fi.IsDir() {
		return errors.New("path is not a directory")
	}
	lck, err := os.Create(this.LckName)
	if err != nil {
		return err
	}
	lck.Close()
	return nil
}

//锁定文件存在与否，存在为锁定状态
func (this *AccessToken) locked() bool {
	_, err := os.Stat(this.LckName)
	return !os.IsNotExist(err)
}

//直接调用API 获取access_token
func (this *AccessToken) fetch() (string, error) {
	rtn, err := get(fmt.Sprintf("%stoken?grant_type=client_credential&appid=%s&secret=%s",
		UrlPrefix,
		this.AppId,
		this.AppSercret,
	))
	if err != nil {
		return "", err
	}
	return rtn.AccessToken, nil
}

//存储指定的access_token
func (this *AccessToken) store(token string) error {
	path := path.Dir(this.TmpName)
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	if !fi.IsDir() {
		return errors.New("path is not a directory")
	}
	tmp, err := os.OpenFile(this.TmpName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer tmp.Close()
	if _, err := tmp.Write([]byte(token)); err != nil {
		return err
	}
	return nil
}

//获取最新的access_token
func (this *AccessToken) Fresh() (string, error) {
	if this.TmpName == "" {
		this.TmpName = this.AppId + "-accesstoken.tmp"
	}
	if this.LckName == "" {
		this.LckName = this.TmpName + ".lck"
	}
	for {
		if this.locked() {
			time.Sleep(time.Second)
			continue
		}
		break
	}
	fi, err := os.Stat(this.TmpName)
	if err != nil && !os.IsExist(err) {
		return this.fetchAndStore()
	}
	expires := fi.ModTime().Add(2 * time.Hour).Unix()
	if expires <= time.Now().Unix() {
		return this.fetchAndStore()
	}
	tmp, err := os.Open(this.TmpName)
	if err != nil {
		return "", err
	}
	defer tmp.Close()
	data, err := ioutil.ReadAll(tmp)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//取得access_token并存储
func (this *AccessToken) fetchAndStore() (string, error) {
	if err := this.lock(); err != nil {
		return "", err
	}
	defer this.unlock()
	token, err := this.fetch()
	if err != nil {
		return "", err
	}
	if err := this.store(token); err != nil {
		return "", err
	}
	return token, nil
}
