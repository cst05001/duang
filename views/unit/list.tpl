<!DOCTYPE HTML>
<html>
<head>
	<script src="/js/jquery-1.11.3.min.js"></script>
</head>
<body>
{{range $_, $unit := $.UnitList}}
	<h1>{{$unit.Name}}</h1>
{{end}}
</body>
</html>