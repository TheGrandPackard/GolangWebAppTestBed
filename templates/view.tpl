{{template "header" .}}

  <div class="container">

    <h1>{{.Page.Title}} <small>[<a href="/edit/{{.Page.Title}}">edit</a>]</small></h1>

    <div>{{.Page.GetBody}}</div>

  </div> <!-- /container -->
