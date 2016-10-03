{{template "header" .}}

<div class="container-fluid">
  <div class="row">
    <div class="main">
      <h1 class="page-header pull-left">Users</h1>
      <button type="submit" class="btn btn-primary pull-right">New User</button>
      <div class="clearfix"></div>
      {{if .ShowDisabled}}
      <small>[<a href="/users">Hide Disabled Users</a>]</small>
      {{else}}
      <small>[<a href="/users?disabled=true">Show Disabled Users</a>]</small>
      {{end}}
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
              <td><button type="button" class="btn btn-default" onclick="editUser({{.Username}})">Edit</button>
              {{if .Enabled}}
                  <button type="button" class="btn btn-danger" onclick="showDisableUserModal({{.Username}})">Disable</button></td>
              {{else}}
                  <button type="button" class="btn btn-success" onclick="showEnableUserModal({{.Username}})">Enable</button></td>
              {{end}}
            </tr>
            {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div> <!-- /container -->

<!-- Disable User Modal -->
<div class="modal fade" id="disableUserModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <form method="POST">
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
          <input type="hidden" name="action" value="disable"/>
          <input type="hidden" id="form_disable_username" name="username"/>
          <button type="submit" class="btn btn-danger">Disable User</button>
        </div>
      </div>
    </div>
  </form>
</div>

<!-- Enable User Modal -->
<div class="modal fade" id="enableUserModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <form method="POST">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">Enable User</h4>
        </div>
        <div class="modal-body">
          Are you sure you want to enable the user <span id="enable_username">username</span>? This user will be able to log in again.
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
          <input type="hidden" name="action" value="enable"/>
          <input type="hidden" id="form_enable_username" name="username"/>
          <button type="submit" class="btn btn-success">Enable User</button>
        </div>
      </div>
    </div>
  </form>
</div>

<script>

  function showDisableUserModal(username) {
    $('#disableUserModal').modal();
    $('#disable_username').text(username)
    $('#form_disable_username').val(username)
}

  function showEnableUserModal(username) {
    $('#enableUserModal').modal();
    $('#enable_username').text(username)
    $('#form_enable_username').val(username)
  }

  function editUser(username) {
    window.location = '/users/edit/' + username;
  }

</script>
