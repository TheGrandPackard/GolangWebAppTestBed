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
              <th>Version</th>
              <th>Title</th>
            </tr>
          </thead>
          <tbody>
            {{range .Pages}}
            <tr>
              <td>{{.ID}}</td>
              <td>{{.Version}}</td>
              <td><a href="view/{{.Title}}">{{.Title}}<a/></td>
              <td><button type="button" class="btn btn-default" onclick="editPage({{.Title}})">Edit</button>
                  <button type="button" class="btn btn-danger" onclick="deletePage({{.Title}})">Delete</button>
              </td>
            </tr>
            {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div> <!-- /container -->

<!-- Delete Page Modal -->
<div class="modal fade" id="deletePageModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <form method="POST">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">Disable User</h4>
        </div>
        <div class="modal-body">
          Are you sure you want to disable the page <span id="delete_title">title</span>? This page cannot be undeleted. Make sure you have backups available before making any major changes.
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
          <input type="hidden" id="form_delete_title" name="title"/>
          <button type="submit" class="btn btn-danger">Delete Page</button>
        </div>
      </div>
    </div>
  </form>
</div>

<script>

function deletePage(title) {
  $('#deletePageModal').modal();
  $('#delete_title').text(title)
  $('#form_delete_title').val(title)
}

function editPage(title) {
  window.location = '/edit/' + title;
}

</script>
