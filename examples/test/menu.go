package main

import (
	"net/url"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/murlokswarm/app"
)

func init() {
	app.Import(&Menu{})
}

// Menu is a component to test menu based elements.
type Menu struct {
	DisableAll       bool
	RandomTitle      uuid.UUID
	Separator        bool
	RenderRootToggle bool
	RenderToggle     bool
}

// Render statisfies the app.Component interface.
func (m *Menu) Render() string {
	return `
<menu>
	<menuitem label="button" 
			  onclick="OnButtonClick"
			  {{if .DisableAll}}disabled{{end}}>
	</menuitem>
	<menuitem label="button with icon" 
			  onclick="OnButtonWithIconClick" 
			  {{if .DisableAll}}disabled{{end}}>
	</menuitem>
	<menuitem label="random button: {{.RandomTitle}}" 
			  onclick="OnButtonWithRandomTitleClicked"
			  {{if .DisableAll}}disabled{{end}}>
	</menuitem>

	<menuitem separator></menuitem>

	<menuitem label="set dock badge" 
			  onclick="OnSetDockBadge"
			  {{if .DisableAll}}disabled{{end}}>
	</menuitem>
	<menuitem label="unset dock badge" 
			  onclick="OnUnsetDockBadge"
			  {{if .DisableAll}}disabled{{end}}>
	</menuitem>
	
	<menuitem separator></menuitem>

	<menuitem label="set dock icon" 
			  onclick="OnSetDockIcon"
			  {{if .DisableAll}}disabled{{end}}>
	</menuitem>
	<menuitem label="unset dock icon" 
			  onclick="OnUnsetDockIcon"
			  {{if .DisableAll}}disabled{{end}}>
	</menuitem>
	
	<menuitem separator></menuitem>

	<menu label="submenu" {{if .DisableAll}}disabled{{end}}>
		<menuitem label="sub button" onclick="OnSubButtonClick"></menuitem>
		<menuitem label="sub button without action"></menuitem>	
	</menu>
	<menu label="disabled submenu" disabled></menu>
	<menu label="random menu: {{.RandomTitle}}" disabled></menu>
	
	<menuitem separator></menuitem>

	<menu label="separator test" {{if .DisableAll}}disabled{{end}}>
		<menuitem label="switch separator" 
				  onclick="OnSwitchSeparatorClick">
		</menuitem>

		{{if .Separator}}
			<menuitem separator></menuitem>		
		{{else}}
			<menuitem label="----- Separator -----"></menuitem>	
		{{end}}

		<menuitem label="disabled button" disabled></menuitem>		
	</menu>

	<menuitem separator></menuitem>

	<menu label="full render test" {{if .DisableAll}}disabled{{end}}>
		<menuitem label="render root" onclick="OnRenderRootClicked"></menuitem>
		{{if .RenderToggle}}
			<menuitem label="render menu" onclick="OnRenderTest"></menuitem>
		{{else}}
			<menu label="render item">
				<menuitem label="render" onclick="OnRenderTest"></menuitem>
			</menu>
		{{end}}
	</menu>

	{{if .RenderRootToggle}}
		<menuitem label="blank button to test root render" disabled></menuitem>
	{{end}}

	<menuitem separator></menuitem>
	
	<menuitem label="enable all" 
			onclick="OnEnableAllClick"
			{{if not .DisableAll}}disabled{{end}}>
	</menuitem>
	<menuitem label="disable all" 
			  onclick="OnDisableAllClick"
			  {{if .DisableAll}}disabled{{end}}>
	</menuitem>	
</menu>
	`
}

// OnNavigate is the function that is called when the component is navigated on.
func (m *Menu) OnNavigate(u *url.URL) {
	m.RandomTitle = uuid.New()
	app.Render(m)
}

// OnButtonClick is the function that is called when the button labelled
// "button" is clicked.
func (m *Menu) OnButtonClick() {
	app.DefaultLogger.Log("button clicked")
}

// OnButtonWithIconClick is the function that is called when the button labelled
// "button with icon" is clicked.
func (m *Menu) OnButtonWithIconClick() {
	app.DefaultLogger.Log("button with icon clicked")
}

// OnSetDockBadge is the function that is called when the button labelled "set
// dock badge" is clicked.
func (m *Menu) OnSetDockBadge() {
	app.DefaultLogger.Log("button set dock badge clicked")

	if app.SupportsDock() {
		app.Dock().SetBadge(uuid.New())
	}
}

// OnUnsetDockBadge is the function that is called when the button labelled
// "unset dock badge" is clicked.
func (m *Menu) OnUnsetDockBadge() {
	app.DefaultLogger.Log("button unset dock badge clicked")

	if app.SupportsDock() {
		app.Dock().SetBadge(nil)
	}
}

// OnSetDockIcon is the function that is called when the button labelled "set
// dock icon" is clicked.
func (m *Menu) OnSetDockIcon() {
	app.DefaultLogger.Log("button set dock icon clicked")

	if app.SupportsDock() {
		app.Dock().SetIcon(filepath.Join(app.Resources(), "logo.png"))
	}
}

// OnUnsetDockIcon is the function that is called when the button labelled
// "unset dock icon" is clicked.
func (m *Menu) OnUnsetDockIcon() {
	app.DefaultLogger.Log("button unset dock icon clicked")

	if app.SupportsDock() {
		app.Dock().SetIcon("")
	}
}

// OnButtonWithRandomTitleClicked is the function that is called when the button
// with randow title is clicked.
func (m *Menu) OnButtonWithRandomTitleClicked() {
	app.DefaultLogger.Log("button with random title clicked")
	m.RandomTitle = uuid.New()
	app.Render(m)
}

// OnSubButtonClick is the function that is called when the button labelled "sub
// button" is clicked.
func (m *Menu) OnSubButtonClick() {
	app.DefaultLogger.Log("sub button clicked")
}

// OnEnableAllClick is the function that is called when the button labelled
// "enable all" is clicked.
func (m *Menu) OnEnableAllClick() {
	app.DefaultLogger.Log("button enable all clicked")
	m.DisableAll = false
	app.Render(m)
}

// OnDisableAllClick is the function that is called when the button labelled
// "disable all" is clicked.
func (m *Menu) OnDisableAllClick() {
	app.DefaultLogger.Log("button disable all clicked")
	m.DisableAll = true
	app.Render(m)
}

// OnSwitchSeparatorClick is the function that is called when the button
// labelled "switch separator" is clicked.
func (m *Menu) OnSwitchSeparatorClick() {
	app.DefaultLogger.Log("button switch separator clicked")
	m.Separator = !m.Separator
	app.Render(m)
}

// OnRenderRootClicked is the function that is called when the button labelled
// "render root" is clicked.
func (m *Menu) OnRenderRootClicked() {
	m.RenderRootToggle = !m.RenderRootToggle
	app.Render(m)
}

// OnRenderTest is the function that is called when a button form the render
// test menu is clicked.
func (m *Menu) OnRenderTest() {
	m.RenderToggle = !m.RenderToggle
	app.Render(m)
}
