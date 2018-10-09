// some scripts

// jquery ready start
$(document).ready(function() {
	// jQuery code
	
	// Bootstrap tooltip
	$("[data-toggle='tooltip']").tooltip();
	
	////scroll top
	var scroll_btn = $("a[href='#top']");
	scroll_btn.hide();
		$(window).scroll(function(){
			if ($(this).scrollTop() > 500) {
				scroll_btn.fadeIn();
			} else {
				scroll_btn.fadeOut();
			}
	    });
	    scroll_btn.click(function () {
			$("html, body").animate({ scrollTop: 0 }, "slow");               
			return false;
	});
	
	$('.open-left').click(function (e) {
		e.preventDefault();
		//Enable sidebar push menu
		if ($(window).width() > 620) {
		  $("body").toggleClass('sidebar-hide');
		 // alert($(window).width());
		}
		else {
			$("body").addClass('sidebar-hide');
		}
    });
  
  
  	$(window).resize(function () {
		if ($(window).width() < 620) {
		$("body").addClass('sidebar-hide');
		// alert($(window).width());
		}
		else {
			$("body").removeClass('sidebar-hide');
		}
	});
	
	$(".sidebar-menu > li > a").click(function (e) {
    //Get the clicked link and the next element
			var $this = $(this);
			var checkElement = $this.next();
				if ((checkElement.is('.sub-menu')) && (checkElement.is(':visible'))) {
			  //Close the menu
			  checkElement.slideUp();
			  checkElement.parent("li").removeClass("active");
			}
			//If the menu is not visible
			else if ((checkElement.is('.sub-menu')) && (!checkElement.is(':visible'))) {
			   //Open the target menu and add the menu-open class
			  
			  var parent = $this.parents('ul').first();
			  //Close all open menus within the parent
			  $(".sidebar-menu > li").removeClass("active");
			  
			  parent.find('ul:visible').slideUp('normal');
			  
			  checkElement.slideDown();
			  checkElement.parent("li").addClass("active");
			  
			  
			}
			//if this isn't a link, prevent the page from being redirected
			if (checkElement.is('.sub-menu')) {
			  e.preventDefault();
			}
	
	});


	function readURL(input) {

    if (input.files && input.files[0]) {
        var reader = new FileReader();

        reader.onload = function (e) {
            $('#img_avatar').attr('src', e.target.result);
        }

        reader.readAsDataURL(input.files[0]);
    }
	}

	$("#img_input").change(function(){
		readURL(this);
	});

	
}); 
// jquery end