/// <reference path="jquery-vsdoc.js" />
/// <reference path="jquery.extend.js" />
/// <reference path="jquery.admin.extend.js" />
document.domain = window.location.hostname;

$.ajaxSetup({
	type: "post",
	timeout: 60000,
	error: function (x, e) {
		if (x.status == 0) {
			$(document).process({ text: "無法連接服務器，可能斷網了。請與管理員聯繫", width: 320 });
		} else if (x.status == 404) {
			$(document).process({ text: "你請求的URL不存在。請與管理員聯繫", width: 320 });
		} else if (x.status == 500) {
			$(document).process({ text: "服務器錯誤。請與管理員聯繫", width: 320 });
		} else if (e == 'parsererror') {
			$(document).process({ text: "請求的JSON格式錯誤。請與管理員聯繫", width: 320 });
		} else if (e == 'timeout') {
			$(document).process({ text: "與服務器鏈接超時。請與管理員聯繫", width: 320 });
		} else {
			$(document).process({ text: "未知的錯誤。請與管理員聯繫", width: 320 });
		}
	}
});

//後台主頁admin/default.html頁面初始化
$.extend($.reg("rbgame.starcloud.admin"), {
	open: function(id) {
		$("#"+id).click();
		return false;
	},
	init: function () {
		//自動適應窗口大小
		$(window).resize(function () {
			var h = 0;
			if ($(".theader").length == 1) h = $(window).height() - $(".theader").height() - $(".content-right > .nav-tab").height() - 2;
			if ($(".mheader").length == 1) h = $(window).height() - $(".mheader").height() - $(".content-right > .nav-tab").height() - 2;
			$(".nav-tab-page > .active").height(h);
			$(".content-left").height(h + $(".content-right > .nav-tab").height() + 2);
			$(".nav-tree").height($(".content-left").height() - $(".content-left-nav").height());
		}).resize();
		//實例一個TAB
		var tab = $(".nav-tab").navTabs();
		window["tab"] = tab;

		//實例樹
		$(".nav-tree").navTree(function (me) {
			tab.add(me.id, $(me).html(), $(me).attr("href"));
			$(".nav-tab-right-menu").hide();
			if ($.isIE6()) {
				setTimeout(function () {
					tab.active();
				}, 500);
			}
			if ($(".nav-tree").attr('mobile') == '1') $(".nav-tree").parent().find('i').click();
			return false;
		}).find("a").first().click();
		this.tabopen();

		//初始化TAB的右鍵菜單事件
		$(".nav-tab-right-menu a").click(function () {
			switch (this.className) {
				case "close_current": tab.current.find("i").click(); break;
				case "close_all": tab.find("li i").click(); break;
				case "refresh":
					var id = tab.current.attr("id").replace("Tab", "");
					tab.refresh(id, true);
					break;
				case "cancel": break;
			};
			$(".nav-tab-right-menu").hide();
			return false;
		});

		//初始化默認頁
		var url = unescape($.getUrlParam("url").trim());
		var title = unescape($.getUrlParam('title'));
		if (url.length > 1) {
			if ($('a[href="' + url + '"]').length > 0)
				$('a[href="' + url + '"]').first().click();
			else tab.add("DefPage", title, url);
		}

		$(document).bind('keydown', 'esc', function (evt) { tab.close(); return false; });
		$(document).bind('keydown', 'f5', function (evt) { tab.refresh(null, true); return false; });

		$(".content-left-nav i").click(function(){
			if (!$(".content-left:first").hasClass("hide")) {
				$(".content-left:first").addClass("hide");
				$(".content-left-block").removeClass("hide");
				$(".content-right").removeClass("ml170").addClass("ml20");
			}
		});
		$(".content-left-block").click(function(){
			if ($(".content-left:first").hasClass("hide")) {
				$(".content-left:first").removeClass("hide");
				$(".content-left-block").addClass("hide");
				$(".content-right").removeClass("ml20").addClass("ml170");
			}
		});
	},
	delselect: function () {
		//選中時刪除按鈕上顯示選擇記錄數
		$("#chkAll,.chkAll").click(function () {
			var obj = $(this).parent().parent().parent();
			if (this.checked) {
				obj.find("td input[type='checkbox']").attr("checked", true);
				obj.find("tr:gt(0)").addClass("active");
			} else {
				obj.find("td input[type='checkbox']").attr("checked", false);
				obj.find("tr:gt(0)").removeClass("active");
			}
			$(".btnDelete em").text("(" + obj.find(":checked:not('#chkAll,.chkAll')").length + ")");
		});
		$(".gridview tr").click(function () {
			var obj = $(this).find(':checkbox:not("#chkAll,.chkAll")').first().get(0);
			if (!obj) return;
			var chk = obj.checked;
			obj.checked = !chk;
			if (obj.checked) $(this).addClass("active"); else $(this).removeClass("active");
			$(".btnDelete em").text("(" + $(this).parent().find(":checked:not('#chkAll,.chkAll')").length + ")");
		});
		$('.gridview :checkbox:not("#chkAll,.chkAll")').click(function () {
			var chk = this.checked;
			this.checked = !chk;
			$(".btnDelete em").text("(" + $(".gridview :checked:not('#chkAll,.chkAll')").length + ")");
		});
	},
	tabopen: function () {
		//所有帶.tab_open的class都在TAB裡打開
		$(".tab_open").click(function () {
			(top.tab || window["tab"]).add(this.id, $(this).html(), $(this).attr("href"));
			return false;
		});
	},
	iframeinit: function () {
		$(document).ready(function () {
			//ＴＡＢ右鍵菜單
			$(document).click(function () {
				$(top.document).find(".nav-tab-right-menu").hide();
			}).click();
			//f5刷新當前頁
			top.$(document).bind('keydown', 'f5', function (evt) { top.tab.refresh(null, true); return false; });
		});
	},
	initDate: function (params, obj) {
		//初始化時間控件
		params = params || {};
		params["readOnly"] = true;
		(obj || $(".Wdate")).each(function () {
			if ($(this).hasClass("WdateY")) params["dateFmt"] = 'yyyy';
			if ($(this).hasClass("WdateYM")) params["dateFmt"] = 'yyyyMM';
			if ($(this).hasClass("WdateALL")) params["dateFmt"] = 'yyyy-MM-dd HH:mm:ss';
			if ($(this).hasClass("WdateYMDHM")) params["dateFmt"] = 'yyyy-MM-dd HH:mm';

			$(this).focus(function () {
				params["el"] = this.id;
				WdatePicker(params);
			}).click(function () {
				params["el"] = this.id;
				WdatePicker(params);
			}).blur(function () {
				if ($(this).val() == $(this).attr('default')) $(this).val("").removeClass("text_gray");
			});
		});
	},
	initUserControl: function() {
		$("body").userControl({ url: '/admin/sys/user_control__' });
	},
	pageLoad: function (fixTH) {
		//頁面加載
		fixTH = typeof fixTH == 'boolean' ? fixTH : (typeof fixTH == 'number' ? fixTH : false);
		$(document).ready(function () {
			rbgame.starcloud.admin.tabopen();
			$(".table-search :text").textTip();
			if (fixTH) $(".gridview").fixTH(fixTH);
			rbgame.starcloud.admin.initDate();
			$(".input_cbo select").change(function () {
				$(this).prev().text($(this).find("option:selected").text());
			}).change();
			rbgame.starcloud.admin.initUserControl();
		});
	}
});


