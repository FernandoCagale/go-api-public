package checker

import (
	mgo "gopkg.in/mgo.v2"
)

type mongodb struct {
	uri string
}

func NewMongodb(uri string) *mongodb {
	return &mongodb{uri}
}

func (m *mongodb) IsAlive() bool {
	Host := []string{
		m.uri,
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		FailFast: true,
	})
	if err != nil {
		return false
	}

	defer session.Close()

	if err = session.Ping(); err != nil {
		return false
	}

	return true
}
