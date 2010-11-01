function coord(n) { return 15 + n *30 }

function drawBoard(id) {
  var canvas = document.getElementById(id);
  var ctx = canvas.getContext('2d');
  ctx.lineWidth = 3;
  ctx.lineCap = 'round';
  ctx.lineJoin = 'round';
  ctx.lineColor = '#000';

  ctx.fillStyle = '#F93';
  ctx.fillRect(0, 0, 570, 570);

  ctx.fillStyle = '#000';

  ctx.beginPath();
  for (var i = 0; i < 19; i++) {
    ctx.moveTo(coord( 0), coord( i));
    ctx.lineTo(coord(18), coord( i));
    ctx.moveTo(coord( i), coord( 0));
    ctx.lineTo(coord( i), coord(18));
  }
  ctx.stroke();
  ctx.closePath();

  ctx.fillColor = '#000';
  $([[3,3], [3,9], [3,15], [9,3], [9,9], [9,15], [15,3], [15,9], [15,15]]).each(function(i, xy) {
    ctx.beginPath();
    ctx.arc(coord(xy[0]), coord(xy[1]), 5, 0, 2*Math.PI, false);
    ctx.closePath();
    ctx.fill();
  });
}

function putPiece(id, color, x, y) {
  var canvas = document.getElementById(id);
  var ctx = canvas.getContext('2d');
  ctx.beginPath();
  ctx.lineColor = '#000';
  ctx.fillStyle = color;
  ctx.arc(coord(x), coord(y), 14, 0, 2*Math.PI, true);
  ctx.fill();
  ctx.stroke();
  ctx.closePath();
}

$(document).ready(function() {
  drawBoard('board');
  //putPiece('board', 'black', 0, 0);
  //putPiece('board', 'white', 4, 3);

  $('#board').click(function(evt) {
    var x = Math.floor((evt.pageX - 30.0) / 30.0);
    var y = Math.floor((evt.pageY - 30.0) / 30.0);
    $.ajax({
      type: 'POST',
      url: '/put',
      data: {x:(x+1), y:(y+1)},
      //dataType: 'json',
      cache: false,
      success: function(dt) {
        var data = eval(dt);
        drawBoard('board');
        for (var i = 0; i < data.size; i++) {
          for (var j = 0; j < data.size; j++) {
            var cell = data.board[i][j];
            if (cell == '@') {
              putPiece('board', 'black', j, i);
            }
            else if (cell == 'O') {
              putPiece('board', 'white', j, i);
            }
          }
        }
      }
    });
  });
});
