package GUI

import (
	"VideoSpider/paraseURL"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"log"
	"strings"
)

var Te, userId *walk.TextEdit
var Mw = new(MyMainWindow)

type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string
	path         string
}

func (mw *MyMainWindow) getUserIdAndOperate() {
	users := userId.Text()
	userId.SetText(users + "shoudaoshuju")

	userIds := strings.Split(users,",")
	//myContext := context.Background()

	fmt.Println(userIds)
	for _,userId := range userIds{
		go paraseURL.ParseId(userId)
	}

}

func UserOperate() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("%s", err)
			//log.Warning(err)
		}
	}()
	if err := (MainWindow{
		AssignTo: &Mw.MainWindow,
		Title:    "NoteBook",
		MenuItems: []MenuItem{
			Menu{
				Text: "&选择榜单",
				Items: []MenuItem{
					Action{
						//AssignTo: &openAction,
						Text: "&打榜动态",
						//Image: "/img/open.png",
						//OnTriggered: Mw.OpenLoopAction_Triggered,
					},
					Action{
						//AssignTo: &openAction,
						Text: "&飙升榜",
						//Image: "/img/open.png",
						//OnTriggered: Opener.Mw.OpenAction_Triggered,
					},

					Action{
						Text:"日榜",
					},

					Action{
						Text:"周榜",
					},

					Separator{},
					Action{
						Text:        "Exit",
						OnTriggered: func() { Mw.Close() },
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text: "About",
						//OnTriggered: Abouter.AboutAction_Triggered,
					},
				},
			},
		},
		MinSize: Size{720, 540},
		Size:    Size{800, 600},
		Layout:  VBox{MarginsZero: true},

		ToolBar: ToolBar{
			Name:    "shiyangongju",
			Visible: true,
			ContextMenuItems: []MenuItem{
				Menu{
					Text: "&Help",
					Items: []MenuItem{
						Action{
							Text: "About",
							//OnTriggered: Abouter.AboutAction_Triggered,
						},
					},
				},
			},
			MaxSize: Size{800, 600},
			Items: []MenuItem{
				Action{
					//AssignTo: &openAction,
					Text: "&选择日期",
					//Image: "/img/open.png",
					//OnTriggered: Mw.OpenLoopAction_Triggered,
				},
			},
		},

		Children: []Widget{
			VSplitter{
				Children: []Widget{
					Label{Text: "玩家ID:"},
					TextEdit{
						AssignTo: &userId,
						MinSize:  Size{Width: 100, Height: 40},
						MaxSize:  Size{Width: 100, Height: 40},
					},
				},
			},
			PushButton{
				Text:    "开始爬取数据",
				MinSize: Size{Width: 100, Height: 40},
				Font:    Font{PointSize: 20},
				OnClicked: func() {
					Mw.getUserIdAndOperate()
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
	Mw.MainWindow.SetVisible(true)
	Mw.MainWindow.Run()
}
