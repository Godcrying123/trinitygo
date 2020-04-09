package application

import (
	"github.com/jinzhu/gorm"
)

// Context record all thing inside one request
type Context interface {
	GetApplication() Application
	setRuntime(map[string]string)
	GetRuntime() map[string]string
	GetDB() *gorm.DB
	GetTXDB() *gorm.DB
	setDB(*gorm.DB)
	cleanRuntime()
}

// ContextImpl Context impl
type ContextImpl struct {
	app     Application
	runtime map[string]string
	db      *gorm.DB
}

// GetApplication get app
func (c *ContextImpl) GetApplication() Application {
	return c.app
}

// GetRuntime get runtime info
func (c *ContextImpl) GetRuntime() map[string]string {
	return c.runtime
}

// setRuntime set runtime info
func (c *ContextImpl) setRuntime(runtime map[string]string) {
	c.runtime = runtime
}

// GetDB get db instance
func (c *ContextImpl) GetDB() *gorm.DB {
	return c.db
}

// GetTXDB get db instance
func (c *ContextImpl) GetTXDB() *gorm.DB {
	c.db = c.db.Begin()
	return c.db
}

// GetRuntime get runtime info
func (c *ContextImpl) setDB(db *gorm.DB) {
	c.db = db
}

// CleanRuntime  clean runtime info
func (c *ContextImpl) cleanRuntime() {
	c.runtime = nil
	c.db = nil
}

// NewContext new contedt
func NewContext(app Application) Context {
	return &ContextImpl{
		app: app,
	}

}
