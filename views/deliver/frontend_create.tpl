<!DOCTYPE HTML>
<html>
<head>
	<script src="/js/jquery-1.11.3.min.js"></script>
</head>
<body>
	<form id="frm_create">
		<input id="name" type="text" placeholder="frontend name" /><br />
		<input id="bind" type="text" placeholder="e.g. 192.168.3.168:5000" /><br />
		<input id="submit" type="button" value="submit" />
	</form>
	<script>
		$("#submit").click(function() {
			
			var o = new Object();
			o.name = $("#name").val();
			o.bind = $("#bind").val();
			
			jQuery.ajax({
				url: "/deliver/frontend/create",
				type: "post",
				data: JSON.stringify(o),
				dataType: "json",
				success: function(data) {
					console.log(data);
				},
			})	
		});		
	</script>
</body>
</html>