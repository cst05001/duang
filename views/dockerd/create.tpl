<!DOCTYPE HTML>
<html>
<head>
	<script src="/js/jquery-1.11.3.min.js"></script>
</head>
<body>
	<form id="frm_dockerd">
		<input name="addr" id="addr" type="text" placeholder="addr: http://192.168.3.147:2375"/><br />
		<input id="submit" type="button" value="submit" />
 	</form>
	<script>
		$("#submit").click(function() {
			
			var dockerd = new Object();
			dockerd.addr = $("#addr").val();
			
			jQuery.ajax({
				url: "/dockerd/create",
				type: "post",
				data: JSON.stringify(dockerd),
				dataType: "json",
				success: function(data) {
					console.log(data);
				},
			})	
		});
	</script>
</body>
</html>