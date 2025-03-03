package bootstrap

import (
	"0121_1/global"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func casbinInit() *casbin.Enforcer {
	db, err := gormadapter.NewAdapterByDB(global.App.DB)
	if err != nil {
		global.App.Log.Error(fmt.Sprintf("failed to initialize Casbin adapter: %w", err))
		return nil
	}

	enforcer, err := casbin.NewEnforcer("model.conf", db)
	if err != nil {
		global.App.Log.Error(fmt.Sprintf("failed to initialize Casbin enforcer: %w", err))
		return nil
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		global.App.Log.Error(fmt.Sprintf("failed to load Casbin policy: %w", err))
		return nil
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

	return enforcer
}
