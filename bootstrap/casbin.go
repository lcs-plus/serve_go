package bootstrap

import (
	"0121_1/global"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func casbinInit() {
	db, err := gormadapter.NewAdapterByDB(global.App.DB)
	if err != nil {
		return
	}

	enforcer, err := casbin.NewEnforcer("model.conf", db)
	if err != nil {
		return
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return
	}

	// 添加策略（如果数据库中没有策略）
	policy, _ := enforcer.HasPolicy("admin", "/api/*", "*")
	if !policy {
		_, _ = enforcer.AddPolicy("admin", "/api/*", "*")
	}

	policy, _ = enforcer.HasPolicy("user", "/api/data", "GET")
	if !policy {
		_, _ = enforcer.AddPolicy("user", "/api/data", "GET")
	}

	policy, _ = enforcer.HasPolicy("alice", "admin")
	if !policy {
		_, _ = enforcer.AddGroupingPolicy("alice", "admin")
	}

	policy, _ = enforcer.HasPolicy("bob", "user")
	if !policy {
		_, _ = enforcer.AddGroupingPolicy("bob", "user")
	}
}
