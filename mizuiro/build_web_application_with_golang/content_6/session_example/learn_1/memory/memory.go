package memory

import (
	"container/list"
	"sync"
	"time"
)

var pder = Provider{list: list.New()}

type Provider struct {
	lock     sync.Mutex // lock
	sessions map[string]*list.Element
	list     *list.List
}

type SessionStore struct {
	sid          string                      // session id unique identifier
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // value on session data
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

func (pder *Provider) SessionInit(sid string) (Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}
func (pder *Provider)