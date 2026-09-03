const statuses = {
  checked: {
    label: "已预检",
    color: "warning",
    icon: "i-tabler-clipboard-check",
  },
  running: {
    label: "执行中",
    color: "primary",
    icon: "i-tabler-loader-2",
  },
  completed: {
    label: "已完成",
    color: "success",
    icon: "i-tabler-circle-check",
  },
  failed: {
    label: "失败",
    color: "error",
    icon: "i-tabler-alert-circle",
  },
  cancelled: {
    label: "已取消",
    color: "neutral",
    icon: "i-tabler-circle-x",
  },
  rolled_back: {
    label: "已回滚",
    color: "neutral",
    icon: "i-tabler-history",
  },
};

const modes = {
  "create-only": "仅创建",
  upsert: "创建或更新",
  "replace-version": "替换版本",
};

export function importStatusMeta(status) {
  return (
    statuses[String(status || "")] ?? {
      label: status || "未知",
      color: "neutral",
      icon: "i-tabler-file-import",
    }
  );
}

export function importStatusBadgeUI(status) {
  return String(status || "") === "running"
    ? { leadingIcon: "animate-spin motion-reduce:animate-none" }
    : undefined;
}

export function importModeLabel(mode) {
  return modes[String(mode || "")] ?? mode ?? "未记录";
}

export function formatImportDate(value) {
  if (!value) return "未记录";
  const parsed = new Date(value);
  if (Number.isNaN(parsed.getTime())) return "未记录";
  return new Intl.DateTimeFormat("zh-CN", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  }).format(parsed);
}
