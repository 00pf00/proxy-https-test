package log

import (
	"00pf00/proxy-https-test/pkg/conf"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"log"
	"time"
)
const logFlushFreqFlagName = "log-flush-frequency"

var logFlushFreq = pflag.Duration(logFlushFreqFlagName, 5*time.Second, "Maximum number of seconds between log flushes")
type KlogWriter struct {

}
func (writer KlogWriter) Write(data []byte) (n int, err error) {
	klog.InfoDepth(1, string(data))
	return len(data), nil
}
// InitLogs initializes logs the way we want for kubernetes.
func InitLogs() {
	klog.InitFlags(nil)
	log.SetOutput(KlogWriter{})
	log.SetFlags(0)
	flag.Set("log_dir",conf.HttpsConf.Log.Dir)
	flag.Set("log_file",time.Now().String()+conf.HttpsConf.Log.File)
	flag.Parse()
	// The default glog flush interval is 5 seconds.
	go wait.Forever(klog.Flush, *logFlushFreq)
}

// FlushLogs flushes logs immediately.
func FlushLogs() {
	klog.Flush()
}

// NewLogger creates a new log.Logger which sends logs to klog.Info.
func NewLogger(prefix string) *log.Logger {
	return log.New(KlogWriter{}, prefix, 0)
}

// GlogSetter is a setter to set glog level.
func GlogSetter(val string) (string, error) {
	var level klog.Level
	if err := level.Set(val); err != nil {
		return "", fmt.Errorf("failed set klog.logging.verbosity %s: %v", val, err)
	}
	return fmt.Sprintf("successfully set klog.logging.verbosity to %s", val), nil
}
