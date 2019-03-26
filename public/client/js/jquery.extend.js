/// <reference path="jquery-vsdoc.js" />

if (!Array.indexOf) {
	Array.prototype.indexOf = function (obj, start) {
		/// <summary>
		///	 obj在數組中是否存在，從start位置開始查找
		/// </summary>
		/// <param name="obj" type="Array item type">obj在數組中是否存在</param>
		/// <param name="start" type="int">start位置開始查找</param>
		for (var i = (start || 0); i < this.length; i++) if (this[i] == obj) return i;
		return -1;
	}
};

Date.prototype.format = function (fmt) {
	/// <summary>
	///	 日期轉換為字符串
	/// </summary>
	/// <param name="fmt" type="string">日期format參數 yyyy-MM-dd HH:mm:ss</param>
	var o = { "M+": this.getMonth() + 1, "d+": this.getDate(), "h+": this.getHours() % 12 == 0 ? 12 : this.getHours() % 12, "H+": this.getHours(), "m+": this.getMinutes(), "s+": this.getSeconds(), "q+": Math.floor((this.getMonth() + 3) / 3), "S": this.getMilliseconds() };
	var week = { "0": "\u65e5", "1": "\u4e00", "2": "\u4e8c", "3": "\u4e09", "4": "\u56db", "5": "\u4e94", "6": "\u516d" };
	if (/(y+)/.test(fmt)) { fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length)); };
	if (/(E+)/.test(fmt)) { fmt = fmt.replace(RegExp.$1, ((RegExp.$1.length > 1) ? (RegExp.$1.length > 2 ? "\u661f\u671f" : "\u5468") : "") + week[this.getDay() + ""]); };
	for (var k in o) { if (new RegExp("(" + k + ")").test(fmt)) { fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length))); }; };
	return fmt;
};
String.prototype.isDateTime = function() {
	/// <summary>
	///	 是否時間類型
	/// </summary>
	var r = this.replace(/(^\s*)|(\s*$)/g, "").match(/^(\d{1,4})(-|\/)(\d{1,2})\2(\d{1,2}) (\d{1,2}):(\d{1,2}):(\d{1,2})$/); 
	if (r == null) return false;
	var d = new Date(r[1], r[3] - 1, r[4], r[5], r[6], r[7]);
	return (d.getFullYear() == r[1] && (d.getMonth() + 1) == r[3] && d.getDate() == r[4] && d.getHours() == r[5] && d.getMinutes() == r[6] && d.getSeconds() == r[7]); 
};
String.prototype.isDate = function() {
	/// <summary>
	///	 是否日期類型
	/// </summary>
	var r = this.replace(/(^\s*)|(\s*$)/g, "").match(/^(\d{1,4})(-|\/)(\d{1,2})\2(\d{1,2})$/);
	if (r == null) return false;
	var d = new Date(r[1], r[3] - 1, r[4]);
	return (d.getFullYear() == r[1] && (d.getMonth() + 1) == r[3] && d.getDate() == r[4]);
};
String.prototype.toDateTime = function () {
	/// <summary>
	///	 字符串轉換為日期
	/// </summary>
	if (!this) return null;
	var val = this.replace(/[-]/g, "/");
	if (val.isDate() || val.isDateTime()) return new Date(Date.parse(val));
	var r = this.match(/(\d+)/);
	if (r) return new Date(parseInt(r));
	return new Date(val);
};
String.prototype.format = function () {
	/// <summary>
	///	 格式化字符串"{0} is {1}".format("tom", "cat");
	/// </summary>
	var args = arguments;
	return this.replace(/{(\d+)}/g, function () { return args[arguments[1]]; });
};
String.prototype.sub = function (len, ext) {
	/// <summary>
	///	 截取字符串漢字是兩個字節
	/// </summary>
	/// <param name="len" type="int">長度</param>
	/// <param name="ext" type="string">擴展字符</param>
	ext = ext || '';
	var r = /[^\x00-\xff]/g;
	if (this.replace(r, "mm").length <= len) return this;
	var m = Math.floor(len / 2);
	for (var i = m; i < this.length; i++) { if (this.substr(0, i).replace(r, "mm").length >= len) { return this.substr(0, i) + ext; } }
	return this;
};
String.prototype.encode = function () {
	/// <summary>
	///	 轉義字符串 防止JS代碼執行
	/// </summary>
	if (!this) return '';
	var re = this;
	var q1 = [/\x26/g, /\x3C/g, /\x3E/g, /\x20/g];
	var q2 = ["&amp;", "&lt;", "&gt;", "&nbsp;"];
	for (var i = 0; i < q1.length; i++) re = re.replace(q1[i], q2[i]);
	return re;
};
String.prototype.lengthEx = function () {
	/// <summary>
	///	 中文字符串的長度
	/// </summary>
	return this.replace(/[^\x00-\xff]/g, "--").length;
};
String.prototype.trim = function () {
	/// <summary>
	///	 截取左右空格
	/// </summary>
	return $.trim(this);
};
String.prototype.isNum = function () {
	/// <summary>
	///	 正整數驗證
	/// </summary>
	return /^[1-9][0-9]*(.[0-9]+)?$/.test(this);
};
String.prototype.isInt = function () {
	/// <summary>
	///	 整數驗證
	/// </summary>
	return /^[-\+]?\d+$/.test(this);
};

(function($) {

$.extend({
	zindex: 10,
	windows: 0,
	operatorTemplate: '<div class="{1}">{0}</div><div class="operator">{2}<input id="btnWinCancel" name="btnWinCancel" class="btnWinCancel" type="button" value="取消" /></div>',
	isIE6: function () {
		/// <summary>
		///	 IE6
		/// </summary>
		return $.browser.msie && ($.browser.version == "6.0") && !$.support.style ? true : false;
	},
	reg: function (nsc) {
		/// <summary>
		///	 註冊命名空間.類名或類名
		/// </summary>
		/// <param name="nsc" type="string">命名空間.類名或類名</param>
		var b = nsc.indexOf(".prototype") != -1;
		if (b) nsc = nsc.substr(0, nsc.length - 10);
		var d = nsc.split(".");
		var o = window[d[0]] = d.length == 0 && b ? window[d[0]] || function () { } : window[d[0]] || {};
		for (var i = 1; i < d.length; ++i) o = o[d[i]] = d.length == i + 1 && b ? o[d[i]] || function () { } : o[d[i]] || {};
		if (b) o.prototype = {};
		return b ? o.prototype : o;
	},
	padLeft: function (len, str) {
		/// <summary>
		///	 左邊補指定長度的字符串
		/// </summary>
		/// <param name="len" type="int">長度</param>
		/// <param name="val" type="string">補的字符串</param>
		var str2 = '';
		for (var i = 0; i < len; i++) str2 += str;
		return str2;
	},
	getUrl: function () {
		/// <summary>
		///	 取當前網址URL
		/// </summary>
		return top.location.href;
	},
	getRef: function () {
		/// <summary>
		///	 取當前頁面的來源
		/// </summary>
		return document.referrer;
	},
	getUrlParam: function (paramName) {
		/// <summary>
		///	 取當前頁面GET參數值
		/// </summary>
		/// <param name="paramName" type="string">參數名稱</param>
		var aParams = document.location.search.substr(1).split('&');
		for (i = 0; i < aParams.length; i++) {
			var aParam = aParams[i].split('=');
			if (paramName.toLowerCase() == aParam[0].toLowerCase()) return $.trim(aParam[1]);
		};
		return "";
	},
	cookies: {
		get: function (name) {
			/// <summary>
			///	 取cookies值
			/// </summary>
			/// <param name="name" type="string">cookies名</param>
			var search = name + "=";
			if (document.cookie.length > 0) {
				var offset = document.cookie.indexOf(search);
				if (offset != -1) {
					offset += search.length;
					var end = document.cookie.indexOf(";", offset);
					if (end == -1) end = document.cookie.length;
					return unescape(document.cookie.substring(offset, end));
				} else return "";
			} else return "";
		},
		set: function (name, value, day) {
			/// <summary>
			///	 寫cookies值
			/// </summary>
			/// <param name="name" type="string">cookies名</param>
			/// <param name="value" type="string">cookies值</param>
			/// <param name="day" type="int">天數</param>
			day = day || 365;
			var today = new Date();
			var expires = new Date();
			expires.setTime(today.getTime() + 1000 * 60 * 60 * 24 * day);
			document.cookie = name + "=" + escape(value) + "; path=/; expires=" + expires.toGMTString();
		},
		del: function (name) {
			/// <summary>
			///	 刪除cookies值
			/// </summary>
			/// <param name="name" type="string">cookies名</param>
			var today = new Date();
			var expires = new Date();
			expires.setTime(today.getTime() - 1);
			document.cookie = name + "=0; path=/; expires=" + expires.toGMTString();
		}
	},
	isDay: function (id, cmd) {
		/// <summary>
		///	 每天只執行一次操作
		/// </summary>
		/// <param name="id" type="string">ID</param>
		/// <param name="cmd" type="string">操作</param>
		var isTrue = false;
		var key = "__jquery_day_{0}_{1}__".format(id, cmd);
		var cookieDate = $.cookies.get(key);
		var nowDate = (new Date()).format("yyyy-MM-dd");
		if (cookieDate.length == 10) { if (nowDate != cookieDate) { $.cookies.set(key, nowDate); isTrue = true; }; } else { $.cookies.set(key, nowDate); isTrue = true; };
		return isTrue;
	},
	goUrl: function (url) {
		/// <summary>
		///	 跳轉新頁面
		/// </summary>
		/// <param name="url" type="string">網址</param>
		self.location = url;
	},
	stopPropagation: function (e) {
		if ($.browser.msie) {
			event.cancelBubble = true;
		} else {
			e.stopPropagation();
			e.preventDefault();
		};
	},
	async: function (fun, delay) {
		setTimeout(fun, delay || 0);
	},
	chunk: function (array, process, end){
		var i = 0;
		if (!array || array.length == 0) { return; }
		(function(){
			var item = array.shift();
			i++;
			process(i, item);
			if (array.length > 0)
				setTimeout(arguments.callee, 0);
			else { if (end) end(); }
		})();
	}
});

$.fn.destroy = function () {
	/// <summary>
	///	 真正銷毀對像
	/// </summary>
	var d, el; if (this[0] && this[0].tagName != 'BODY') {
		this.each(function () {
			el = this; $(el).unbind();
			if ($.browser.msie) { d = d || document.createElement('div'); d.appendChild(el); d.innerHTML = ''; d.outerHTML = ''; } else { if (el.parentNode) el.parentNode.removeChild(el); }; 
		});
	}; el = d = null;
	return this;
};
$.fn.log = $.log = function (msg) {
	/// <summary>
	///	 寫日誌console
	/// </summary>
	/// <param name="msg" type="string">日誌消息</param>
	if (console) console.log("%s: %o", msg, this);
	return this;
};
$.fn.drag = function (obj, moveCode, downCode, upCode) {
	/// <summary>
	///	 拖動可控制在某個區域move時執行代碼按下時執行代碼
	/// </summary>
	/// <param name="obj" type="jquery">控制在某個區域</param>
	/// <param name="moveCode" type="function">move時執行代碼</param>
	/// <param name="downCode" type="function">按下時執行代碼</param>
	obj = obj || null; moveCode = moveCode || null; downCode = downCode || null; upCode = upCode || null;
	return this.each(function () {
		var draging = false; var startLeft, startTop; var startX, startY;
		//$(this).css('cursor', 'move');
		$(this).mousedown(function (event) {
			if (downCode) downCode();
			var offset = $(this).offset();
			startLeft = offset.left;
			startTop = offset.top;
			startX = event.clientX;
			startY = event.clientY;
			draging = true;
		}).mousemove(function (event) {
			if (draging == false && moveCode) moveCode(this.id);
			if (draging == false) return;
			var deltaX = event.clientX - startX;
			var deltaY = event.clientY - startY;
			var left = startLeft + deltaX;
			var top = startTop + deltaY;
			if (obj) {
				var m = obj.offset() || $("body").offset();
				if (left < m.left) left = m.left;
				if (top < m.top) top = m.top;
				var width = obj.width() - $(this).width();
				var height = obj.height() - $(this).height();
				if (left > (m.left + width)) left = m.left + width;
				if (top > (m.top + height)) top = m.top + height;
			};
			$(this).css('left', left + 'px').css('top', top + 'px');
		}).mouseup(function (event) {
			if (upCode) moveCode(this.id);
			draging = false;
		});
	});
};
$.fn.valEx = function (val) {
	/// <summary>
	///	 IE6 下拉列表無法賦值可調用$("select").valEx(0)
	/// </summary>
	/// <param name="val" type="string">值</param>
	var obj = this;
	setTimeout(function () {
		obj.val(val);
		obj = null;
	}, 1);
};
$.fn.selectBox = function (data, val, isgroup, title) {
	/// <summary>
	///	 無限級下拉列表
	/// </summary>
	/// <param name="data" type="string">
	///	 數據列表 [{ id = "1", name = "1", node = [{ id = "11", name = "11", node = []}, { id = "12", name = "12", node = []}]}]
	/// </param>
	/// <param name="val" type="string">默認值</param>
	/// <param name="isgroup" type="bool">是否分組顯示true分組顯示只支持二級分類</param>
	var sel = this;
	isgroup = typeof isgroup == 'boolean' ? isgroup : false;
	var html = '<option value="-1">' + (title || "請選擇") + '</option>';
	this.init = function (arr1, level) {
		for (var i = 0; i < arr1.length; i++) {
			if (isgroup && level == 0) html += '<optgroup value="' + arr1[i].id + '" label="' + $.padLeft(level, '　') + arr1[i].name + '">';
			if (level > 0 || !isgroup) html += '<option value="' + arr1[i].id + '">' + $.padLeft(isgroup ? 0 : level, '　') + arr1[i].name + '</option>';
			if (arr1[i].node.length > 0) this.init(arr1[i].node, level + 1);
		};
		if (isgroup && level == 0) html += '<optgroup>';
	};
	this.init(data, 0);
	sel.html(html);
	sel.valEx(val);
	html = sel = init = null;
};
$.fn.asyncSelectBox = function (data, val, isgroup, title, fun) {
	/// <summary>
	///	 無限級下拉列表
	/// </summary>
	/// <param name="data" type="string">
	///	 數據列表 [{ id = "1", name = "1", node = [{ id = "11", name = "11", node = []}, { id = "12", name = "12", node = []}]}]
	/// </param>
	/// <param name="val" type="string">默認值</param>
	/// <param name="isgroup" type="bool">是否分組顯示true分組顯示只支持二級分類</param>
	var sel = this, oldsel = this;
	oldsel.parent().addClass("async_input_cbo");
	isgroup = typeof isgroup == 'boolean' ? isgroup : false;
	sel.append('<option value="-1">-' + (title || "請選擇") + '-</option>');
	var init = function (arr1, level) {
		$.chunk(arr1, function (i, item) {
			if (item) {
				if (isgroup && level == 0) {
					sel.append('<optgroup value="' + item.id + '" label="' + $.padLeft(level, '　') + item.name + '"><optgroup>');
					sel = sel.find("optgroup").last();
				}
				if (level > 0 || !isgroup) sel.append('<option value="' + item.id + '">' + $.padLeft(isgroup ? 0 : level, '　') + item.name + '</option>');
				if (item.node.length > 0) init(item.node, level + 1);
			}
		}, function(){
			oldsel.valEx(val);
			if (fun) fun();
			oldsel.parent().removeClass("async_input_cbo");
		});
	};
	init(data, 0);
};
$.fn.swfPlay = function (width, height, url, params) {
	/// <summary>
	///	 播放FLASH
	/// </summary>
	/// <param name="width" type="int">寬</param>
	/// <param name="height" type="int">高</param>
	/// <param name="url" type="string">flash地址</param>
	/// <param name="params" type="string">參數</param>
	var obj = this;
	params = params || '';
	var html = '<object type="application/x-shockwave-flash" id="swfplayer" name="swfplayer" data="{2}" width="{0}" height="{1}"><param name="movie" value="{2}"/><param name="allowFullScreen" value="true" /><param name="FlashVars" value="{2}" />{3}</object>'.format(width, height, url, params);
	obj.html(html);
	return;
};
$.fn.disable = function (disable) {
	/// <summary>
	///	 對像不可用
	/// </summary>
	this.each(function(){ $(this).get(0).disabled = true; });
	return this;
};
$.fn.enable = function (disable) {
	/// <summary>
	///	 對像可用
	/// </summary>
	this.each(function(){ $(this).get(0).disabled = false; });
	return this;
};
$.fn.checkTextData = function (output, min, max, title, submit, isNVarchar) {
	/// <summary>
	///	 驗證文本框數據
	/// </summary>
	/// <param name="output" type="jquery">輸出JQUERY對像</param>
	/// <param name="min" type="int">最小長</param>
	/// <param name="max" type="int">最大長</param>
	/// <param name="title" type="string">顯示內容</param>
	/// <param name="submit" type="bool">是否提交</param>
	/// <param name="isNVarchar" type="bool">是否NVarchar類型的數據</param>
	var obj = this;
	var len = (isNVarchar) ? obj.val().length : obj.val().lengthEx();
	title = title.split('：')[0].split(':')[0];
	if (len == 0 && min != 0) {
		if (output != "") {
			output.remove("*").notice(title + "不能為空", "error");
			try { if (submit) { obj.focus(); }; } catch (e) { };
		} else {
			$(document).alert({
				id: "sysalert",
				text: title + "不能為空！",
				cancelClick: function () { if (submit) { obj.focus(); } }
			});
		};
		return false;
	} else if (len < min) {
		if (output != "") {
			output.remove("*").notice(title + "不能少於" + min + "位", "error");
			try { if (submit) { obj.focus(); }; } catch (e) { };
		} else {
			$(document).alert({
				id: "sysalert",
				text: title + "不能少於" + min + "位",
				cancelClick: function () { if (submit) { obj.focus(); } }
			});
		};
		return false;
	} else if (len > max) {
		if (output != "") {
			output.remove("*").notice(title + "不能大於" + max + "位", "error");
			try { if (submit) { obj.focus(); }; } catch (e) { };
		} else {
			$(document).alert({
				id: "sysalert",
				text: title + "不能大於" + max + "位",
				cancelClick: function () { if (submit) { obj.focus(); } }
			});
		};
		return false;
	} else {
		if (output != "") output.remove("*").html(this.tableTemplate.format("&nbsp;", "okMsg border0"));
	};
	return true;
};
$.fn.checkSelectData = function (output, title, submit) {
	/// <summary>
	///	 驗證文本框數據
	/// </summary>
	/// <param name="output" type="jquery">輸出JQUERY對像</param>
	/// <param name="title" type="string">顯示內容</param>
	/// <param name="submit" type="bool">是否提交</param>
	var obj = this;
	var value = parseInt(obj.val()) || 0;
	title = title.split('：')[0].split(':')[0];
	if (value <= 0) {
		if (output != "") {
			output.remove("*").notice("請選擇" + title + "！", "error");
			try { if (submit) { obj.focus(); }; } catch (e) { };
		} else {
			$(document).alert({
				id: "sysalert",
				text: "請選擇" + title + "！",
				cancelClick: function () { if (submit) { obj.focus(); } }
			});
		};
		return false;
	}
	return true;
};
$.fn.navMenu = function (fix, align, valign, left, top) {
	/// <summary>
	///	 相對顯示菜單位置
	/// </summary>
	/// <param name="fix" type="jquery">相對位置的JQUERY對像</param>
	/// <param name="align" type="string">左右對齊 left/center/right</param>
	/// <param name="valign" type="string">上下對齊 top/middle/bottom</param>
	/// <param name="left" type="int">左偏移</param>
	/// <param name="top" type="int">上偏移</param>
	var obj = this;
	fix = fix || null; left = left || 0; top = top || 0; align = align || 'left'; valign = valign || 'bottom';
	if (obj.length == 0 || fix.length == 0 || obj.css("display") != "none") return this;

	fix.hover(function(){
		var offset = fix.offset();

		if (align == 'center') obj.css({ "left": offset.left + offset.width() / 2 + left });
		else if (align == 'right') obj.css({ "left": offset.left + fix.width() + left });
		else obj.css({ "left": offset.left + left });

		if (valign == 'top') obj.css({ "top": offset.top + top });
		else if (valign == 'middle') obj.css({ "top": offset.top + fix.height() / 2 + top });
		else obj.css({ "top": offset.top + fix.height() + top });

		setTimeout(function () { obj.show(); }, 50);
	}, function(){
		setTimeout(function () { if (!obj.hasClass("mousemove")) obj.hide(); }, 50);
	});

	obj.hover(function () {
		obj.addClass("mousemove");
	}, function(){
		obj.removeClass("mousemove").hide();
	});

	return this;
};
$.fn.showMenu = function (fix, align, valign, left, top) {
	/// <summary>
	///	 相對顯示菜單位置
	/// </summary>
	/// <param name="fix" type="jquery">相對位置的JQUERY對像</param>
	/// <param name="align" type="string">左右對齊 left/center/right</param>
	/// <param name="valign" type="string">上下對齊 top/middle/bottom</param>
	/// <param name="left" type="int">左偏移</param>
	/// <param name="top" type="int">上偏移</param>
	var obj = this;
	fix = fix || null; left = left || 0; top = top || 0; align = align || 'left'; valign = valign || 'top';
	if (obj.length == 0 || fix.length == 0 || obj.css("display") != "none") return;
	var offset = fix.offset(), docH = $(document).height(), docW = $(document).width(), lastH = offset.top + obj.height(), lastW = offset.left + obj.width();

	if (align == 'center') obj.css({ "left": offset.left + offset.width() / 2 + left });
	else if (align == 'right') obj.css({ "left": offset.left + fix.width() + left });
	else obj.css({ "left": offset.left + left });

	if (valign == 'bottom') obj.css({ "top": offset.top + fix.width() + top });
	else if (valign == 'middle') obj.css({ "top": offset.top + fix.height() / 2 + top });
	else obj.css({ "top": offset.top + top });

	if (lastH > docH && valign == 'top') obj.css({ "top": offset.top - obj.height() });
	if (lastW > docW && align == 'left') obj.css({ "left": offset.left + fix.width() - obj.width() });

	obj.mousemove(function () { $(this).addClass("mousemove"); });
	setTimeout(function () { obj.show(); }, 50);
	offset = null;
};
$.fn.hideMenu = function () {
	/// <summary>
	///	 隱藏菜單
	/// </summary>
	this.hide().removeClass("mousemove");
	return false;
};
$.fn.form = function (active, checkFun) {
	/// <summary>
	///	 表單數據轉成對像
	/// </summary>
	/// <param name="active" type="string">操作</param>
	/// <param name="checkFun" type="function">驗證函數返回true/false</param>
	if (checkFun && !checkFun()) return null;
	var obj = { };
	obj["active"] = active;
	this.find("input:text,input:password,input:hidden,textarea,:radio,:checkbox,select").each(function(){
		var name = this.name, id = this.id, type = this.type;
		if (!this.disabled) {
			if (type == 'radio' && this.checked) {
				if (name) obj[name] = (obj[name] || "") + this.value + ','; else if (id) obj[id] = this.checked;
			} else if (type == 'checkbox' && this.checked) {
				if (name) obj[name] = (obj[name] || "") + this.value + ','; else if (id) obj[id] = this.checked;
			} else if (type != 'radio' && type != 'checkbox') {
				if (name) obj[name] = (obj[name] || "") + this.value + ','; else if (id) obj[id] = this.value;
			}
		}
	});
	for (key in obj) if (obj[key] && obj[key].toString().substr(obj[key].toString().length - 1) == ',') obj[key] = obj[key].toString().substr(0, obj[key].toString().length - 1);
	return obj;
};

$.fn.closeWindow = function (id) {
	/// <summary>
	///	 關閉窗口
	/// </summary>
	id = id || "syswindow";
	$("#" + id + ",#" + id + "_gray").remove();
};
$.fn.window = function (options) {
	/// <summary>
	///	 在容器裡顯示窗口
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id				  窗口ID
	///	 &#10;	title			   標題
	///	 &#10;	html				內容
	///	 &#10;	width			   寬
	///	 &#10;	height			  高
	///	 &#10;	minWidth			最小寬
	///	 &#10;	minHeight		   最小高
	///	 &#10;	left: 0,			左
	///	 &#10;	top: 0,			 上
	///	 &#10;	okClick			 確定單擊事件
	///	 &#10;	cancelClick		 取消單擊事件
	///	 &#10;	okTitle			 確定
	///	 &#10;	cancelTitle		 取消
	///	 &#10;	onActive			活動窗口事件
	///	 &#10;	onMax			   最大化事件
	///	 &#10;	onMin			   最小化事件
	///	 &#10;	onActive			活動窗口事件
	///	 &#10;	onTitleClick		標題栏單擊
	///	 &#10;	isTitle			 是否顯示標題
	///	 &#10;	isClose			 是否顯示關閉按鈕
	///	 &#10;	isMax			   是否顯示最大化按鈕
	///	 &#10;	isMin			   是否顯示最小化按鈕
	///	 &#10;	isDrag			  是否允許拖動
	///	 &#10;	isResize			是否允許改變大小
	///	 &#10;	isFull			  是否全屏
	///	 &#10;	type				默認狀態 0默認 1最小化 2最大化
	///	 &#10;	algin			   對齊
	///	 &#10;	classEx			 附加className樣式
	///	 &#10;	zindex			  層次
	///	 &#10;	icon				小圖標
	///	 &#10;	isMode			  顯示模態窗口
	/// </param>
	var defaultOptions = {
		id: "syswindow",		//窗口ID
		title: "系统提示",	   	//標題
		html: "",			   //內容
		width: 300,			 //寬
		height: 0,			  //高
		minWidth: 538,		  //最小寬
		minHeight: 361,		 //最小高
		left: 0,				//左
		top: 0,				 //上
		okClick: null,		  //確定單擊事件
		cancelClick: null,	  //取消單擊事件
		onActive: null,		 //活動窗口事件
		onMax: null,			//最大化事件
		onMin: null,			//最小化事件
		onTitleClick: null,	 //標題栏單擊
		okTitle: "確定",			//確定
		cancelTitle: "取消",		//取消
		isTitle: true,		  //是否顯示標題
		isClose: true,		  //是否顯示關閉按鈕
		isMax: true,			//是否顯示最大化按鈕
		isMin: true,			//是否顯示最小化按鈕
		isDrag: true,		   //是否允許拖動
		isResize: true,		 //是否允許改變大小
		isFull: false,		  //是否全屏
		type: 0,				//默認狀態 0默認 1最小化 2最大化
		align: null,			//對齊
		classEx: "",			//附加className樣式
		zindex: 4000,		   //層次
		icon: null,			 //小圖標
		isMode: true			//顯示模態窗口
	};
	var opts = $.extend({}, defaultOptions, options);
	var winStatus = 0; //0默認 1最小化 2最大化
	var winLocal = { w: opts.width, h: opts.height, offset: null };

	if ($.zindex > 3000000) $.zindex = 10;
	//if ($.zindex < 3000000 && opts.isMode) $.zindex = $.zindex || 3000001;
	if (opts.align == null) opts.align = "center center";

	this.each(function () {
		var me = $(this);
		var winSize = { w: me.width(), h: me.height() };
		if (opts.html.indexOf('<div class="operator">') == -1 && opts.okClick) opts.html = $.operatorTemplate.replace('取消', opts.cancelTitle || "取消").format(opts.html, 'operatorBox', '<input id="btnWinOK" name="btnWinOK" class="btnWinOK" type="button" value="' + (opts.okTitle ? opts.okTitle : '保存') + '" /> ')
		if (!opts.isTitle) opts.html = opts.html.replace("operator", "operator hide");
		$("#" + opts.id + ",#" + opts.id + "_gray").remove();

		var htmlTemplate = (opts.isMode ? '<div id="{0}_gray" class="syswindow_gray"></div>' : "") + '<div id="{0}" class="syswindow{6}" style="width:{1}px"><div class="msg_top" onselectstart="event.returnValue=false;return false;" onselect="document.selection.empty();"><div class="msg_rt"><a href="#" class="msg_close2"></a></div><div class="msg_lt"></div><div class="msg_ct" style="width:{2}px"><div class="msg_close"><div class="b_close"></div><div class="b_max"></div><div class="b_min"></div></div><div class="title"><img src="" class="b_icon" style="display:none;">{5}</div></div></div><div class="msg_mid"><div class="msg_rm">&nbsp;</div><div class="msg_lm">&nbsp;</div><div class="msg_cm" style="width:{2}px;">{4}</div></div><div class="msg_bott"><div class="msg_rb"></div><div class="msg_lb"></div><div class="msg_cb" style="width:{2}px"></div></div><div class="changesize"></div></div>';
		htmlTemplate = htmlTemplate.format(opts.id, opts.width, opts.width, opts.height, opts.html, opts.title, opts.classEx);
		if (this == window || this == document) $("body").append(htmlTemplate); else me.append(htmlTemplate);
		$("html").addClass("syshtml");
		var win = $("#" + opts.id);
		win.width(opts.width + win.find(".msg_lt").width() + win.find(".msg_rt").width());
		if (opts.icon) { win.find(".b_icon").attr("src", opts.icon).show(); };

		if (!opts.isTitle) {
			win.find(".msg_top,.msg_rt,.msg_lt,.msg_ct").addClass("msgheight");
			win.find(".msg_ct *:not(.b_close)").remove();
			win.find(".b_close").hide();
			win.find(".msg_close2").show();
		};
		if (!opts.isClose) win.find(".msg_close2").hide();

		if (opts.isMode) {
			win.find(".b_max,.b_min").remove();
			win.find(".changesize").css("cursor", "default");
		} else {
			win.find(".msg_lt,.msg_rt,.msg_rm,.msg_lm,.msg_rb,.msg_lb,.msg_cb,.changesize").mousedown(function (e) { $.stopPropagation(e); if (opts.onActive) opts.onActive(); });
		}

		if (!opts.isMax) win.find(".b_max").remove(); else win.find(".msg_ct").dblclick(function () { win.find(".b_max").click(); });
		if (!opts.isMin) win.find(".b_min").remove();
		if (!opts.isResize) win.find(".msg_lt,.msg_rt,.msg_rm,.msg_lm,.msg_rb,.msg_lb,.msg_cb").css("cursor", "default");
		if (opts.isTitle && opts.onTitleClick != null) win.find(".msg_ct").click(function () { opts.onTitleClick(); });

		win.find(".b_close").click(function () {
			$('#loading' + options.id).remove();
			if (opts.cancelClick) opts.cancelClick();
			$("html").removeClass("syshtml");
			$("#" + opts.id + ",#" + opts.id + "_gray").remove();
			$.windows--;
			htmlTemplate = null;
		}).hover(function () {
			win.find(".b_close").addClass("b_close_active");
		}, function () {
			win.find(".b_close").removeClass("b_close_active");
		});

		win.find(".b_max").click(function () {
			if (winStatus == 0) {
				winSize = { w: me.width(), h: me.height() };
				winLocal.offset = win.offset();
				if (opts.onMax) opts.onMax();
				var offset = win.offset();
				var ext = offset.left * 2;
				win.width(winSize.w - ext).height(winSize.h);
				win.find(".msg_cm").width(winSize.w - win.find(".msg_lt").width() - win.find(".msg_rt").width() - ext).height(winSize.h - win.find(".msg_top").height() - win.find(".msg_bott").height());
				win.find(".msg_ct,.msg_cb").width(win.find(".msg_cm").width());
				win.find(".msg_rm,.msg_lm").height(win.find(".msg_cm").height());
				win.unbind("mousedown").unbind("mousemove").unbind("mouseup").find(".title,.msg_ct").css("cursor", "default");
				win.find(".msg_lt,.msg_rt,.msg_rm,.msg_lm,.msg_rb,.msg_lb,.msg_cb").css("cursor", "default");
				win.find(".b_max").addClass("b_restore");
				winStatus = 2;
				opts.isFull = true;
				win.find(".b_restore").hover(function () {
					win.find(".b_restore").addClass("b_restore_active");
				}, function () {
					win.find(".b_restore").removeClass("b_restore_active");
				});
				win.find(".msg_bott").hide();
			} else if (winStatus == 2) {
				if (winLocal.offset.left == -1000) winLocal.offset.left = 100 + ($.windows - 2) * 30;
				if (winLocal.offset.top == -1000) winLocal.offset.top = 50 + ($.windows - 2) * 30;
				win.width(winLocal.w).height(winLocal.h).css({ left: winLocal.offset.left, top: winLocal.offset.top });
				win.find(".msg_cm").width(winLocal.w - ($("#" + opts.id + " .msg_lt").width() * 2)).height(winLocal.h - $("#" + opts.id + " .msg_top").height() - $("#" + opts.id + " .msg_bott").height());
				win.find(".msg_ct,.msg_cb").width(win.find(".msg_cm").width());
				win.find(".msg_rm,.msg_lm").height(win.find(".msg_cm").height());
				win.unbind().find(".title,.msg_ct").css("cursor", "move");
				if (opts.isDrag) win.drag(me, null, function () {
					$.zindex = $.zindex + 1;
					win.css("z-index", $.zindex);
				});
				win.find(".b_max").removeClass("b_restore").removeClass("b_restore_active");
				winStatus = 0;
				opts.isFull = false;
				win.find(".msg_bott").show();
			}
			return false;
		}).hover(function () {
			win.find(".b_max").addClass("b_max_active");
		}, function () {
			win.find(".b_max").removeClass("b_max_active");
		});

		win.find(".b_min").click(function () {
			$("#" + opts.id + ",#" + opts.id + "_gray").hide();
			if (opts.onMin) opts.onMin();
			return false;
		}).hover(function () {
			win.find(".b_min").addClass("b_min_active");
		}, function () {
			win.find(".b_min").removeClass("b_min_active");
		});

		win.find(".msg_close2").click(function () {
			$('#loading' + options.id).remove();
			if (opts.cancelClick) opts.cancelClick();
			$("html").removeClass("syshtml");
			$("#" + opts.id + ",#" + opts.id + "_gray").remove();
			$.windows--;
			htmlTemplate = null;
		});
		win.css('position', 'fixed');
		$("#iframe_" + opts.id).css({ 'width': '100%', 'height': '100%' });

		if (opts.height > 0) {
			win.find(".icon").css({ 'height': (opts.height - 76) + 'px' }); //, 'min-height': (opts.height - 76) + 'px'
			if (win.find(".icon").length == 0) win.find(".msg_cm").css({ 'height': opts.height + 'px' }); //, 'min-height': opts.height + 'px'
		};

		win.find(".operator #btnWinCancel").click(function () { win.find(".b_close").click(); }).focus();
		win.find(".operator #btnWinOK").click(function () {
			if (opts.okClick) { $("html").removeClass("syshtml"); if (opts.okClick() == false) return false; };
			win.remove();
			$("#" + opts.id + "_gray").remove();
			htmlTemplate = null;
		});

		win.find(".msg_lm,.msg_rm").height(win.find(".msg_cm").height());
		if (opts.isClose) $(document).keyup(function (e) { e = window.event || e; if (e.keyCode == 27) { win.find(".b_close").click(); $(document).unbind("keyup"); } });

		var initPos = function (isopen) {
			var winSize = { w: me.width(), h: me.height() };
			if (opts.isFull) {
				var offset = win.offset();
				var ext = offset.left * 2;
				win.width(winSize.w - ext).height(winSize.h);
				win.find(".msg_cm").width(winSize.w - win.find(".msg_lt").width() - win.find(".msg_rt").width() - ext).height(winSize.h - win.find(".msg_top").height() - win.find(".msg_bott").height());
				win.find(".msg_ct,.msg_cb").width(win.find(".msg_cm").width());
				win.find(".msg_rm,.msg_lm").height(win.find(".msg_cm").height());

				return false;
			};

			if (opts.align == "center center") {
				win.css('left', ((winSize.w - opts.width + 26) / 2 - 20) + 'px');
				win.css('top', ((winSize.h - (win.find(".msg_cm").height() + 10)) / 2 - 30) + 'px');
				if ($.isIE6()) {
					win.css('top', (document.documentElement.scrollTop) + 'px');
					win.find(".msg_cm").width(win.find(".msg_cm").width() - 2);
					win.css('position', 'absolute');
					$("#" + opts.id + "_gray").css({ 'position': 'absolute', 'width': $("body").width(), 'height': $("body").height() });
					win.css('top', (($("body").height() - $("#" + opts.id + " .msg_cm").height()) / 2 - 30) + 'px');
					win.css('left', ((winSize.w - opts.width + 26) / 2 - 100) + 'px');
					//if (win.offset().top - document.documentElement.scrollTop == 0) 
					//win.css('top', (document.documentElement.scrollTop + ((winSize.h - ($("#" + opts.id + " .msg_cm").height() + 10)) / 2 - 30)) + 'px');
				};
			} else if (opts.align == "left top") {
				win.css('left', '0px');
				win.css('top', '0px');
			} else if (opts.align == "left center") {
				win.css('left', '0px');
				win.css('top', ((winSize.h - (win.height() + 10)) / 2 - 30) + 'px');
			} else if (opts.align == "left bottom") {
				win.css('left', '0px');
				win.css('top', (winSize.h - win.height()) + 'px');
			} else if (opts.align == "center top") {
				win.css('left', ((winSize.w - win.width()) / 2) + 'px');
				win.css('top', '0px');
			} else if (opts.align == "center bottom") {
				win.css('left', ((winSize.w - win.width()) / 2) + 'px');
				win.css('top', (winSize.h - win.height()) + 'px');
			} else if (opts.align == "right top") {
				win.css('left', (winSize.w - win.width()) + 'px');
				win.css('top', '0px');
			} else if (opts.align == "right center") {
				win.css('left', (winSize.w - win.width()) + 'px');
				win.css('top', ((winSize.h - (win.height())) / 2) + 'px');
			} else if (opts.align == "right bottom") {
				win.css('left', (winSize.w - win.width()) + 'px');
				win.css('top', (winSize.h - win.height()) + 'px');
			} else {
				if (isopen) {
					win.css('left', (100 + ($.windows - 1) * 30) + 'px');
					win.css('top', (50 + ($.windows - 1) * 30) + 'px');
				};
			};
			if (opts.isDrag) win.drag(me, null, function () {
				$.zindex = $.zindex + 1;
				win.css("z-index", $.zindex);
				opts.align = "";
				opts.left = win.offset().left;
				opts.top = win.offset().top;
			});
		};
		$(window).resize(function () { initPos(false); });
		initPos(true);
		if (opts.isFull) win.find(".b_max").click();

		if (win.find("iframe").length > 0) {
			win.find(".msg_cm").tooltip({
				id: 'loading' + opts.id,
				text: '正在加載數據...',
				type: 'process',
				width: 100,
				left: -140,
				top: 10,
				styleEx: "z-index:4000000"
			});
			//win.find(".msg_cm").popupLoading();
		}
		$.zindex = $.zindex + 10;
		$("#" + opts.id + "_gray").css("z-index", $.zindex);
		$.zindex = $.zindex + 1;
		win.css("z-index", $.zindex);

		var _x = 0, _y = 0;
		var ResizeStart = function (e) {
			_x = (win[0].offsetLeft || 0);
			_y = (win[0].offsetTop || 0);
			if ($.browser.msie) {
				win[0].attachEvent("onlosecapture", ResizeStop);
				win[0].setCapture();
			} else {
				e.preventDefault();
				$(window).blur(ResizeStop);
			}
			$(document).mousemove(ResizeMove);
			$(document).mouseup(ResizeStop);
		};
		var ResizeMove = function (e) {
			window.getSelection ? window.getSelection().removeAllRanges() : document.selection.empty();
			var i_x = e.clientX - _x, i_y = e.clientY - _y;
			win.width(Math.max(i_x, opts.minWidth));
			win.height(Math.max(i_y, opts.minHeight));

			win.find(".msg_cm").width(Math.max(win.width() - win.find(".msg_lt").width() - win.find(".msg_rt").width(), opts.minWidth - win.find(".msg_lt").width() - win.find(".msg_rt").width()))
				.height(Math.max(win.height() - win.find(".msg_top").height() - win.find(".msg_bott").height(), opts.minHeight - win.find(".msg_top").height() - win.find(".msg_bott").height()));
			win.find(".msg_ct,.msg_cb").width(win.find(".msg_cm").width());
			win.find(".msg_rm,.msg_lm").height(win.find(".msg_cm").height());
		};
		var ResizeStop = function () {
			$(document).unbind('mousemove');
			$(document).unbind('mouseup');
			if ($.browser.msie) {
				win[0].detachEvent("onlosecapture", ResizeStop);
				win[0].releaseCapture();
			} else {
				$(window).unbind("blur");
			};
		};
		win.find(".changesize").mousedown(ResizeStart);

		$.windows++;
		htmlTemplate = null;
	});
};
$.fn.alert = function (options) {
	/// <summary>
	///	 在容器里顯示消息窗口
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id				  窗口ID
	///	 &#10;	title			   標題
	///	 &#10;	html				內容
	///	 &#10;	width			   寬
	///	 &#10;	height			  高
	///	 &#10;	okClick			 確定單擊
	///	 &#10;	cancelClick		 取消單擊
	///	 &#10;	okTitle			 確定
	///	 &#10;	cancelTitle		 取消
	///	 &#10;	isTitle			 顯示標題
	///	 &#10;	isClose			 顯示關閉按鈕
	///	 &#10;	isDrag			  允許拖動
	///	 &#10;	isResize			允許改變大小
	///	 &#10;	type				默認狀態 0默認 1最小化 2最大化
	///	 &#10;	algin			   對齊
	///	 &#10;	classEx			 附加類樣式
	///	 &#10;	zindex			  層次
	///	 &#10;	icon				小圖標
	/// </param>
	if (!options) options = {};
	var html = $.operatorTemplate.replace("取消", options.cancelTitle || "關閉").format('<div class="left alert">&nbsp;</div><div class="text">{0}</div>'.format(options.text || "无提示"), 'icon', '');
	if (!options.html) options.html = html;
	if (!options.align) options.align = "center center";

	this.window(options);
};
$.fn.confirm = function (options) {
	/// <summary>
	///	 在容器里顯示确认窗口
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id				  窗口ID
	///	 &#10;	title			   標題
	///	 &#10;	html				內容
	///	 &#10;	width			   寬
	///	 &#10;	height			  高
	///	 &#10;	okClick			 確定單擊
	///	 &#10;	cancelClick		 取消單擊
	///	 &#10;	okTitle			 確定
	///	 &#10;	cancelTitle		 取消
	///	 &#10;	isTitle			 顯示標題
	///	 &#10;	isClose			 顯示關閉按鈕
	///	 &#10;	isDrag			  允許拖動
	///	 &#10;	isResize			允許改變大小
	///	 &#10;	type				默認狀態 0默認 1最小化 2最大化
	///	 &#10;	algin			   對齊
	///	 &#10;	classEx			 附加類樣式
	///	 &#10;	zindex			  層次
	///	 &#10;	icon				小圖標
	/// </param>
	if (!options) options = { };
	var html = $.operatorTemplate.replace("取消", options.cancelTitle || "取消").format('<div class="left confirm">&nbsp;</div><div class="text">{0}</div>'.format(options.text || "无提示"), 'icon', '<input id="btnWinOK" name="btnWinOK" class="btnWinOK" type="button" value="' + (options.okTitle || "確定") + '" /> ');
	if (!options.html) options.html = html;
	if (!options.align) options.align = "center center";
	this.window(options);
};
$.fn.iframe = function (options) {
	/// <summary>
	///	 在容器里顯示iframe窗口
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id				  窗口ID
	///	 &#10;	title			   標題
	///	 &#10;	src				 URL
	///	 &#10;	width			   寬
	///	 &#10;	height			  高
	///	 &#10;	okClick			 確定單擊
	///	 &#10;	cancelClick		 取消單擊
	///	 &#10;	okTitle			 確定
	///	 &#10;	cancelTitle		 取消
	///	 &#10;	isTitle			 顯示標題
	///	 &#10;	isClose			 顯示關閉按鈕
	///	 &#10;	isDrag			  允許拖動
	///	 &#10;	isResize			允許改變大小
	///	 &#10;	type				默認狀態 0默認 1最小化 2最大化
	///	 &#10;	algin			   對齊
	///	 &#10;	classEx			 附加類樣式
	///	 &#10;	zindex			  層次
	///	 &#10;	icon				小圖標
	/// </param>
	if (!options) options = {};
	if (!options.id) options.id = "syswindow";
	var html = '<iframe id="iframe_{0}" src="{1}" frameborder="0" style="border:none;display:none;"></iframe>'.format(options.id, options.src);
	if (!options.html) options.html = html;
	if (options.align==null) options.align = "center center";
	this.window(options);
	$("#iframe_" + options.id).load(function () {
		$(this).show();
		$('#loading' + options.id).remove();
	});
};
$.fn.popupProcess = function (options) {
	/// <summary>
	///	 在容器里顯示处理窗口
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id				  窗口ID
	///	 &#10;	title			   標題
	///	 &#10;	html				內容
	///	 &#10;	width			   寬
	///	 &#10;	height			  高
	///	 &#10;	okClick			 確定單擊
	///	 &#10;	cancelClick		 取消單擊
	///	 &#10;	okTitle			 確定
	///	 &#10;	cancelTitle		 取消
	///	 &#10;	isTitle			 顯示標題
	///	 &#10;	isClose			 顯示關閉按鈕
	///	 &#10;	isDrag			  允許拖動
	///	 &#10;	isResize			允許改變大小
	///	 &#10;	type				默認狀態 0默認 1最小化 2最大化
	///	 &#10;	algin			   對齊
	///	 &#10;	classEx			 附加類樣式
	///	 &#10;	zindex			  層次
	///	 &#10;	icon				小圖標
	/// </param>
	if (!options) options = { };
	var html = '<div class="processMsg" style="border:0px;background-color:#fff;margin:5px; 0px;">{0}</div>'.format(options.text || "正在處理數據，請等待...");
	if (!options.html) options.html = html;
	if (!options.width) options.width = 200;
	if (!options.align) options.align = "center center";
	options.isClose = false;
	options.isTitle = false;
	this.window(options);
};
$.fn.popupLoading = function (options) {
	/// <summary>
	///	 在容器里顯示加载窗口
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id				  窗口ID
	///	 &#10;	title			   標題
	///	 &#10;	html				內容
	///	 &#10;	width			   寬
	///	 &#10;	height			  高
	///	 &#10;	okClick			 確定單擊
	///	 &#10;	cancelClick		 取消單擊
	///	 &#10;	okTitle			 確定
	///	 &#10;	cancelTitle		 取消
	///	 &#10;	isTitle			 顯示標題
	///	 &#10;	isClose			 顯示關閉按鈕
	///	 &#10;	isDrag			  允許拖動
	///	 &#10;	isResize			允許改變大小
	///	 &#10;	type				默認狀態 0默認 1最小化 2最大化
	///	 &#10;	algin			   對齊
	///	 &#10;	classEx			 附加類樣式
	///	 &#10;	zindex			  層次
	///	 &#10;	icon				小圖標
	/// </param>
	if (!options) options = { };
	if (!options.text) options.text = "正在加載數據...";
	if (!options.width) options.width = 150;
	options.isClose = false;
	options.isTitle = false;
	this.popupProcess(options);
};

$.fn.notice = function (options) {
	/// <summary>
	///	 在容器里顯示消息
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id:		 窗口ID
	///	 &#10;	text:	   內容 為null時是隱藏
	///	 &#10;	width:	  寬
	///	 &#10;	type:	   默認狀態 ok,warning,error,process
	///	 &#10;	interval:   自動關閉(秒)
	///	 &#10;	classEx:	附加樣式類
	///	 &#10;	styleEx:	附加樣式
	///	 &#10;	isTop:	  true 加載到最上面false 加載到最下面
	/// </param>

	var defaultOptions = {
		id: "sysnotice",		//窗口ID
		text: null,			 //內容
		width: 0,			   //寬
		type: "ok",			 //默認狀態 ok,warning,error,process
		interval: 0,			//自动關閉
		classEx: "",			//附加樣式類
		styleEx: "",			//附加樣式
		isTop: false			//true 加載到最上面false 加載到最下面
	};
	var opts = $.extend({}, defaultOptions, options);

	this.each(function () {
		var me = $(this);
		if (!opts.text) { me.children().remove(); return false; };

		var html = '<div id="{3}" class="{0}Msg {4}" style="float:left;{1}">{2}</div>'.format(opts.type, opts.styleEx + (opts.width > 0 ? "width:{0}px;".format(opts.width) : ""), opts.text, opts.id, opts.classEx);
		if (opts.interval == 0) me.children().remove();

		if (opts.isTop) me.prepend(html); else me.append(html);
		if (opts.interval > 0) { setTimeout(function () { $("#" + opts.id).remove(); }, opts.interval * 1000); };
		html = me = null;
	});
};
$.fn.tooltip = function (options) {
	/// <summary>
	///	 在容器里顯示消息
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id:		 窗口ID
	///	 &#10;	text:	   內容 為null時是隱藏
	///	 &#10;	width:	  寬
	///	 &#10;	type:	   默認狀態 ok,warning,error,process
	///	 &#10;	classEx:	附加樣式類
	///	 &#10;	styleEx:	附加樣式
	///	 &#10;	left:	   左偏移
	///	 &#10;	top:		上偏移
	/// </param>

	var defaultOptions = {
		id: "systooltip",		//窗口ID
		text: null,			 //內容
		width: 0,			   //寬
		type: "ok",			 //默認狀態 ok,warning,error,process
		classEx: "",			//附加樣式類
		styleEx: "",			//附加樣式
		left: 0,				//左偏移
		top: 0				  //上偏移
	};
	var opts = $.extend({}, defaultOptions, options);

	this.each(function () {
		var me = $(this);
		$("#" + opts.id).remove();
		if (!me || !opts.text) return false;

		var html = '<div class="{5}Msg {4}" id="{0}" style="position:absolute;z-index:2000;display:none;width:{1};{3}">{2}</div>'
			.format(opts.id, opts.width == 0 ? "100%" : (opts.width + "px"), opts.text, opts.styleEx, opts.classEx, opts.type);
		$("body").prepend(html);
		$("#" + opts.id).showMenu(me, "right", "top", opts.left, opts.top);
		html = me = null;
	});
};
$.fn.popup = function (options) {
	/// <summary>
	///	 在容器里顯示飘浮窗口
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id:		 窗口ID
	///	 &#10;	text:	   內容 為null時是隱藏
	///	 &#10;	width:	  寬
	///	 &#10;	type:	   默認狀態 ok,warning,error,process
	///	 &#10;	interval:   自動關閉(秒)
	///	 &#10;	classEx:	附加樣式類
	///	 &#10;	styleEx:	附加樣式
	///	 &#10;	isTop:	  true 加載到最上面false 加載到最下面
	/// </param>
	if (!options) options = { };
	if (!options.align) options.align = "right bottom";
	this.window(options);
};
$.fn.process = function (options) {
	/// <summary>
	///	 在容器里顯示处理消息
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id:		 窗口ID
	///	 &#10;	text:	   內容 為null時是隱藏
	///	 &#10;	width:	  寬
	///	 &#10;	type:	   默認狀態 ok,warning,error,process
	///	 &#10;	interval:   自動關閉(秒)
	///	 &#10;	classEx:	附加樣式類
	///	 &#10;	styleEx:	附加樣式
	///	 &#10;	isTop:	  true 加載到最上面false 加載到最下面
	/// </param>
	if (!options) options = {};
	if (!options.text) options.text = "正在處理數據，請等待...";
	options.type = "process";
	$(this).notice(options);
};
$.fn.loading = function (options) {
	/// <summary>
	///	 在容器里加载消息
	/// </summary>
	/// <param name="options" type="object">
	///	 参数列表：
	///	 &#10;	id:		 窗口ID
	///	 &#10;	text:	   內容 為null時是隱藏
	///	 &#10;	width:	  寬
	///	 &#10;	type:	   默認狀態 ok,warning,error,process
	///	 &#10;	interval:   自動關閉(秒)
	///	 &#10;	classEx:	附加樣式類
	///	 &#10;	styleEx:	附加樣式
	///	 &#10;	isTop:	  true 加載到最上面false 加載到最下面
	/// </param>
	if (!options) options = {};
	if (!options.text) options.text = "正在加載數據...";
	options.type = "process";
	$(this).notice(options);
};

$.fn.textTip = function () {
	/// <summary>
	///	 text textarea顯示默認值
	/// </summary>
	return this.each(function () {
		var me = $(this), defV = $.trim(me.attr("default"));

		me.focus(function () {
			var val = $.trim(me.val());
			if (defV == val || val.length == 0) {
				me.val('');
				me.removeClass("text_gray");
			}
		}).blur(function () {
			init();
		});

		function init() {
			var val = $.trim(me.val());
			if (defV != val && val.length > 0) me.removeClass("text_gray");
			if (defV == val || val.length == 0) me.addClass("text_gray");
			if (val.length == 0) me.val(defV);
		}
		init();
	});
}
$.fn.btnSubmit = function(form) {
	/// <summary>
	///	 在text textarea內按回車或ctrl+回車時調用
	/// </summary>
	/// <param name="form" type="jquery">form內的所有text textarea</param>
	var me = this;
	form.find("input:text,input:password,:radio,:checkbox,select").keypress(function (e) {
		e = window.event || e;
		var code = e.keyCode || e.which;
		if (code == 13) me.click();
	});
	form.find("textarea").keypress(function (e) {
		e = window.event || e;
		var code = e.keyCode || e.which;
		if (e.ctrlKey && code == 13) me.click();
	});
	return this;
}
$.fn.textPassword = function () {
	/// <summary>
	///	 text 輸入密碼時顯示●
	/// </summary>
	return this.each(function () {
		var ome = $(this),
			oid = this.name || this.id,
			nid = oid + "Temp",
			html = ome.parent().html().replace(oid, nid);
		ome.hide().before(html);
		var nme = $("#" + nid);
		nme.keyup(function () {
			var val = this.value, len = val.length;
			ome.val((ome.val() + val).replace(/●/g, '').substring(0, len));
			nme.val($.padLeft(len, "●"));
		});
	});
}

}(jQuery));
