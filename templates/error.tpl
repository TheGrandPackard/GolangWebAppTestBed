{{template "header" .}}

  <div class="container">
      {{if .Site.Error}}
      <div class="alert alert-danger" style="margin-top: 15px;">{{.Site.Error}}</div>
      {{end}}
  </div> <!-- /container -->
