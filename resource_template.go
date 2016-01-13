package gorma

const resourceTmpl = `
{{ $typename  := .TypeName }}
{{ $belongs := .BelongsTo }}
{{ $ml := .MediaLower }}
{{ if .DoMedia }}
{{ $version := .APIVersion }}
{{ range $idx, $action := .TypeDef.Actions  }}
{{ if hasusertype $action }}
func {{$typename}}From{{version $version}}{{title $action.Name}}Payload(ctx *{{$version}}.{{title $action.Name}}{{$typename}}Context) {{$typename}} {
	var {{$ml}} {{$typename}}
	payload := ctx.Payload
	{{ $atd :=  attributefields $action.Payload.Definition }}
	{{ range $pf := typefields $action.Payload }}
	{{if $pf.Pointer}} {{$pf.Name}} := payload.{{ range $cols := $atd }}{{ if eq $pf.Name $cols.Name }}{{goify $cols.Name true}}{{end}}{{end}}
	{{$ml}}.{{ goify $pf.Name true}} = {{$pf.Name}}
	//pointer? {{ $pf.Pointer }}
	{{ else }}
	// Not Pointer {{$pf.Name }} {{$pf.Pointer}}
	{{$ml}}.{{ goify $pf.Name true}} = {{ range $cols := $atd }}{{ if eq $pf.Name $cols.Name }}{{if $cols.Pointer }}*{{end}}payload.{{goify $cols.Name true}}{{end}}{{end}}
	{{ end }}
	{{ end }}
/*
	{{ range $pf := typefields $action.Payload }}
	{{ $pf.Name }} - {{$pf.Pointer}} / {{ range $cols := $atd }}{{ if eq $pf.Name $cols.Name }}{{$cols.Name }} - {{$cols.Pointer}}{{end}}{{end}}
	{{ end }}

*/
	return {{$ml}}	
}
{{ end }}{{end}}{{end}}
`
