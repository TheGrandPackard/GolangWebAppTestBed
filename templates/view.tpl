{{template "header" .}}

  <div class="container">

    <h1>{{.Page.Title}}</h1>

    <p>[<a href="/edit/{{.Page.Title}}">edit</a>]</p>

    <div class=".mce-content-body">{{.Page.Body}}</div>

  </div> <!-- /container -->
