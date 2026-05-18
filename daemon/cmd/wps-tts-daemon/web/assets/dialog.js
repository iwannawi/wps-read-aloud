(function () {
  "use strict";

  function fromBase64(value) {
    try {
      return decodeURIComponent(escape(atob(value || "")));
    } catch (_) {
      return "{}";
    }
  }

  function text(value) {
    return value === undefined || value === null ? "" : String(value);
  }

  function appendText(parent, tag, value, className) {
    var el = document.createElement(tag);
    if (className) {
      el.className = className;
    }
    el.textContent = text(value);
    parent.appendChild(el);
    return el;
  }

  function render() {
    var params = new URLSearchParams(window.location.search);
    var payload = {};
    try {
      payload = JSON.parse(fromBase64(params.get("payload")));
    } catch (_) {
      payload = {};
    }

    document.body.className = payload.compact ? "compact" : (payload.about ? "about" : "");
    document.title = payload.title || "文档朗读";
    document.getElementById("dialogTitle").textContent = payload.title || "文档朗读";
    document.getElementById("dialogMessage").textContent = payload.message || "";

    var fields = payload.fields || [];
    if (fields.length) {
      document.getElementById("fieldsSection").hidden = false;
      var dl = document.getElementById("fields");
      fields.forEach(function (item) {
        appendText(dl, "dt", item.label);
        appendText(dl, "dd", item.value);
      });
    }

    var details = payload.details || [];
    if (details.length) {
      document.getElementById("detailsSection").hidden = false;
      var detailsEl = document.getElementById("details");
      details.forEach(function (item) {
        var row = document.createElement("div");
        row.className = "detail-item";
        appendText(row, "span", (item.name || "播放器") + (item.message ? "：" + item.message : ""));
        appendText(row, "span", item.status || "", "detail-status");
        detailsEl.appendChild(row);
      });
    }

    var links = payload.links || [];
    if (links.length) {
      document.getElementById("linksSection").hidden = false;
      var linksEl = document.getElementById("links");
      links.forEach(function (item) {
        var row = document.createElement("div");
        row.className = "link-item";
        appendText(row, "span", item.label);
        var btn = document.createElement("button");
        btn.type = "button";
        btn.className = "link-button";
        btn.textContent = "打开";
        btn.onclick = function () {
          openDoc(item.label, item.url);
        };
        row.appendChild(btn);
        linksEl.appendChild(row);
      });
    }

    document.getElementById("backBtn").onclick = function () {
      showMainView();
    };
    document.getElementById("closeBtn").onclick = function () {
      window.close();
    };
    if (payload.startup) {
      closeWhenPlaybackStarts(payload.startupId || "");
    }
  }

  function closeWhenPlaybackStarts(startupId) {
    var openedAt = Date.now();
    setInterval(function () {
      try {
        if (startupId && localStorage.getItem("wpsReadAloudCloseStartup") === startupId) {
          window.close();
          return;
        }
      } catch (_) {}
      fetch("/read/status", { cache: "no-store" })
        .then(function (response) {
          return response.ok ? response.json() : null;
        })
        .then(function (data) {
          if (!data) {
            return;
          }
          if (data.state === "playing" || data.state === "done" || data.state === "stopped" || data.state === "error") {
            window.close();
          }
          if (data.state === "idle" && Date.now() - openedAt > 1000) {
            window.close();
          }
        })
        .catch(function () {
          // Keep the startup dialog open when the service is still starting.
        });
    }, 300);
  }

  function showMainView() {
    document.getElementById("fieldsSection").hidden = !document.getElementById("fields").children.length;
    document.getElementById("detailsSection").hidden = !document.getElementById("details").children.length;
    document.getElementById("linksSection").hidden = !document.getElementById("links").children.length;
    document.getElementById("docSection").hidden = true;
    document.querySelector("footer").hidden = false;
  }

  function showDocView(title, content) {
    document.getElementById("fieldsSection").hidden = true;
    document.getElementById("detailsSection").hidden = true;
    document.getElementById("linksSection").hidden = true;
    document.getElementById("docSection").hidden = false;
    document.querySelector("footer").hidden = true;
    document.getElementById("docTitle").textContent = title || "说明文件";
    document.getElementById("docContent").textContent = content || "";
  }

  function openDoc(title, url) {
    showDocView(title, "正在加载，请稍候...");
    fetch(url, { cache: "no-store" })
      .then(function (response) {
        if (!response.ok) {
          throw new Error("HTTP " + response.status);
        }
        return response.text();
      })
      .then(function (content) {
        showDocView(title, content);
      })
      .catch(function (error) {
        showDocView(title, "说明文件加载失败：" + (error && error.message ? error.message : String(error)));
      });
  }

  render();
})();
