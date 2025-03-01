// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package homepage

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "alexlupatsiy.com/personal-website/frontend/src/views/components"

func LandingPage() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"sm:h-[calc(100vh-96px)] min-h-[700px] flex\"><div class=\"h-full w-full hidden md:flex flex-col justify-center items-center py-10 pb-14\"><div class=\"relative flex-1 w-full h-full flex flex-col justify-center items-center py-10 pb-14\"><div class=\"bg-main-500 h-3/4 rounded-3xl absolute left-0 top-0 w-[calc(50%+20px)] shadow-xl z-10\"><div class=\"flex h-full flex-col justify-between px-14 py-16 lg:px-16 xl:px-24 xl:py-20 text-white\"><div><h1 class=\"text-6xl lg:text-7xl\">I <span class=\"underline\">solve</span> <span>&nbsp</span><br class=\"2xl:hidden\">problems!</h1><h2 class=\"text-3xl lg:text-3xl mt-6\">Even the difficult ones</h2></div><div class=\"2xl:mx-auto 2xl:mb-10\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = getInTouchButton().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></div></div><div class=\"bg-[#DE672B] h-3/4 rounded-3xl absolute right-0 bottom-0 w-[calc(50%+20px)] shadow-md \n\t\t\t\t\t\t\tflex flex-col pl-20 pr-12 py-10 lg:pl-20 lg:pr-12 lg:py-14\"><div class=\"text-white text-lg lg:text-xl [&amp;&gt;p]:mb-5\"><p class=\"text-2xl\">Hey there! &nbsp; 👋🏻</p><p class=\"\">I'm Alex, a software engineer based in Dortmund, Germany.</p><p class=\"\">I'm passionate about building software that solves real-world problems and helps people.</p><p class=\"\">Currently, I'm finishing my Master's degree in Computer Science with a focus on Robotics at TU Dortmund.</p></div><div class=\"h-16 flex justify-end mt-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Signature().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div></div></div></div><div class=\"flex flex-1 md:hidden flex-col gap-6 justify-center items-center\"><div class=\"relative flex-1 w-full h-full flex flex-col justify-center items-center mt-4\"><div class=\"bg-main-500 h-80 sm:h-[calc(50%+10px)] rounded-3xl absolute left-0 top-0 w-10/12 shadow-xl z-10\"><div class=\"flex flex-col h-full justify-between p-6 xs:p-8 sm:px-10 sm:py-10 text-white\"><div class=\"flex flex-col gap-4\"><h1 class=\"text-4xl sm:text-6xl\">I <span class=\"underline\">solve</span> problems!</h1><h2 class=\"text-xl sm:text-3xl\">Even the difficult ones</h2></div><div class=\"mx-auto xs:mx-0\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = getInTouchButton().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</div></div></div><div class=\"bg-[#DE672B] h-[calc(50%+10px)] rounded-3xl absolute right-0 top-[calc(320px-20px)] sm:top-auto sm:bottom-0 w-10/12 shadow-md \n\t\t\t\t\t\t\tflex flex-col px-8 pb-8 pt-10\"><div class=\"text-white text-lg [&amp;&gt;p]:mb-2.5\"><p class=\"text-xl\">Hey there! &nbsp; 👋🏻</p><p class=\"\">I'm Alex, a software engineer based in Dortmund, Germany.</p><p class=\"\">I'm passionate about building software that solves real-world problems and helps people.</p></div><div class=\"h-12 flex justify-end mt-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Signature().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func getInTouchButton() templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<div class=\"bg-white flex w-fit shadow-md transition-all duration-700 px-3 py-2.5 sm:px-6 sm:py-4 \n\t\t\t\tjustify-center gap-2 lg:gap-4 items-center rounded-2xl overflow-hidden\"><p class=\"text-lg sm:text-xl text-main-800\">Get in touch!</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.EmailButtonSM().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
