package main

import "sync"

/*
 Simple interface that allows us to switch out both implementations of the Manager
*/
type ConfigManager interface {
	Set(*Config)
	Get() *Config
	Close()
}

/*
 This struct manages the configuration instance by
 preforming locking around access to the Config struct.
*/
type MutexConfigManager struct {
	conf  *Config
	mutex *sync.Mutex
}

func NewMutexConfigManager(conf *Config) *MutexConfigManager {
	return &MutexConfigManager{conf, &sync.Mutex{}}
}

func (self *MutexConfigManager) Set(conf *Config) {
	self.mutex.Lock()
	self.conf = conf
	self.mutex.Unlock()
}

func (self *MutexConfigManager) Get() *Config {
	self.mutex.Lock()
	temp := self.conf
	self.mutex.Unlock()
	return temp
}

func (self *MutexConfigManager) Close() {
	//Do Nothing
}
