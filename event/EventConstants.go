package event

const (
	// SelectNumScroll 对于数独单元格的选择数字滚动事件
	SelectNumScroll = "selectedNum::scroll"
	// SelectNumTap 对于下方数字按钮的点击事件
	SelectNumTap = "selectedNum::tap"
	// NumCellTap 数独数字框体点击事件
	NumCellTap = "numCell::tap"
	// GameStart 对局开始
	GameStart = "game::start"
	// GameVictory 游戏胜利
	GameVictory = "game::victory"
	// GameRefresh 对局刷新
	GameRefresh = "game::refresh"
	// TimeStart 计时开式
	TimeStart = "time::start"
	// TimeStop 计时终止
	TimeStop = "time::stop"
	// TimeReStart 计时重置
	TimeReStart = "time::restart"
	// NumberFillCompleted 数字填充完成
	NumberFillCompleted = "number::fillCompleted"
	// CandidatesView 查看数独对应的候选数字信息
	CandidatesView = "candidates::view"
	// CandidatesHide 候选数字隐藏
	CandidatesHide = "candidates::hide"
)
