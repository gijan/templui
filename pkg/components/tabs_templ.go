// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

type Tab struct {
	// ID uniquely identifies the tab
	ID string

	// Title displays in tab button
	Title string

	// Content renders when tab is active
	Content templ.Component
}

type TabsProps struct {
	// Tabs defines the tabs structure and content
	Tabs []Tab

	// TabsContainerClass adds classes to tabs header
	TabsContainerClass string

	// ContentContainerClass adds classes to content area
	ContentContainerClass string
}

// Tabs renders a tabbed interface with animated transitions.
// Uses Alpine.js for state management and interactions.
// For detailed examples and usage guides, visit https://goilerplate.com/docs/components/tabs
//
// Props:
// - Tabs: Tab definitions and content
// - TabsContainerClass: Additional classes for header
// - ContentContainerClass: Additional classes for content
//
// Features:
// - Animated tab switching
// - Keyboard navigation
// - Responsive layout
// - ARIA support
func Tabs(props TabsProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{     \n      tabSelected: &#39;1&#39;,\n      tabId: $id(&#39;tabs&#39;),\n      tabButtonClicked(tabButton) {\n        this.tabSelected = tabButton.id.replace(this.tabId + &#39;-&#39;, &#39;&#39;);\n        this.tabRepositionMarker(tabButton);\n      },\n      tabRepositionMarker(tabButton) {\n        this.$refs.tabMarker.style.width = tabButton.offsetWidth + &#39;px&#39;;\n        this.$refs.tabMarker.style.height = tabButton.offsetHeight + &#39;px&#39;;\n        this.$refs.tabMarker.style.left = tabButton.offsetLeft + &#39;px&#39;;\n      },\n      tabContentActive(tabContent) {\n        return this.tabSelected === tabContent.id.replace(this.tabId + &#39;-content-&#39;, &#39;&#39;);\n      }\n    }\" x-init=\"$nextTick(() =&gt; tabRepositionMarker($refs.tabButtons.firstElementChild));\" class=\"relative\"><!-- Tabs buttons container -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{"relative flex items-center justify-center h-10 p-1 rounded-lg select-none",
			"bg-muted text-muted-foreground",
			props.TabsContainerClass}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-ref=\"tabButtons\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/tabs.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><!-- Individual tab buttons -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, tab := range props.Tabs {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button :id=\"$id(tabId)\" @click=\"tabButtonClicked($el);\" type=\"button\" class=\"relative z-20 flex-1 inline-flex items-center justify-center h-8 px-3 text-sm font-medium transition-all rounded-md cursor-pointer whitespace-nowrap\" :class=\"{&#39;text-foreground bg-background shadow-sm&#39;: tabSelected === &#39;{tab.ID}&#39;, &#39;hover:text-foreground&#39;: tabSelected !== &#39;{tab.ID}&#39;}\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(tab.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/tabs.templ`, Line: 76, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Active tab marker --><div x-ref=\"tabMarker\" class=\"absolute left-0 z-10 h-full duration-300 ease-out\" x-cloak><div class=\"w-full h-full bg-background rounded-md shadow-sm\"></div></div></div><!-- Tab content container -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 = []any{"relative mt-2 content", props.ContentContainerClass}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/tabs.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><!-- Individual tab content -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, tab := range props.Tabs {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div :id=\"$id(tabId + &#39;-content&#39;)\" x-show=\"tabContentActive($el)\" class=\"relative\" x-cloak>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = tab.Content.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
