package systemicdb_service

import (
	"errors"
	"github.com/vmihailenco/msgpack/v5"
	"strings"
	"time"
)

func (ks *SystemicDBService) GetByKey(key string) (string, error) {

	var returnDatum string

	datumBytes := ks.keyStore.Get(key)
	if datumBytes == nil {
		return returnDatum, errors.New("key not found")
	}

	err := msgpack.Unmarshal(datumBytes.Value, &returnDatum)
	if err != nil {
		return returnDatum, err
	}

	return returnDatum, nil
}

func (ks *SystemicDBService) SetByKey(key string, datum string, expiry int32) error {

	if len(strings.TrimSpace(key)) < 1 || len(strings.TrimSpace(datum)) < 1 {
		return errors.New("key and datum cannot be blank")
	}

	mpd, err := msgpack.Marshal(datum)
	if err != nil {
		return err
	}

	ks.keyStore.Insert(key, mpd, time.Second*time.Duration(expiry))

	return nil
}

func (ks *SystemicDBService) DelByKey(key string) error {
	ks.keyStore.Remove(key)

	return nil
}
