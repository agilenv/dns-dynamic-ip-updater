package track

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	timeLayout         = time.RFC1123Z
	delimiter          = "\t"
	defaultFileStorage = "linkip_tracks.log"
	fileEnvvar         = "TRACK_FILE"
)

type fileStorage struct {
	filepath string
}

func NewFileStorage() *fileStorage {
	filepath := defaultFileStorage
	if os.Getenv(fileEnvvar) != "" {
		filepath = os.Getenv(fileEnvvar)
	}
	return &fileStorage{
		filepath: filepath,
	}
}

func (f *fileStorage) Save(event Event) error {
	file, err := os.OpenFile(f.filepath, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}

	data := strings.Join([]string{
		time.Now().Format(timeLayout),
		event.IP,
		event.PublicAPI,
	}, delimiter)

	_, err = file.WriteString(data + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (f *fileStorage) LastEvent() Event {
	data, err := ioutil.ReadFile(f.filepath)
	if err != nil {
		return Event{}
	}
	return parseEvent(string(data))
}

func parseEvent(data string) Event {
	pieces := strings.Split(data, delimiter)
	eventTime, err := time.Parse(timeLayout, pieces[0])
	if err != nil {
		return Event{}
	}
	return NewEvent(eventTime, pieces[1], pieces[2])
}
