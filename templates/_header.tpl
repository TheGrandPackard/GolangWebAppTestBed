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
				<li><a href="/pages">Pages</a></li>
				{{if .Site.Session.Values }}<li><a href="/logout">Log Out</a></li>{{end}}
			</ul>
			<form class="navbar-form pull-right" action="/search/" role="search" method="GET">
				<div class="form-group">
					<input type="text" class="form-control" name="name" placeholder="Search">
				</div>
				<button type="submit" class="btn btn-default">Submit</button>
			</form>
			<div class="clearfix"></div>
		</div>
	</div>
</nav>
