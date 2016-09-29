{{template "header" .}}

<div class="container-fluid">
  <div class="row">
    <div class="main">
      <h1 class="page-header pull-left">Users</h1>
      <button type="submit" class="btn btn-primary pull-right">New User</button>
      <div class="clearfix"></div>
      <div class="table-responsive">
        <table class="table table-striped table-hover">
          <thead>
            <tr>
              <th>ID</th>
              <th>Username</th>
              <th>Date Added</th>
              <th>Date Last Login</th>
              <th>Manage Users</th>
              <th>Manage Pages</th>
            </tr>
          </thead>
          <tbody>
            {{range .Users}}
            <tr>
              <td>{{.ID}}</td>
              <td>{{.Username}}</td>
              <td>{{.DateAdded}}</td>
              <td>{{.DateLastLogin}}</td>
              <td>{{.ManageUsers}}</td>
              <td>{{.ManagePages}}</td>
              <td><button type="button" class="btn btn-default">Edit</button>
                  <button type="button" class="btn btn-danger" onclick="showDisableUserModal({{.Username}},{{.ID}})">Disable</button></td>
            </tr>
            {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div> <!-- /container -->

<!-- Delete User Modal -->
<div class="modal fade" id="disableUserModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="myModalLabel">Disable User</h4>
      </div>
      <div class="modal-body">
        Are you sure you want to disable the user <span id="disable_username">username</span>? This user will no longer be able to log in until an administrator enables their account again.
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
        <button type="button" class="btn btn-danger" onclick="disableUser()">Disable User</button>
      </div>
    </div>
  </div>
</div>

<script>

  function showDisableUserModal(username, id) {
    $('#disableUserModal').modal();
    $('#disable_username').text(username)
  }

  function disableUser() {
    console.log("POST User Disable Request Here")
  }

</script>
