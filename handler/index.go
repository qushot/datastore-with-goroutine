package handler

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Index</title>
	</head>
	<body>
		{{ range . }}
		<a href="{{ .Path }}" target="_blank">{{ .Name }}</a><br>
		{{ end }}
	</body>
</html>`

	t, err := template.New("index").Parse(tpl)
	if err != nil {
		http.Error(w, "template.Parse error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := []struct {
		Name string
		Path template.URL
	}{
		{
			Name: "SetUp",
			Path: template.URL("/setup"),
		},
		{
			Name: "TearDown",
			Path: template.URL("/teardown"),
		},
		{
			Name: "Sync",
			Path: template.URL("/sync"),
		},
		{
			Name: "Async",
			Path: template.URL("/async"),
		},
	}

	t.Execute(w, data)
}
