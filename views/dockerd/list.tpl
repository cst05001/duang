<!DOCTYPE HTML>
<html>
<head>
	<script src="/js/jquery-1.11.3.min.js"></script>
</head>
<body>
{{range $_, $dockerd := $.DockerdList}}
	<h1>{{$dockerd.Addr}}</h1>
{{end}}
</body>
</html>