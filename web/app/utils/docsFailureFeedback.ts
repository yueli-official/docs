import { resolveFailureFeedback, type ProblemParams } from "@yueli/http-runtime";
import { docsFailurePresentation, type DocsFailureCode } from "../generated/docsFailure";

const messages: Record<DocsFailureCode, string> = {
  "docs.abuse_attempt_replayed": "这次提交已经处理，请勿重复提交。",
  "docs.abuse_unavailable": "内容安全服务暂时不可用。",
  "docs.administrator_grant_protected": "不能移除受保护的管理员授权。",
  "docs.authorization_unavailable": "权限服务暂时不可用。",
  "docs.project_sync_unavailable": "项目同步尚未配置，请联系站点管理员。",
  "docs.challenge_required": "请先完成人机验证。",
  "docs.forbidden": "你没有执行此操作的权限。",
  "docs.initial_administrator_already_claimed": "本站已由另一位用户完成初始化。",
  "docs.import.compression_unsupported": "压缩包使用了当前不支持的压缩方式。",
  "docs.import_blocked": "导入预检发现必须先修复的问题。",
  "docs.invalid_input": "提交内容不符合要求。",
  "docs.not_found": "请求的内容不存在。",
  "docs.rate_limited": "操作过于频繁，请稍后再试。",
  "docs.slug_taken": "这个路径已经被使用。",
  "docs.upstream_failed": "依赖服务暂时不可用。",
};

const recoveries: Record<string, string> = {
  "recovery.retry_later": "请稍后重试。",
  "recovery.fix_import_package": "请根据预检结果修正导入包。",
  "recovery.repack_archive": "请改用受支持的 ZIP 压缩方式重新打包。",
  "recovery.choose_another_slug": "请换一个路径后重试。",
};

const sourceInputMessages: Record<string,string> = {
  github_repository_invalid: '请输入公开 GitHub 仓库地址，例如 https://github.com/owner/project。',
  github_asset_invalid: '附件名称必须是 ZIP 文件名，例如 docs.zip。',
  exclusive_mirror_required: '完整镜像需确认该文档集由这个项目独占维护。',
  default_locale_required: '请选择有效的默认语言。',
  collection_already_has_source: '这个文档集已经绑定项目，请先解除原绑定。',
  source_paused_or_missing: '项目已暂停或不存在，请恢复后再检查。',
  source_id_invalid: '项目绑定地址无效。',
};

function resolveText(code: string, params: ProblemParams) {
  if (code === 'docs.invalid_input' && typeof params?.reason === 'string' && sourceInputMessages[params.reason]) {
    return { message: sourceInputMessages[params.reason]! };
  }
  if (!(code in docsFailurePresentation)) return undefined;
  const typedCode = code as DocsFailureCode;
  const presentation = docsFailurePresentation[typedCode];
  const recoveryKey = "recoveryKey" in presentation ? presentation.recoveryKey : undefined;
  return { message: messages[typedCode], ...(recoveryKey ? { recovery: recoveries[recoveryKey] } : {}) };
}

export function docsFailureFeedback(error: unknown, fallback: string) {
  return resolveFailureFeedback(error, { fallback, resolveText });
}

export function docsFailureMessage(error: unknown, fallback: string) {
  const feedback = docsFailureFeedback(error, fallback);
  return feedback.recovery ? `${feedback.message}${feedback.recovery}` : feedback.message;
}
