<nav class="navbar navbar-fixed-top navbar-bootsnipp animate">
	<div class="container-fluid" style="background-color: #000000">
		<div class="navbar-header">
			<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
			<span class="sr-only">Toggle navigation</span>
			<span class="icon-bar"></span>
			<span class="icon-bar"></span>
			<span class="icon-bar"></span>
			</button>
			<a href="/" class="navbar-brand">Wiki</a>
		</div>
		<div class="collapse navbar-collapse" id="nav-header">
		<ul class="nav navbar-nav pull-left">
			{{if .Site.User }}
				{{if .Site.User.ManagePages }}
				<li><a href="/pages">Pages</a></li>
				{{end}}
					{{if .Site.User.ManageUsers }}
					<li><a href="/users">Users</a></li>
					{{end}}
				<li><a href="/logout">Log Out</a></li>
			{{else}}
				<li><a href="/login">Sign In</a></li>
				<li><a href="/signup">Sign Up</a></li>
			{{end}}
		</ul>
			<form class="navbar-form pull-right" action="/search" role="search" method="GET">
				<div class="form-group">
					<input type="text" class="form-control" name="query" placeholder="Search">
				</div>
				<button type="submit" class="btn btn-default">Search</button>
			</form>
			<div class="clearfix"></div>
		</div>
	</div>
</nav>
