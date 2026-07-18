package hub

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

const adminI18nScriptPath = "/api/beszel/admin-zh-cn.js"

func (h *Hub) registerAdminI18n(se *core.ServeEvent) {
	se.Router.GET(adminI18nScriptPath, func(e *core.RequestEvent) error {
		e.Response.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		e.Response.Header().Set("Cache-Control", "no-cache")
		_, err := e.Response.Write([]byte(adminZhCNScript))
		return err
	})

	se.Router.BindFunc(func(e *core.RequestEvent) error {
		if e.Request.Method != http.MethodGet || (e.Request.URL.Path != "/_/" && e.Request.URL.Path != "/_/index.html") {
			return e.Next()
		}

		originalResponse := e.Response
		recorder := httptest.NewRecorder()
		e.Response = recorder
		err := e.Next()
		e.Response = originalResponse
		if err != nil {
			return err
		}

		for key, values := range recorder.Header() {
			for _, value := range values {
				originalResponse.Header().Add(key, value)
			}
		}
		body := recorder.Body.Bytes()
		if strings.Contains(recorder.Header().Get("Content-Type"), "text/html") {
			body = []byte(injectAdminI18nScript(string(body)))
			originalResponse.Header().Set("Content-Length", strconv.Itoa(len(body)))
		}
		originalResponse.WriteHeader(recorder.Code)
		_, err = io.Copy(originalResponse, bytes.NewReader(body))
		return err
	})
}

func injectAdminI18nScript(html string) string {
	if strings.Contains(html, adminI18nScriptPath) {
		return html
	}
	script := `<script defer src="` + adminI18nScriptPath + `"></script>`
	if strings.Contains(html, "</body>") {
		return strings.Replace(html, "</body>", script+"\n</body>", 1)
	}
	return html + script
}

const adminZhCNScript = `(function () {
  const dictionary = new Map(Object.entries({
    "Collections": "集合",
    "Collection": "集合",
    "Records": "记录",
    "Record": "记录",
    "Settings": "设置",
    "Logs": "日志",
    "Backups": "备份",
    "New collection": "新建集合",
    "Create collection": "创建集合",
    "Edit collection": "编辑集合",
    "Delete collection": "删除集合",
    "New record": "新建记录",
    "Create record": "创建记录",
    "Edit record": "编辑记录",
    "Delete record": "删除记录",
    "View record": "查看记录",
    "Import collections": "导入集合",
    "Export collections": "导出集合",
    "Refresh": "刷新",
    "Reload": "重新加载",
    "Search": "搜索",
    "Filter": "筛选",
    "Filters": "筛选",
    "Sort": "排序",
    "Reset": "重置",
    "Clear": "清空",
    "Save": "保存",
    "Cancel": "取消",
    "Create": "创建",
    "Update": "更新",
    "Delete": "删除",
    "Edit": "编辑",
    "Add": "添加",
    "Remove": "移除",
    "Close": "关闭",
    "Confirm": "确认",
    "Apply": "应用",
    "Submit": "提交",
    "Continue": "继续",
    "Back": "返回",
    "Next": "下一步",
    "Previous": "上一步",
    "More": "更多",
    "Options": "选项",
    "Name": "名称",
    "Type": "类型",
    "System": "系统",
    "Base": "基础",
    "Auth": "认证",
    "View": "视图",
    "Email": "邮箱",
    "Password": "密码",
    "Username": "用户名",
    "Verified": "已验证",
    "Created": "创建时间",
    "Updated": "更新时间",
    "Actions": "操作",
    "Fields": "字段",
    "Indexes": "索引",
    "Rules": "规则",
    "API Rules": "API 规则",
    "Options": "选项",
    "Overview": "概览",
    "Preview": "预览",
    "File": "文件",
    "Files": "文件",
    "Select": "选择",
    "Select all": "全选",
    "Selected": "已选择",
    "No records found": "没有找到记录",
    "No collections found": "没有找到集合",
    "Loading...": "加载中...",
    "Nothing found": "未找到内容",
    "Are you sure?": "确定要继续吗？",
    "This action cannot be undone.": "此操作无法撤销。",
    "Dashboard": "仪表盘",
    "Superusers": "超级用户",
    "Superuser": "超级用户",
    "Log out": "退出登录",
    "Sign in": "登录",
    "Forgot password?": "忘记密码？",
    "Request password reset": "请求重置密码",
    "Confirm password": "确认密码",
    "Old password": "旧密码",
    "New password": "新密码",
    "Collection name": "集合名称",
    "Field name": "字段名称",
    "Field type": "字段类型",
    "Required": "必填",
    "Hidden": "隐藏",
    "Presentable": "展示字段",
    "Unique": "唯一",
    "Min": "最小值",
    "Max": "最大值",
    "Pattern": "正则表达式",
    "Values": "可选值",
    "Relation": "关联",
    "Cascade delete": "级联删除",
    "Expand": "展开",
    "Copy": "复制",
    "Copied": "已复制",
    "API Preview": "API 预览",
    "Documentation": "文档",
    "Logs retention": "日志保留",
    "Application": "应用",
    "Mail settings": "邮件设置",
    "Auth providers": "认证提供方",
    "OAuth2": "OAuth2",
    "Backups": "备份",
    "Cron jobs": "定时任务",
    "Token": "令牌",
    "Tokens": "令牌",
    "Status": "状态",
    "Message": "消息",
    "Level": "级别",
    "Data": "数据",
    "Date": "日期",
    "Time": "时间",
    "Yes": "是",
    "No": "否",
    "true": "是",
    "false": "否"
  }));

  const placeholderDictionary = new Map(Object.entries({
    "Search collections": "搜索集合",
    "Search records": "搜索记录",
    "Filter records": "筛选记录",
    "Type to search": "输入关键词搜索",
    "Enter email": "输入邮箱",
    "Enter password": "输入密码"
  }));

  const titleDictionary = new Map(Object.entries({
    "Refresh": "刷新",
    "Create": "创建",
    "Edit": "编辑",
    "Delete": "删除",
    "Copy": "复制",
    "Close": "关闭"
  }));

  const skipTags = new Set(["SCRIPT", "STYLE", "CODE", "PRE", "TEXTAREA"]);
  let scheduled = false;

  function translateValue(value, map) {
    if (!value) return value;
    const direct = map.get(value.trim());
    return direct || value;
  }

  function translateTextNode(node) {
    const value = node.nodeValue;
    if (!value || !value.trim()) return;
    const leading = value.match(/^\s*/)[0];
    const trailing = value.match(/\s*$/)[0];
    const translated = dictionary.get(value.trim());
    if (translated) node.nodeValue = leading + translated + trailing;
  }

  function translateElement(el) {
    if (!(el instanceof HTMLElement) || skipTags.has(el.tagName)) return;
    for (const attr of ["aria-label", "data-tooltip", "data-title"]) {
      const next = translateValue(el.getAttribute(attr), titleDictionary);
      if (next !== el.getAttribute(attr)) el.setAttribute(attr, next);
    }
    if (el.hasAttribute("title")) {
      const next = translateValue(el.getAttribute("title"), titleDictionary);
      if (next !== el.getAttribute("title")) el.setAttribute("title", next);
    }
    if (el instanceof HTMLInputElement || el instanceof HTMLTextAreaElement) {
      const next = translateValue(el.placeholder, placeholderDictionary);
      if (next !== el.placeholder) el.placeholder = next;
    }
  }

  function walk(root) {
    const walker = document.createTreeWalker(root, NodeFilter.SHOW_ELEMENT | NodeFilter.SHOW_TEXT);
    let node = root;
    while (node) {
      if (node.nodeType === Node.TEXT_NODE) translateTextNode(node);
      else translateElement(node);
      node = walker.nextNode();
    }
  }

  function run() {
    scheduled = false;
    document.documentElement.lang = "zh-CN";
    document.title = document.title.replace("PocketBase", "后台管理");
    walk(document.body);
  }

  function schedule() {
    if (scheduled) return;
    scheduled = true;
    requestAnimationFrame(run);
  }

  new MutationObserver(schedule).observe(document.documentElement, {
    childList: true,
    subtree: true,
    characterData: true,
    attributes: true,
    attributeFilter: ["placeholder", "title", "aria-label", "data-tooltip", "data-title"]
  });

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", schedule, { once: true });
  } else {
    schedule();
  }
})();`
