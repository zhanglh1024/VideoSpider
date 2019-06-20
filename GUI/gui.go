package GUI

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"log"
)

var(
	userId *walk.TextEdit
	timeText *walk.TextEdit
	runStopBtn *walk.PushButton
	pauseRecoverBtn *walk.PushButton
)
var Mw = new(MyMainWindow)

//type MyMainWindow struct {
//	*walk.MainWindow
//	prevFilePath string
//	path         string
//}
//
//func (mw *MyMainWindow) getUserIdAndOperate() {
//	users := userId.Text()
//	userId.SetText(users + "shoudaoshuju")
//
//	userIds := strings.Split(users, ",")
//	//myContext := context.Background()
//
//	fmt.Println(userIds)
//	for _, userId := range userIds {
//		go paraseURL.ParseId(userId)
//	}
//
//}

func UserOperate() {
	if err := (MainWindow{
		AssignTo: &Mw.MainWindow,
		Title:    "酷狗视频爬虫",
		MinSize: Size{720, 540},
		Size:    Size{800, 600},
		Layout:  VBox{MarginsZero: true},

		Children: []Widget{
			Composite{
				MinSize:Size{Width: 720, Height:200},
				Layout: HBox{},
				Children: []Widget{
					VSplitter{
						MaxSize: Size{Width: 120, Height: 100},
						Children:[]Widget{
							PushButton{
								Text:    "打榜动态",
								OnClicked: func() {
									Mw.dynamicList()
								},
							},
						},
					},
					VSplitter{
						MaxSize: Size{Width: 120, Height: 100},
						Children:[]Widget{
							PushButton{
								Text:    "飙升榜",
								OnClicked: func() {
									Mw.soaringList()
								},
							},
						},
					},
					//VSplitter{
					//	MaxSize: Size{Width: 120, Height: 100},
					//	Children:[]Widget{
					//		PushButton{
					//			Text:    "日榜",
					//			OnClicked: func() {
					//				Mw.dayList()
					//			},
					//		},
					//	},
					//},
					//VSplitter{
					//	MaxSize: Size{Width: 120, Height: 50},
					//	Children:[]Widget{
					//		PushButton{
					//			Text:    "周榜",
					//			OnClicked: func() {
					//				Mw.mouthList()
					//			},
					//		},
					//	},
					//},
				},
			},

			VSplitter{
				//ColumnSpan: 1,
				MaxSize:    Size{550, 10},
				Children: []Widget{
					Label{Text: "输入时间:"},
					TextEdit{
						AssignTo: &timeText,
						MaxSize:  Size{Width: 550, Height: 20},
					},
				},
			},

			Composite{
				MinSize:    Size{550, 100},
				Layout:HBox{},
				Children:[]Widget{
					VSplitter{
						MaxSize: Size{Width: 120, Height: 100},
						Children:[]Widget{
							PushButton{
								Text:    "日榜",
								OnClicked: func() {
									Mw.dayList()
								},
							},
						},
					},
					VSplitter{
						MaxSize: Size{Width: 120, Height: 50},
						Children:[]Widget{
							PushButton{
								Text:    "周榜",
								OnClicked: func() {
									Mw.mouthList()
								},
							},
						},
					},
				},
			},

			VSplitter{
				ColumnSpan: 1,
				MinSize:    Size{550, 100},
				Children: []Widget{
					Label{Text: "玩家ID:"},
					TextEdit{
						AssignTo: &userId,
						MaxSize:  Size{Width: 100, Height: 40},
					},
				},
			},
			Composite{
				MinSize:    Size{550, 100},
				Layout:HBox{},
				Children:[]Widget{
					VSplitter{
						MaxSize: Size{90, 50},
						Children: []Widget{
							PushButton{
								Text:      "退出",
								AssignTo:  &pauseRecoverBtn,
								OnClicked: Mw.offlinePauseRecover,
							},
						},
					},
					VSplitter{
						MaxSize: Size{90, 50},
						Children: []Widget{
							PushButton{
								Text:      "开始运行",
								AssignTo:  &runStopBtn,
								OnClicked: Mw.getUserIdAndOperate,
							},
						},
					},
				},
			},
		},
	}).Create(); err != nil {
		log.Fatal(err)
	}

	screenX := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	screenY := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	if err := Mw.SetBounds(walk.Rectangle{
		X: (screenX - Mw.MainWindow.Width()) / 2,
		Y: (screenY - Mw.MainWindow.Height()) / 2,
	}); err != nil {
		log.Print(err)
		return
	}

	//pauseRecoverBtn.SetVisible(false)
	Mw.MainWindow.SetVisible(true)
	Mw.MainWindow.Run()
}
