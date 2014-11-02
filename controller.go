package goweb

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Controller struct {
	Data map[string]interface{}
}

func (ct *Controller) Init(w http.ResponseWriter, r *http.Request) {
	ct.Data = make(map[string]interface{})
	ct.Data["test"] = 123
	Logger.Debug("Init")
}

func (ct Controller) Post(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No such page", http.StatusNotFound)
}

func (ct Controller) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func (ct Controller) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No such page", http.StatusNotFound)
}

func (ct Controller) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No such page", http.StatusNotFound)
}

func (ct *Controller) Err400(w http.ResponseWriter, r *http.Request, title string, info string) {
	ct.Execute(w, "view/layout.tpl", "view/400.tpl")
}

func (ct *Controller) GetTime() (t int64) {
	t = time.Now().Unix()
	return
}

func (ct *Controller) SetSession(w http.ResponseWriter, r *http.Request, key string, value string) {
	session := SessionManager.StartSession(w, r)
	session.Set(key, value)
}

func (ct *Controller) GetSession(w http.ResponseWriter, r *http.Request, key string) (value string) {
	session := SessionManager.StartSession(w, r)
	value = session.Get(key)
	return
}

func (ct *Controller) DeleteSession(w http.ResponseWriter, r *http.Request) {
	SessionManager.DeleteSession(w, r)
}

func (c *Controller) Execute(w io.Writer, tplfiles ...string) {
	t, err := ParseFiles(tplfiles...)
	if err == nil {
		err = t.Execute(w, c.Data)
	}
	if err != nil {
		//模板产生的错误应该属于debug错误，所以不对用户显示
		Logger.Debug(err)
	}
}

func (ct *Controller) GetAction(path string, pos int) string {
	path = strings.Trim(path, "/")
	pathsplit := strings.Split(path, "/")
	if pos >= 0 && pos < len(pathsplit) {
		return pathsplit[pos]
	}
	return ""
}

func (ct *Controller) PostReader(i interface{}) (r io.Reader, err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return
	}
	r = strings.NewReader(string(b))
	return
}
