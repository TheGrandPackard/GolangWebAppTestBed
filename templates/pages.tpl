{{template "header" .}}

<div class="container-fluid">
  <div class="row">
    <div class="main">
      <h1 class="page-header">Wiki Pages</h1>
      <div class="table-responsive">
        <table class="table table-striped">
          <thead>
            <tr>
              <th>ID</th>
              <th>Title</th>
              <th>Body</th>
            </tr>
          </thead>
          <tbody>
            {{range .Pages}}
            <tr>
              <td>{{.ID}}</td>
              <td><a href="view/{{.Title}}">{{.Title}}<a/></td>
              <td>{{.GetBody}}</td>
            </tr>
            {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div> <!-- /container -->
