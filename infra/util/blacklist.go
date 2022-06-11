package util

import (
	"bufio"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"crgo/infra/log"
)

var Blacklist struct {
	sync.RWMutex
	Data map[string]struct{}
}

func init() {
	Blacklist.Data = make(map[string]struct{})
}

//监听黑名单文件
func WatchBlacklist() {
	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigFile(viper.GetString("blacklist.filePath"))
	v.OnConfigChange(onBlacklistChange)
	v.WatchConfig()
}

func onBlacklistChange(in fsnotify.Event) {
	const writeOrCreateMask = fsnotify.Write | fsnotify.Create
	if in.Op&writeOrCreateMask != 0 {
		updateBlacklist()
	}
}
//可以用atomic 替换
func updateBlacklist() {
	filePath := viper.GetString("blacklist.filePath")
	fp, err := os.Open(filePath)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	defer fp.Close()

	data := make(map[string]struct{})
	f := bufio.NewReader(fp)
	for {
		line, _, err := f.ReadLine()
		if err != nil {
			break
		}
		data[string(line)] = struct{}{}
	}
	Blacklist.Lock()
	Blacklist.Data = data
	Blacklist.Unlock()
}

//是都在黑名单中
func InBlacklist(uid string) bool {
	Blacklist.RLock()
	_, ok := Blacklist.Data[uid]
	Blacklist.RUnlock()
	return ok
}
