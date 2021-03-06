{{template "header" .}}

  <div class="container">

    <form class="form-signin" method="POST">
      <h2 class="form-signin-heading">Please sign in</h2>
      <label for="inputEmail" class="sr-only">Email address</label>
      <input type="email" id="inputEmail" class="form-control" placeholder="Email address" name="username" required autofocus>
      <label for="inputPassword" class="sr-only">Password</label>
      <input type="password" id="inputPassword" class="form-control" placeholder="Password" name="password" required>
      <div class="checkbox">
        <label>
          <input type="checkbox" value="remember-me"> Remember me
        </label>
      </div>
      <button class="btn btn-lg btn-primary btn-block" type="submit">Sign In</button>
      <button class="btn btn-lg btn-success btn-block" type="submit" onclick="window.location='/signup'; return false;">Sign Up</button>

      {{if .Site.Error}}
      <div class="alert alert-danger" style="margin-top: 15px;">{{.Site.Error}}</div>
      {{end}}
    </form>

  </div> <!-- /container -->
