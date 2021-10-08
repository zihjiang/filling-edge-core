
package common

import (
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

const (
	EdgeIdFile = "/data/edge.id"
)

type RuntimeInfo struct {
	ID           string
	BaseDir      string
	HttpUrl      string
	DPMEnabled   bool
	AppAuthToken string
}

func (r *RuntimeInfo) init() error {
	r.ID = r.getSdeId()
	return nil
}

func (r *RuntimeInfo) getSdeId() string {
	var edgeId string
	if _, err := os.Stat(r.getSdeIdFilePath()); os.IsNotExist(err) {
		f, err := os.Create(r.getSdeIdFilePath())
		check(err)

		defer f.Close()
		edgeId = uuid.NewV4().String()
		_, _ = f.WriteString(edgeId)
	} else {
		buf, err := ioutil.ReadFile(r.getSdeIdFilePath())
		if err != nil {
			log.Fatal(err)
		}
		edgeId = string(buf)
	}

	return edgeId
}

func (r *RuntimeInfo) getSdeIdFilePath() string {
	return r.BaseDir + EdgeIdFile
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewRuntimeInfo(httpUrl string, baseDir string) (*RuntimeInfo, error) {
	runtimeInfo := RuntimeInfo{
		HttpUrl: httpUrl,
		BaseDir: baseDir,
	}
	err := runtimeInfo.init()
	if err != nil {
		return nil, err
	}
	return &runtimeInfo, nil
}
