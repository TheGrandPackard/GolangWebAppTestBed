{{template "header" .}}

  <div class="container">

    <form class="form-signup" method="POST">
      <h2 class="form-signup-heading">Create your account</h2>

      <div class="form-group">
        <label for="inputEmail">Email address</label>
        <input type="email" id="inputEmail" class="form-control" placeholder="Email address" name="username" required autofocus>
      </div>

      <div class="form-group">
        <label for="inputPassword">Password</label>
        <input type="password" id="inputPassword" class="form-control" placeholder="Password" name="password" required>
      </div>

      <div class="form-group">
        <label for="inputRepeatPassword">Repeat Password</label>
        <input type="password" id="inputRepeatPassword" class="form-control" placeholder="Repeat Password" name="repeat_password" required>
      </div>

      <button class="btn btn-lg btn-success btn-block" type="submit">Create an Account</button>

      {{if .Site.Error}}
      <div class="alert alert-danger" style="margin-top: 15px;">{{.Site.Error}}</div>
      {{end}}
    </form>

  </div> <!-- /container -->
