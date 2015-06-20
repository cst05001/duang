<!DOCTYPE HTML>
<html>
<head>
	<script src="/js/jquery-1.11.3.min.js"></script>
</head>
<body>
	<form>
		<input name="name" id="name" type="text" placeholder="name"/><br />
		<input name="image" id="image" type="text" placeholder="image "/><br />
		<input name="number" id="number" type="number" placeholder="number" /><br />
		<input name="parameteres" id="parameteres" type="text" placeholder="parameter" /><br />
		<input id="submit" type="button" value="submit" />
 	</form>
	<script>
		$("#submit").click(function() {
			var unit = new Object();
			unit.name = $("#name").val();
			unit.image = $("#image").val();
			unit.number = parseInt($("#number").val());
			unit.parameteres = $("#parameteres").val().split(" ");
			jQuery.ajax({
				url: "/unit/create",
				type: "post",
				data: JSON.stringify(unit),
				dataType: "json",
				success: function(data) {
					console.log(data);
				},
			})	
		});

	</script>
</body>
</html>