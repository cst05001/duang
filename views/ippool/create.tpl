<!DOCTYPE HTML>
<html>
<head>
	<script src="/js/jquery-1.11.3.min.js"></script>
</head>
<body>
	<form id="frm_create">
		<input id="ip" type="text" /><br />
		<input id="submit" type="button" value="submit" />
	</form>
	<script>
		$("#submit").click(function() {
			
			var o = new Object();
			o.ip = $("#ip").val();
			
			jQuery.ajax({
				url: "/ippool/create",
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