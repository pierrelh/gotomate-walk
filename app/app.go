package app

import (
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var aw = &AutomateWindow{}
var fiber = &Fiber{}

// CreateApp Initiate the app
func CreateApp() {

	aw.plbmodel = NewAutomateModel()

	window := MainWindow{
		AssignTo:   &aw.mw,
		Icon:       "icon.ico",
		Title:      "Gotomate",
		Background: SolidColorBrush{Color: walk.RGB(11, 11, 11)},
		MinSize:    Size{Width: 320, Height: 240},
		Size:       Size{Width: 800, Height: 600},
		Layout:     VBox{MarginsZero: true, SpacingZero: true},
		Children: []Widget{
			HSplitter{
				MaxSize: Size{Height: 150},
				Children: []Widget{
					ListBox{
						Name:            "PrimaryList",
						Font:            Font{Family: "Segoe UI", PointSize: 9},
						MultiSelection:  false,
						AssignTo:        &aw.plb,
						Model:           aw.plbmodel,
						OnItemActivated: aw.plbItemActivated,
					},
					ListBox{
						Name:            "SecondaryList",
						Font:            Font{Family: "Segoe UI", PointSize: 9},
						MultiSelection:  false,
						AssignTo:        &aw.slb,
						OnItemActivated: aw.slbItemActivated,
						BindingMember:   "test",
						DisplayMember:   "test",
					},
				},
			},
			ScrollView{
				AssignTo:        &aw.sv,
				Layout:          VBox{MarginsZero: true},
				Background:      SolidColorBrush{Color: walk.RGB(11, 11, 11)},
				HorizontalFixed: false,
				VerticalFixed:   false,
			},
			PushButton{
				AssignTo:   &aw.pb,
				Font:       Font{Family: "Segoe UI", PointSize: 9},
				Background: SolidColorBrush{Color: walk.RGB(106, 215, 229)},
				Text:       "RUN",
				OnClicked: func() {
					go runFiber()
				},
			},
		},
	}
	window.Run()
}

// AutomateWindow Setting the automate window structure
type AutomateWindow struct {
	mw         *walk.MainWindow
	plb        *walk.ListBox
	plbmodel   *AutomateModel
	slb        *walk.ListBox
	slbmodel   *AutomateModel
	pb         *walk.PushButton
	sv         *walk.ScrollView
	pushButton *walk.PushButton
	compose    *walk.Composite
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

		fiberButton, _ := CreatePushButton(aw.sv, funcName, packageName)

		newInstruction := &FiberInstruction{
			Package:  packageName,
			FuncName: funcName,
			Button:   fiberButton,
		}
		fiber.Instructions = append(fiber.Instructions, newInstruction)
	}
}

// CreatePushButton Create a new push button in the fiber
func CreatePushButton(parent walk.Container, funcName, packageName string) (*FiberButton, error) {
	compose, err := walk.NewComposite(parent)
	if err != nil {
		return nil, err
	}

	compose.SetLayout(walk.NewVBoxLayout())
	layout := compose.Layout()
	layout.SetMargins(walk.Margins{HNear: 100, VNear: 30, HFar: 100, VFar: 30})

	imageView, err := walk.NewImageView(compose)
	if err != nil {
		return nil, err
	}

	imageView.SetMode(0)
	imageView.SetSize(walk.Size{Width: 64, Height: 64})

	path, err := walk.NewImageFromFile(walk.Resources.RootDirPath() + "/func-icons/" + packageName + ".png")
	if err != nil {
		path, _ = walk.NewImageFromFile(walk.Resources.RootDirPath() + "/func-icons/default.png")
	}

	if err := imageView.SetImage(path); err == nil {
		imageView.SetImage(path)
	}

	linkLabel, err := walk.NewLinkLabel(compose)
	if err != nil {
		return nil, err
	}

	bg, err := walk.NewSolidColorBrush(walk.RGB(106, 215, 229))
	if err == nil {
		compose.SetBackground(bg)
		linkLabel.SetBackground(bg)
	}

	font, err := walk.NewFont("Segoe UI", 9, walk.FontBold)
	if err == nil {
		linkLabel.SetFont(font)
	}
	linkLabel.SetText(funcName)
	linkLabel.SetWidth(500)

	dialog := CreateNewDialog(funcName)

	fiberButton := &FiberButton{
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
		return nil, err
	}

	if err := walk.InitWrapperWindow(fiberImageView); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(fiberLinkLabel); err != nil {
		return nil, err
	}

	return fiberButton, nil
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
