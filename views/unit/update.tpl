<!DOCTYPE HTML>
<html>
<head>
	<script src="/js/jquery-1.11.3.min.js"></script>
</head>
<body>
	<form id="frm_unit">
		<input name="name" id="name" type="text" placeholder="name" value="{{$.Unit.Name}}"/><br />
		<input name="image" id="image" type="text" placeholder="image" value="{{$.Unit.Image}}"/><br />
		<input name="number" id="number" type="number" placeholder="number" value="{{$.Unit.Number}}"/><br />
		{{range $_, $p := $.Unit.Parameteres}}
		<div class="div_parameteres">
		<!-- 此处有bug，等 Type 表实现了一定要改 -->
		<input name="parameteres[]" id="parameteres" class="input_parameter" type="text" placeholder="parameter" value="{{$p.Value}}" /><select class="select_parameter"><option value="v">Volume</option></select><input id="btn_add_parameter" class="btn_add_parameter" type="button" value="+" /><input id="btn_del_parameter" class="btn_del_parameter" type="button" value="-" />
		</div>
		{{end}}
		<div class="div_parameteres">
		<!-- 此处有bug，等 Type 表实现了一定要改 -->
		<input name="parameteres[]" id="parameteres" class="input_parameter" type="text" placeholder="parameter" /><select class="select_parameter"><option value="v">Volume</option></select><input id="btn_add_parameter" class="btn_add_parameter" type="button" value="+" /><input id="btn_del_parameter" class="btn_del_parameter" type="button" value="-" />
		</div>		

		<input id="submit" type="button" value="submit" />
 	</form>
	<script>
		function add_parameter() {
			$(this).parent().after('		<div class="div_parameteres"> \
		<input name="parameteres[]" id="parameteres" class="input_parameter" type="text" placeholder="parameter" /><select class="select_parameter"><option value="v">Volume</option><input id="btn_add_parameter" class="btn_add_parameter" type="button" value="+" /><input id="btn_del_parameter" class="btn_del_parameter" type="button" value="-" /> \
		</div>');

			$(".btn_add_parameter").off("click");
			$(".btn_add_parameter").on("click", add_parameter);
		}
		$(".btn_add_parameter").click(add_parameter);
		
		$(".btn_del_parameter").click(function() {
			$(this).parent().remove();
		});

		$("#submit").click(function() {
			
			var unit = new Object();
			unit.name = $("#name").val();
			unit.image = $("#image").val();
			unit.number = parseInt($("#number").val());
			unit.parameteres = new Array();
			$(".div_parameteres").each(function() {
				var p = new Object()
				p.value = $(this).children(".input_parameter").val();
				p.type = $(this).children(".select_parameter").val();
				unit.parameteres.push(p);
			});

			
			jQuery.ajax({
				url: "/unit/update/{{$.Unit.Id}}",
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