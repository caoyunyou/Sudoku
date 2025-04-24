package event

const (
	// SelectedNumChange 选中数字变更
	SelectedNumChange = "selectedNum::change"
	// GameVictory 游戏胜利
	GameVictory = "game::victory"
	// GameRefresh 对局刷新
	GameRefresh = "game::refresh"
	// GameUndoStep 回退到上一步
	GameUndoStep = "gameUndoStep::undoStep"
	// TimeStart 计时开式
	TimeStart = "time::start"
	// TimeStop 计时终止
	TimeStop = "time::stop"
	// TimeReStart 计时重新计算
	TimeReStart = "time::restart"
	// NumberFillCompleted 数字填充完成
	NumberFillCompleted = "number::fillCompleted"
	// CandidatesView 查看数独对应的候选数字信息
	CandidatesView = "candidates::view"
	// CandidatesHide 候选数字隐藏
	CandidatesHide = "candidates::hide"
)
