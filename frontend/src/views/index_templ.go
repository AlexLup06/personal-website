// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Layout() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><title>Alex Lupatsiy - Portfolio and Blog</title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"Welcome to My personal website. Explore projects, blogs, and professional background.\"><meta property=\"og:site_name\" content=\"Alex Lupatsiy\"><link rel=\"stylesheet\" href=\"/css/tailwind.css\"><link rel=\"icon\" type=\"image/png\" href=\"/public/favicon/favicon-96x96.png\" sizes=\"96x96\"><link rel=\"icon\" type=\"image/svg+xml\" href=\"/public/favicon/favicon.svg\"><link rel=\"shortcut icon\" href=\"/public/favicon/favicon.ico\"><link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/public/favicon/apple-touch-icon.png\"><link rel=\"manifest\" href=\"/public/favicon/site.webmanifest\"><link rel=\"preload\" href=\"/public/fonts/FunnelSans-VariableFont_wght.ttf\" as=\"font\" type=\"font/ttf\" crossorigin=\"anonymous\"><style>\n\t\t\t\t@font-face {\n\t\t\t\t\tfont-family: 'FunnelSans';\n\t\t\t\t\tsrc: url('/public/fonts/FunnelSans-VariableFont_wght.ttf') format('opentype');\n\t\t\t\t\tfont-display: swap;\n\t\t\t\t}\n\t\t\t\tbody {\n\t\t\t\t\tfont-family: 'FunnelSans', Arial, sans-serif;\n\t\t\t\t}\n\t\t\t</style><script type=\"application/ld+json\">\n\t\t\t{\n\t\t\t\t\"@context\": \"https://schema.org\",\n\t\t\t\t\"@type\": \"WebSite\",\n\t\t\t\t\"name\": \"Alex Lupatsiy\",\n\t\t\t\t\"url\": \"https://alexlupatsiy.com\",\n\t\t\t\t\"sameAs\": [\n\t\t\t\t\t\"https://www.linkedin.com/in/alex-lupatsiy-2730b41b5/\",\n\t\t\t\t\t\"https://github.com/AlexLup06\"\n\t\t\t\t]\n\t\t\t}\n\t\t\t</script></head><body class=\"m-0\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</body><script src=\"/js/htmx/htmx.min.js\" defer></script><script src=\"/js/build/bundle.js\" type=\"module\" defer></script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
