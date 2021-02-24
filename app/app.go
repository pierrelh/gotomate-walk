package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

var aw = &AutomateWindow{}
var fiber = &Fiber{}

// CreateApp Initiate the app
func CreateApp() {

	aw.plbmodel = NewAutomateModel()
	var openAction *walk.Action

	if err := (declarative.MainWindow{
		AssignTo:   &aw.mw,
		Icon:       "icon.ico",
		Title:      "Gotomate",
		Background: declarative.SolidColorBrush{Color: walk.RGB(11, 11, 11)},
		MinSize:    declarative.Size{Width: 320, Height: 240},
		Size:       declarative.Size{Width: 800, Height: 600},
		Layout:     declarative.VBox{MarginsZero: true, SpacingZero: true},
		MenuItems: []declarative.MenuItem{
			declarative.Menu{
				Text: "&File",
				Items: []declarative.MenuItem{
					declarative.Action{
						AssignTo: &openAction,
						Text:     "Open",
						Image:    "/open.png",
						Enabled:  declarative.Bind("enabledCB.Checked"),
						Visible:  declarative.Bind("!openHiddenCB.Checked"),
						Shortcut: declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
						// OnTriggered: aw.openActionTriggered,
					},
					declarative.Menu{
						AssignTo: &aw.folders,
						Image:    "/folder.png",
						Text:     "My Fibers",
					},
					declarative.Action{
						Text:     "Run",
						Image:    "/run.png",
						Shortcut: declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyR},
						OnTriggered: func() {
							go runFiber()
						},
					},
					declarative.Action{
						Text:        "Save",
						Image:       "/save.png",
						Shortcut:    declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyS},
						OnTriggered: aw.saveFiber,
					},
					declarative.Action{
						Text:        "Exit",
						OnTriggered: func() { aw.mw.Close() },
					},
				},
			},
		},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout:  declarative.HBox{MarginsZero: true, SpacingZero: true},
				MaxSize: declarative.Size{Height: 120},
				Children: []declarative.Widget{
					declarative.ListBox{
						Name:            "PrimaryList",
						Font:            declarative.Font{Family: "Roboto", PointSize: 9},
						MultiSelection:  false,
						AssignTo:        &aw.plb,
						Model:           aw.plbmodel,
						OnItemActivated: aw.plbItemActivated,
					},
					declarative.ListBox{
						Name:            "SecondaryList",
						Font:            declarative.Font{Family: "Roboto", PointSize: 9},
						MultiSelection:  false,
						AssignTo:        &aw.slb,
						OnItemActivated: aw.slbItemActivated,
					},
					declarative.Composite{
						Layout: declarative.VBox{},
						Children: []declarative.Widget{
							declarative.Composite{
								Layout: declarative.HBox{},
								Children: []declarative.Widget{
									declarative.Label{
										Alignment: declarative.Alignment2D(walk.AlignHFarVCenter),
										Font:      declarative.Font{Family: "Roboto", PointSize: 12, Underline: true, Bold: true},
										Text:      "Fiber Name :",
										TextColor: walk.Color(0xffffff),
									},
									declarative.TextEdit{
										AssignTo:      &aw.name,
										Alignment:     declarative.Alignment2D(walk.AlignHNearVCenter),
										Font:          declarative.Font{Family: "Roboto", PointSize: 9},
										CompactHeight: true,
										MaxSize:       declarative.Size{Width: 150},
									},
								},
							},
							declarative.Composite{
								Layout: declarative.HBox{},
								Children: []declarative.Widget{
									declarative.PushButton{
										MaxSize:    declarative.Size{Width: 100},
										Font:       declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
										Background: declarative.SolidColorBrush{Color: walk.RGB(106, 215, 229)},
										Text:       "RUN",
										OnClicked: func() {
											go runFiber()
										},
									},
									declarative.PushButton{
										MaxSize:    declarative.Size{Width: 100},
										Font:       declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
										Background: declarative.SolidColorBrush{Color: walk.RGB(106, 215, 229)},
										Text:       "Save",
										OnClicked:  aw.saveFiber,
									},
								},
							},
						},
					},
				},
			},
			declarative.ScrollView{
				AssignTo:        &aw.sv,
				Column:          5,
				Layout:          declarative.Flow{Margins: declarative.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Background:      declarative.SolidColorBrush{Color: walk.RGB(11, 11, 11)},
				HorizontalFixed: false,
				VerticalFixed:   false,
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	AddSavedFibersActions()
	aw.mw.Run()
}

// AutomateWindow Setting the automate window structure
type AutomateWindow struct {
	mw       *walk.MainWindow
	folders  *walk.Menu
	name     *walk.TextEdit
	plb      *walk.ListBox
	plbmodel *AutomateModel
	slb      *walk.ListBox
	slbmodel *AutomateModel
	sv       *walk.ScrollView
}

// AutomateItem Setting automate packages items
type AutomateItem struct {
	name  string
	value string
}

// AutomateModel Setting the model of automate ListBox
type AutomateModel struct {
	walk.ListModelBase
	items []AutomateItem
}

func (aw *AutomateWindow) saveFiber() {
	name := aw.name.Text()
	if name == "" {
		var dlg *walk.Dialog
		var acceptPB *walk.PushButton

		errDialog := declarative.Dialog{
			Icon:          "/icon.ico",
			Title:         "Error",
			AssignTo:      &dlg,
			DefaultButton: &acceptPB,
			MinSize: declarative.Size{
				Width:  200,
				Height: 150,
			},
			Layout: declarative.VBox{},
			Children: []declarative.Widget{
				declarative.Composite{
					Layout: declarative.VBox{},
					Children: []declarative.Widget{
						declarative.TextLabel{
							Text:      "No name given for the fiber",
							Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
							Font:      declarative.Font{Family: "Roboto", PointSize: 9},
						},
						declarative.PushButton{
							AssignTo: &acceptPB,
							Text:     "OK",
							Font:     declarative.Font{Family: "Roboto", PointSize: 9},
							OnClicked: func() {
								dlg.Cancel()
								aw.name.SetFocus()
							},
						},
					},
				},
			},
		}

		errDialog.Run(aw.mw)
	} else {
		path := "saves/" + name + ".json"
		file, _ := json.Marshal(fiber)
		ioutil.WriteFile(path, file, 0644)
		AddSavedFibersActions()
	}
}

func (aw *AutomateWindow) plbItemActivated() {
	i := aw.plb.CurrentIndex()
	if i != -1 {
		item := &aw.plbmodel.items[i]
		value := item.name
		newModel := NewAutomateSubModel(value)
		aw.slb.SetModel(newModel)
		aw.slbmodel = newModel
	}
}

func (aw *AutomateWindow) slbItemActivated() {
	i := aw.slb.CurrentIndex()
	if i != -1 {
		plbIndex := aw.plb.CurrentIndex()
		plbItem := &aw.plbmodel.items[plbIndex]
		packageName := plbItem.name

		item := &aw.slbmodel.items[i]
		funcName := item.name

		data, dialog := CreateNewDialog(funcName)

		newInstruction := &FiberInstruction{
			Package:  packageName,
			FuncName: funcName,
			Data:     data,
		}
		CreatePushButton(aw.sv, funcName, packageName, dialog)
		fiber.Instructions = append(fiber.Instructions, newInstruction)
	}
}

// CreatePushButton Create a new push button in the fiber
func CreatePushButton(parent *walk.ScrollView, funcName, packageName string, dialog declarative.Dialog) error {
	compose, err := walk.NewComposite(parent)
	if err != nil {
		return err
	}

	compose.SetLayout(walk.NewVBoxLayout())
	compose.SetMinMaxSizePixels(walk.Size{Width: 300, Height: 150}, walk.Size{Width: 300, Height: 150})
	bg, err := walk.NewSolidColorBrush(walk.RGB(106, 215, 229))
	if err == nil {
		compose.SetBackground(bg)
	}

	imageView, err := walk.NewImageView(compose)
	if err != nil {
		return err
	}

	imageView.SetMinMaxSizePixels(walk.Size{Width: 64, Height: 64}, walk.Size{Width: 64, Height: 64})
	imageView.SetMode(3)

	path, err := walk.NewImageFromFile(walk.Resources.RootDirPath() + "/func-icons/" + packageName + ".png")
	if err != nil {
		path, _ = walk.NewImageFromFile(walk.Resources.RootDirPath() + "/func-icons/default.png")
	}

	if err := imageView.SetImage(path); err == nil {
		imageView.SetImage(path)
	}

	linkLabel, err := walk.NewLinkLabel(compose)
	if err != nil {
		return err
	}

	font, err := walk.NewFont("Roboto", 9, walk.FontBold)
	if err == nil {
		linkLabel.SetFont(font)
	}
	linkLabel.SetText(funcName)

	_ = &FiberButton{
		compose,
		imageView,
		linkLabel,
		dialog,
	}

	fiberComposite := &FiberComposite{
		compose,
		dialog,
	}

	fiberImageView := &FiberImageView{
		imageView,
		dialog,
	}

	fiberLinkLabel := &FiberLinkLabel{
		linkLabel,
		dialog,
	}

	if err := walk.InitWrapperWindow(fiberComposite); err != nil {
		return err
	}

	if err := walk.InitWrapperWindow(fiberImageView); err != nil {
		return err
	}

	if err := walk.InitWrapperWindow(fiberLinkLabel); err != nil {
		return err
	}

	return nil
}

// NewAutomateModel Getting the automates packages
func NewAutomateModel() *AutomateModel {
	env := Packages

	m := &AutomateModel{items: make([]AutomateItem, len(env))}

	for i, e := range env {
		name := strconv.Itoa(i)
		value := e

		m.items[i] = AutomateItem{value, name}
	}

	return m
}

// NewAutomateSubModel Gettings the automate's subpackages
func NewAutomateSubModel(key string) *AutomateModel {
	env := SubPackages

	m := &AutomateModel{items: make([]AutomateItem, len(env[key]))}

	for i, e := range env[key] {
		j := strings.Index(e, "=")
		if j == 0 {
			continue
		}

		name := strconv.Itoa(i)
		value := e

		m.items[i] = AutomateItem{value, name}
	}

	return m
}

// ItemCount return the length of an AutomateModel items
func (m *AutomateModel) ItemCount() int {
	return len(m.items)
}

// Value return the value of an AutomateModel item
func (m *AutomateModel) Value(index int) interface{} {
	return m.items[index].name
}
