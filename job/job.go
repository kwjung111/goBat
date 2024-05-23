package job

type Job interface {
	GetContext() jobContext
	SetStep(Step) error
	NextStep() error
}

// 서비스 로직을 포함 시킴
type Step interface {
	SetService()
}

// job을 실행시킴
type JobLauncher interface {
	Launch(*Job) bool
}
