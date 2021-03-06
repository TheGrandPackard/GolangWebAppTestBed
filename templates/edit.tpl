{{template "header" .}}

<div class="container">

  <h1>Editing {{.Page.Title}}</h1>

  <form action="/save/{{.Page.Title}}" method="POST">
    <div><textarea name="body" rows="20" cols="80">{{printf "%s" .Page.GetBody}}</textarea></div>
    <input type="hidden" name="id" value="{{.Page.ID}}">
    <div>
      <button class="btn btn-primary">Save</button>
      <button class="btn btn-default" onclick="window.history.back(); return;">Cancel</button>
    </div>
  </form>

</div> <!-- /container -->

<script>
  tinymce.init({
  selector: 'textarea',
  height: 500,
  plugins: [
    'advlist autolink lists link image charmap print preview anchor',
    'searchreplace visualblocks code fullscreen',
    'insertdatetime media table contextmenu paste code'
  ],
  toolbar: 'insertfile undo redo | styleselect | bold italic | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | link image'
});
</script>
