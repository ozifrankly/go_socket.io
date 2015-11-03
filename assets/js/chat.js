var socket1 = io();
socket1.on('chat message', function(msg){
  var ta = $('#ta1');
  ta.html(ta.html()+"\n"+msg);
});

function send1(){
  socket1.emit('chat message', $('#in1').val());
}

var socket2 = io();
socket2.on('chat message', function(msg){
  var ta = $('#ta2');
  ta.html(ta.html()+"\n"+msg);
});

function send2(){
  socket2.emit('chat message', $('#in2').val());
}

$(function(){
  $('#bt1').click(send1);
  $('#bt2').click(send2);
});
