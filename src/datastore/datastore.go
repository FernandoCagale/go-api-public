package datastore

import (
	mgo "gopkg.in/mgo.v2"
)

func New(connection string) (*mgo.Session, error) {
	Host := []string{
		connection,
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		FailFast: true,
	})

	if err != nil {
		return nil, err
	}

	return session, nil
}
