var version = -1;

function coord(n) { return 15 + n * 30 }

function drawBoard(id) {
  var canvas = document.getElementById(id);
  var ctx = canvas.getContext('2d');
  ctx.lineWidth = 3;
  ctx.lineCap = 'round';
  ctx.lineJoin = 'round';
  ctx.strokeStyle = '#000';

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
  ctx.stokeStyle = '#000';
  ctx.fillStyle = color;
  ctx.arc(coord(x), coord(y), 14, 0, 2*Math.PI, true);
  ctx.fill();
  ctx.stroke();
  ctx.closePath();
}

function drawConsole(id, turn, agehama) {
  function drawCircle(ctx, lineColor, fillColor, x, y) {
    ctx.beginPath();
    ctx.lineWidth = 5;
    ctx.strokeStyle = lineColor;
    ctx.fillStyle = fillColor;
    ctx.arc(x, y, 25, 0, 2*Math.PI, true);
    ctx.fill();
    ctx.stroke();
    ctx.closePath();
  }

  var canvas = document.getElementById(id);
  var ctx = canvas.getContext('2d');
  ctx.font = "50px 'ＭＳ Ｐゴシック'";

  ctx.fillStyle = '#FFF';
  ctx.fillRect(0, 0, 570, 50);

  drawCircle(ctx, turn==0?"#F00":"#000", "#000", 70, 35);
  drawCircle(ctx, turn==1?"#F00":"#000", "#FFF", 370, 35);
  ctx.strokeStyle = '#000';
  ctx.strokeText('' + agehama[0], 120, 55, 100);
  ctx.strokeText('' + agehama[1], 420, 55, 100);
}

function drawMatch(data) {
  var board = data.board;
  drawBoard('board');

  for (var i = 0; i < board.size; i++) {
    for (var j = 0; j < board.size; j++) {
      var cell = board.board[i][j];
      if      (cell == '@')  putPiece('board', 'black', j, i);
      else if (cell == 'O')  putPiece('board', 'white', j, i);
    }
  }

  drawConsole('console', data.turn, data.agehama);
}

function draw() {
  $.ajax({
    type: 'POST',
    url: '/get',
    cache: false,
    success: function(json) {
      var data = eval(json);
      if (version < data.version) {
        version = data.version;
        drawMatch(data);
      }
      setTimeout(draw, 1000);
    }
  });
}

$(document).ready(function() {
  drawBoard('board');

  $('#board').click(function(evt) {
    //var x = Math.round((evt.pageX - 30.0) / 30.0);
    //var y = Math.round((evt.pageY - 30.0) / 30.0);
    var x = Math.round((evt.pageX - 35.0) / 30.0);
    var y = Math.round((evt.pageY - 35.0) / 30.0);
    $.ajax({
      type: 'POST',
      url: '/put',
      data: {x:(x+1), y:(y+1)},
      cache: false,
      success: function(json) {
        var data = eval(json);
        if (data.message) {
          alert(data.message);
          return;
        }
        if (data.version <= version) return;
        version = data.version;
        drawMatch(data);
      }
    });
  });

  draw();
});
