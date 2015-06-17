<!DOCTYPE HTML>
<html>
<head>
	<script src="http://ajax.aspnetcdn.com/ajax/jQuery/jquery-1.8.0.js"></script>
</head>
<body>
	<form>
		<input name="name" id="name" type="text" placeholder="name"/><br />
		<input name="image" id="image" type="text" placeholder="image "/><br />
		<input name="number" id="number" type="number" placeholder="number" /><br />
		<input name="parameter" id="parameter" type="text" placeholder="parameter" /><br />
		<input id="submit" type="button" value="submit" />
 	</form>
	<script>
		$("#submit").click(function() {
			var unit = new Object();
			unit.name = $("#name").val();
			unit.image = $("#image").val();
			unit.number = parseInt($("#number").val());
			unit.parameter = $("#parameter").val().split(" ");
			jQuery.ajax({
				url: "/unit/create",
				type: "post",
				data: JSON.stringify(unit),
				success: function(data, textStatus, jqXHR) {
					alert(data);
				}(),
			})	
		});

	</script>
</body>
</html>