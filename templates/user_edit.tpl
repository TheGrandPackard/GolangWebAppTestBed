{{template "header" .}}

  <div class="container">

    <form class="form-signup form-horizontal" method="POST" action="/users">
      <h2 class="form-signup-heading">Edit User: {{.User.Username}}</h2>

      <div class="form-group">
        <label for="inputEmail3" class="col-sm-2 control-label">Email Address</label>
        <div class="col-sm-10">
          <input type="email" class="form-control" id="inputEmail3" placeholder="Email Address" name="username" value="{{.User.Username}}">
        </div>
      </div>

      <div class="form-group">
        <label for="inputEmail3" class="col-sm-2 control-label">Enabled</label>
        <div class="col-sm-10">
          <input type="checkbox" name="enabled" {{if .User.Enabled}}checked{{end}}>
        </div>
      </div>

      <div class="form-group">
        <label for="inputEmail3" class="col-sm-2 control-label">Manage Users</label>
        <div class="col-sm-10">
          <input type="checkbox" name="manageUsers" {{if .User.ManageUsers}}checked{{end}}>
        </div>
      </div>
      
      <div class="form-group">
        <label for="inputEmail3" class="col-sm-2 control-label">Manage Pages</label>
        <div class="col-sm-10">
          <input type="checkbox" name="managePages" {{if .User.ManagePages}}checked{{end}}>
        </div>
      </div>

      <input type="hidden" name="action" value="edit"/>

      <div class="form-group">
        <div class="col-sm-offset-2 col-sm-10">
          <button type="submit" class="btn btn-default">Save User</button>
          <button type="submit" class="btn btn-default" onclick="window.location='/users'; return false;">Cancel Changes</button>
        </div>
      </div>

      {{if .Site.Error}}
      <div class="alert alert-danger" style="margin-top: 15px;">{{.Site.Error}}</div>
      {{end}}
    </form>

  </div> <!-- /container -->
