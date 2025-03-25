package application

type App struct {
	mode        string
	version     string // 版本号
	initialized bool   // 是否初始化
}

// Version 应用版本号
func (t *App) Version() string {
	return t.version
}

func (t *App) IsDebug() bool {
	return t.mode == "debug"
}

// Initialize() 初始化应用
func (t *App) Initialize() {
	t.initialized = true
}

func (t *App) Initialized() bool {
	return t.initialized
}

func New() *App {
	return &App{
		version: "1.0.0",
	}
}
