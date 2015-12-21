$(document).ready(function() {
  // Full page
  $('#fullpage').fullpage({
    scrollBar: true,
    controlArrows: true,
		scrollingSpeed: 1000,
    keyboardScrolling: true
  });
  // Countdown
  $(".countdown").countdown("2016/01/01", function(event) {
    $(this).text(
       event.strftime('%D Ngày %H Giờ %M Phút %S Giây')
    );
  });
  // Section 1
  var click = false
  $(".face.back").hide();
  $("#section-1 .btn-view").click(function() {
    if (!click) {
      $("#invitaion-card").css({
        "transform": "rotateY(180deg)"
      })
      $(".face.front").hide();
      $(".face.back").show();
      $("#timeline").find(".work").each(function(index, element) {
        $(element).delay(500 * index).fadeTo(1000, 1);
      });
      $(".btn-view").text("Xem mặt trước");
      click = true
    } else {
      $("#invitaion-card").css({
        "transform": "rotateY(0deg)"
      });
      $(".face.front").show();
      $(".face.back").hide();
      $("#timeline").find(".work").each(function(index, element) {
        $(element).hide();
      });
      $(".btn-view").text("Xem mặt sau");
      click = false
    }
  })
  // Section 3
  $("#section-3 .btn").on("click", function() {
    var form = $("#section-3 form");
    var name = $("#section-3 #name").val();
    var email = $("#section-3 #email").val();
    var content = $("#section-3 #textarea").val();
    $.ajax({
      url: '/bless',
      data: {
        name: name,
        email: email,
        content: content
      },
      beforeSend: function() {
        $("#section-3 .btn").addClass("disabled");
      },
      method: 'POST',
      success: function(data){
        $("#section-3 .btn").removeClass("disabled");
        console.log(data);
      }
    })
  })
})
