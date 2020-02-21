package conf

import "github.com/BurntSushi/toml"

var HttpsConf *Conf

type Conf struct {
	Addr Addr `toml:"addr"`
	Tls  Tls  `toml:"tls"`
	Log  Log  `toml:"log"`
}
type Tls struct {
	Key  string `toml:"key"`
	Cert string `toml:"cert"`
}
type Log struct {
	Dir  string `toml:"dir"`
	File string `toml:"file"`
}
type Addr struct {
	Kubelet string `toml:"kubelet"`
}

func InitConf(path string) error{
	_,err := toml.DecodeFile(path,HttpsConf)
	if err != nil {
		return err
	}
	return nil
}
