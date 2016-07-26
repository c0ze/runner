package stats

import (
	"time"

	"github.com/Sirupsen/logrus"
)

type Log15Reporter struct {
}

func NewLog15Reporter() *Log15Reporter {
	return (&Log15Reporter{})
}

func (lr *Log15Reporter) report(stats []*collectedStat) {
	for _, s := range stats {
		f := make(logrus.Fields)
		for k, v := range s.Counters {
			f[k] = v
		}
		for k, v := range s.Values {
			f[k] = v
		}
		for k, v := range s.Timers {
			f[k] = time.Duration(v)
		}

		logrus.WithFields(f).Info(s.Name)
	}
}
