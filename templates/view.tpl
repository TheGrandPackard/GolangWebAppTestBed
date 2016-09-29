{{template "header" .}}

  <div class="container">
    <h1>{{.Page.Title}}
			{{if .Site.User }}
        {{if .Site.User.ManagePages }}
        <small>[<a href="/edit/{{.Page.Title}}">edit</a>]</small>
        {{end}}
      {{end}}
    </h1>
    <div>{{.Page.GetBody}}</div>
  </div> <!-- /container -->
