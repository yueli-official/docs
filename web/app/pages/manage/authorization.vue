<script setup lang="ts">
import { createDocsNotifier } from "~/utils/feedback";
import { ManageEmpty, SkeletonList } from "~/utils/manageComponents";
import { useMinimumLoading } from "@yueli/ui/feedback";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "权限与申请 · 控制台" });

interface Role {
  key: string;
  displayName: string;
  protected: boolean;
  status: string;
  capabilities: string[];
  assignmentSources: string[];
}
interface Policy {
  number: number;
  base: number;
  state: string;
  createdAt: string;
  activatedAt?: string;
}
interface Application {
  id: string;
  subject: string;
  role: string;
  reason: string;
  state: string;
  createdAt: string;
}
interface PolicySnapshot {
  roles: Role[];
  automaticRules: Array<{ key: string; enabled: boolean }>;
}
interface AuthorizationData {
  roles: Role[];
  policies: Policy[];
  applications: Application[];
  snapshot: PolicySnapshot | null;
}
interface Impact {
  addedBindings: number;
  removedBindings: number;
}
type PolicyMutation =
  | { type: "set-role"; role: string; capabilities: string[] }
  | { type: "automatic"; enabled: boolean }
  | {
      type: "create-role";
      key: string;
      displayName: string;
      capabilities: string[];
      sources: string[];
    }
  | { type: "retire-role"; role: string };
interface PendingPolicyChange {
  title: string;
  description: string;
  added: string[];
  removed: string[];
  mutation: PolicyMutation;
}

const capabilityOptions = [
  {
    key: "docs.collection.manage",
    label: "管理文档集",
    description: "创建、修改和删除本站文档集。",
  },
  {
    key: "docs.document.create",
    label: "新建文档",
    description: "在允许的文档集内创建内容。",
  },
  {
    key: "docs.document.read",
    label: "查看管理文档",
    description: "查看自己拥有的管理文档与统计。",
  },
  {
    key: "docs.document.update",
    label: "编辑自己的文档",
    description: "修改自己拥有的文档及层级。",
  },
  {
    key: "docs.document.publish",
    label: "发布自己的文档",
    description: "把自己拥有的文档公开发布。",
  },
  {
    key: "docs.document.archive",
    label: "归档自己的文档",
    description: "归档自己拥有的文档。",
  },
] as const;
const sourceOptions = [
  { key: "application", label: "申请" },
  { key: "invitation", label: "邀请" },
  { key: "direct", label: "管理员直接授予" },
  { key: "group", label: "用户组" },
] as const;

const { call } = useApi();
const { isAdministrator } = useMe();
const toast = createDocsNotifier(useToast());
const mounted = ref(false);
const publishing = ref(false);
const operationError = ref("");
const lastImpact = ref<Impact | null>(null);
onMounted(() => {
  mounted.value = true;
});

const { data, pending, error, refresh } = await useAsyncData(
  "docs-authorization-management",
  async (): Promise<AuthorizationData> => {
    if (!isAdministrator.value) {
      return { roles: [], policies: [], applications: [], snapshot: null };
    }
    const [roles, policies, applications] = await Promise.all([
      call<{ items: Role[] }>("/api/v1/authorization/manage/roles"),
      call<{ items: Policy[] }>("/api/v1/authorization/manage/policies"),
      call<{ items: Application[] }>(
        "/api/v1/authorization/manage/applications?state=pending",
      ),
    ]);
    const active = policies.items.find((policy) => policy.state === "active");
    const snapshot = active
      ? await call<PolicySnapshot>(
          `/api/v1/authorization/manage/policies/${active.number}`,
        )
      : null;
    return {
      roles: roles.items,
      policies: policies.items,
      applications: applications.items,
      snapshot,
    };
  },
  {
    server: false,
    default: () => ({
      roles: [],
      policies: [],
      applications: [],
      snapshot: null,
    }),
  },
);

const showSkeleton = useMinimumLoading(
  computed(() => !mounted.value || pending.value),
);
const roles = computed(() => data.value?.roles ?? []);
const applications = computed(() => data.value?.applications ?? []);
const activePolicy = computed(() =>
  data.value?.policies.find((policy) => policy.state === "active"),
);
const author = computed(() =>
  roles.value.find((role) => role.key === "author"),
);
const customRoles = computed(() =>
  roles.value.filter(
    (role) =>
      !role.protected && role.key !== "author" && role.status !== "retired",
  ),
);
const selectedAuthorCapabilities = ref<string[]>([]);
watch(
  author,
  (role) => {
    selectedAuthorCapabilities.value = [...(role?.capabilities ?? [])];
  },
  { immediate: true },
);
const authorDirty = computed(
  () =>
    normalizedCapabilities(selectedAuthorCapabilities.value).join("\u0000") !==
    normalizedCapabilities(author.value?.capabilities ?? []).join("\u0000"),
);
const autoAuthorEnabled = computed(
  () =>
    data.value?.snapshot?.automaticRules.find(
      (rule) => rule.key === "docs.registration_author",
    )?.enabled ?? false,
);

function normalizedCapabilities(values: readonly string[]) {
  return [...new Set(values)].sort();
}

function capabilityLabel(key: string) {
  return capabilityOptions.find((item) => item.key === key)?.label ?? key;
}

function toggleCapability(
  target: string[],
  capability: string,
  enabled: boolean,
) {
  return enabled
    ? [...new Set([...target, capability])]
    : target.filter((item) => item !== capability);
}

function capabilityDiff(current: string[], next: string[]) {
  const currentSet = new Set(current);
  const nextSet = new Set(next);
  return {
    added: next.filter((item) => !currentSet.has(item)).map(capabilityLabel),
    removed: current.filter((item) => !nextSet.has(item)).map(capabilityLabel),
  };
}

const policyConfirmOpen = ref(false);
const pendingPolicyChange = ref<PendingPolicyChange | null>(null);

function preparePolicyChange(change: PendingPolicyChange) {
  operationError.value = "";
  pendingPolicyChange.value = change;
  policyConfirmOpen.value = true;
}

function prepareAuthorCapabilities() {
  if (!author.value || !authorDirty.value) return;
  const next = normalizedCapabilities(selectedAuthorCapabilities.value);
  const diff = capabilityDiff(author.value.capabilities, next);
  preparePolicyChange({
    title: "发布作者能力策略",
    description:
      "变更只影响本站作者角色，所有者约束仍然生效，作者不会获得他人文档的操作权限。",
    ...diff,
    mutation: { type: "set-role", role: "author", capabilities: next },
  });
}

function prepareAutomaticAuthor(enabled: boolean) {
  if (enabled === autoAuthorEnabled.value) return;
  preparePolicyChange({
    title: enabled ? "开启注册自动作者" : "关闭注册自动作者",
    description: enabled
      ? "以后收到的用户注册事件会自动授予作者角色；已有用户不会被追溯处理。"
      : "以后注册的用户不再自动获得作者角色；已有作者授权不会被撤销。",
    added: enabled ? ["注册自动作者规则"] : [],
    removed: enabled ? [] : ["注册自动作者规则"],
    mutation: { type: "automatic", enabled },
  });
}

const roleEditorOpen = ref(false);
const roleEditorMode = ref<"create" | "edit">("create");
const roleForm = reactive({
  key: "",
  displayName: "",
  capabilities: [] as string[],
  sources: ["application", "invitation", "direct"] as string[],
});
const roleFormError = ref("");

function openCreateRole() {
  roleEditorMode.value = "create";
  Object.assign(roleForm, {
    key: "",
    displayName: "",
    capabilities: [],
    sources: ["application", "invitation", "direct"],
  });
  roleFormError.value = "";
  roleEditorOpen.value = true;
}

function openEditRole(role: Role) {
  roleEditorMode.value = "edit";
  Object.assign(roleForm, {
    key: role.key,
    displayName: role.displayName,
    capabilities: [...role.capabilities],
    sources: [...role.assignmentSources],
  });
  roleFormError.value = "";
  roleEditorOpen.value = true;
}

function toggleRoleCapability(capability: string, enabled: boolean) {
  roleForm.capabilities = toggleCapability(
    roleForm.capabilities,
    capability,
    enabled,
  );
}

function toggleRoleSource(source: string, enabled: boolean) {
  roleForm.sources = toggleCapability(roleForm.sources, source, enabled);
}

function prepareRoleSave() {
  const key = roleForm.key.trim().toLowerCase();
  const displayName = roleForm.displayName.trim();
  if (!/^[a-z][a-z0-9._-]{1,63}$/.test(key)) {
    roleFormError.value =
      "角色标识需以字母开头，仅使用小写字母、数字、点、横线或下划线。";
    return;
  }
  if (!displayName) {
    roleFormError.value = "请填写角色名称。";
    return;
  }
  if (!roleForm.sources.length) {
    roleFormError.value = "至少选择一种角色授予来源。";
    return;
  }
  const capabilities = normalizedCapabilities(roleForm.capabilities);
  if (roleEditorMode.value === "create") {
    preparePolicyChange({
      title: `创建角色“${displayName}”`,
      description:
        "新角色仅能使用 Docs 声明的普通能力，不能获得管理员专属的导入、设置、永久删除或转让能力。",
      added: capabilities.map(capabilityLabel),
      removed: [],
      mutation: {
        type: "create-role",
        key,
        displayName,
        capabilities,
        sources: [...roleForm.sources],
      },
    });
  } else {
    const current = roles.value.find((role) => role.key === key);
    if (!current) return;
    const diff = capabilityDiff(current.capabilities, capabilities);
    if (!diff.added.length && !diff.removed.length) {
      roleFormError.value = "角色能力没有变化。";
      return;
    }
    preparePolicyChange({
      title: `更新角色“${displayName}”`,
      description: "变更会在新策略激活后影响该角色的后续授权判断。",
      ...diff,
      mutation: { type: "set-role", role: key, capabilities },
    });
  }
  roleEditorOpen.value = false;
}

function prepareRoleRetirement(role: Role) {
  preparePolicyChange({
    title: `停用角色“${role.displayName}”`,
    description:
      "角色停用后不能再授予，现有授权会按策略影响预览计算；管理员与内置作者角色不能停用。",
    added: [],
    removed: role.capabilities.map(capabilityLabel),
    mutation: { type: "retire-role", role: role.key },
  });
}

async function mutatePolicy(revision: number, mutation: PolicyMutation) {
  if (mutation.type === "set-role") {
    await call(
      `/api/v1/authorization/manage/policies/${revision}/roles/${mutation.role}/capabilities`,
      {
        method: "PUT",
        body: { capabilities: mutation.capabilities },
      },
    );
    return;
  }
  if (mutation.type === "automatic") {
    await call(
      `/api/v1/authorization/manage/policies/${revision}/automatic/docs.registration_author`,
      { method: "PUT", body: { enabled: mutation.enabled } },
    );
    return;
  }
  if (mutation.type === "create-role") {
    await call(`/api/v1/authorization/manage/policies/${revision}/roles`, {
      method: "POST",
      body: {
        key: mutation.key,
        displayName: mutation.displayName,
        capabilities: mutation.capabilities,
        assignmentSources: mutation.sources,
      },
    });
    return;
  }
  await call(
    `/api/v1/authorization/manage/policies/${revision}/roles/${mutation.role}/retire`,
    { method: "POST" },
  );
}

async function publishPolicyChange() {
  const active = activePolicy.value;
  const change = pendingPolicyChange.value;
  if (!active || !change) return;
  publishing.value = true;
  operationError.value = "";
  try {
    const draft = await call<{ policy: Policy }>(
      "/api/v1/authorization/manage/policies/drafts",
      { method: "POST", body: { expectedActiveRevision: active.number } },
    );
    await mutatePolicy(draft.policy.number, change.mutation);
    const validation = await call<{ valid: boolean; violations: string[] }>(
      `/api/v1/authorization/manage/policies/${draft.policy.number}/validate`,
      { method: "POST" },
    );
    if (!validation.valid) {
      throw new Error(validation.violations.join("；") || "策略校验失败");
    }
    const impact = await call<Impact>(
      `/api/v1/authorization/manage/policies/${draft.policy.number}/preview`,
      { method: "POST" },
    );
    await call(
      `/api/v1/authorization/manage/policies/${draft.policy.number}/activate`,
      {
        method: "POST",
        body: { expectedActiveRevision: active.number },
      },
    );
    lastImpact.value = impact;
    policyConfirmOpen.value = false;
    pendingPolicyChange.value = null;
    await refresh();
    toast.add({
      title: "权限策略已发布",
      description: `新增 ${impact.addedBindings} 项绑定，移除 ${impact.removedBindings} 项绑定。`,
      color: "success",
    });
  } catch (caught) {
    operationError.value =
      caught instanceof Error ? caught.message : "权限策略发布失败，请重试。";
  } finally {
    publishing.value = false;
  }
}

const reviewOpen = ref(false);
const reviewTarget = ref<Application | null>(null);
const reviewDecision = ref<"approve" | "reject">("approve");
const reviewReason = ref("");
const reviewing = ref(false);
const reviewError = ref("");

function openReview(application: Application, decision: "approve" | "reject") {
  reviewTarget.value = application;
  reviewDecision.value = decision;
  reviewReason.value = "";
  reviewError.value = "";
  reviewOpen.value = true;
}

async function submitReview() {
  if (!reviewTarget.value) return;
  if (reviewDecision.value === "reject" && !reviewReason.value.trim()) {
    reviewError.value = "拒绝申请时请填写原因。";
    return;
  }
  reviewing.value = true;
  reviewError.value = "";
  try {
    await call(
      `/api/v1/authorization/manage/applications/${reviewTarget.value.id}/review`,
      {
        method: "POST",
        body: {
          decision: reviewDecision.value,
          reason: reviewReason.value.trim(),
        },
      },
    );
    reviewOpen.value = false;
    await refresh();
    toast.add({
      title: reviewDecision.value === "approve" ? "申请已批准" : "申请已拒绝",
      color: "success",
    });
  } catch (caught) {
    reviewError.value =
      caught instanceof Error ? caught.message : "申请处理失败，请重试。";
  } finally {
    reviewing.value = false;
  }
}

function formatDate(value: string) {
  if (!value) return "未记录";
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return "未记录";
  return new Intl.DateTimeFormat("zh-CN", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
}
</script>

<template>
  <YAdminPage
    id="authorization"
    title="权限与申请"
    icon="i-tabler-shield-lock"
    main-id="manage-main"
    body-class="mx-auto w-full max-w-screen-2xl space-y-4"
  >
    <template #actions>
      <UButton
        v-if="isAdministrator"
        label="新建自定义角色"
        icon="i-tabler-user-shield"
        color="neutral"
        variant="outline"
        @click="openCreateRole"
      />
    </template>

    <div
      v-if="!isAdministrator"
      class="rounded-lg border border-default bg-default p-8"
    >
      <div class="mx-auto max-w-md text-center">
        <span
          class="mx-auto grid size-12 place-items-center rounded-lg bg-warning/10 text-warning"
        >
          <UIcon name="i-tabler-lock" class="size-6" />
        </span>
        <h2 class="mt-4 text-lg font-semibold text-highlighted">
          仅管理员可管理权限
        </h2>
        <p class="mt-2 text-sm leading-6 text-muted">
          作者可以发布内容，但不能修改角色、审批申请或发布权限策略。
        </p>
      </div>
    </div>

    <SkeletonList v-else-if="showSkeleton" :rows="8" />

    <div
      v-else-if="error"
      class="rounded-lg border border-default bg-default p-8"
    >
      <div class="mx-auto max-w-md text-center">
        <span
          class="mx-auto grid size-12 place-items-center rounded-lg bg-error/10 text-error"
        >
          <UIcon name="i-tabler-alert-circle" class="size-6" />
        </span>
        <h2 class="mt-4 text-lg font-semibold text-highlighted">
          权限数据加载失败
        </h2>
        <p class="mt-2 text-sm leading-6 text-muted">
          当前策略、角色或申请队列暂时不可用。
        </p>
        <UButton
          class="mt-5"
          label="重试"
          icon="i-tabler-refresh"
          color="neutral"
          variant="soft"
          @click="() => refresh()"
        />
      </div>
    </div>

    <template v-else>
      <UAlert
        v-if="lastImpact"
        color="success"
        variant="soft"
        icon="i-tabler-circle-check"
        title="权限策略已生效"
        :description="`新增 ${lastImpact.addedBindings} 项绑定，移除 ${lastImpact.removedBindings} 项绑定。`"
      />
      <UAlert
        v-if="operationError"
        color="error"
        variant="soft"
        icon="i-tabler-alert-circle"
        title="权限策略未发布"
        :description="operationError"
      />

      <div class="grid gap-4 xl:grid-cols-[minmax(0,1fr)_22rem]">
        <div class="min-w-0 space-y-4">
          <section
            aria-labelledby="applications-title"
            class="overflow-hidden rounded-xl border border-default bg-default"
          >
            <div
              class="flex items-center justify-between gap-3 border-b border-default px-4 py-3 sm:px-5"
            >
              <div>
                <h2
                  id="applications-title"
                  class="text-sm font-semibold text-highlighted"
                >
                  待审批申请
                </h2>
                <p class="mt-0.5 text-xs text-muted">
                  普通用户申请作者或可申请的自定义角色。
                </p>
              </div>
              <UBadge
                :label="String(applications.length)"
                color="neutral"
                variant="soft"
              />
            </div>
            <div v-if="applications.length" class="divide-y divide-default">
              <div
                v-for="application in applications"
                :key="application.id"
                class="grid gap-3 px-4 py-3 sm:grid-cols-[minmax(0,1fr)_auto] sm:items-center sm:px-5"
              >
                <div class="min-w-0">
                  <div class="flex flex-wrap items-center gap-2">
                    <p class="truncate text-sm font-medium text-highlighted">
                      {{ application.subject }}
                    </p>
                    <UBadge
                      :label="
                        roles.find((role) => role.key === application.role)
                          ?.displayName || application.role
                      "
                      color="neutral"
                      variant="outline"
                      size="sm"
                    />
                  </div>
                  <p class="mt-1 text-sm text-muted">
                    {{ application.reason || "未填写申请理由" }}
                  </p>
                  <time
                    class="mt-1 block text-xs text-dimmed"
                    :datetime="application.createdAt"
                  >
                    {{ formatDate(application.createdAt) }}
                  </time>
                </div>
                <div class="flex gap-2">
                  <UButton
                    label="批准"
                    icon="i-tabler-check"
                    size="sm"
                    @click="openReview(application, 'approve')"
                  />
                  <UButton
                    label="拒绝"
                    icon="i-tabler-x"
                    size="sm"
                    color="neutral"
                    variant="outline"
                    @click="openReview(application, 'reject')"
                  />
                </div>
              </div>
            </div>
            <ManageEmpty v-else icon="i-tabler-inbox" text="暂无待审批申请" />
          </section>

          <section
            aria-labelledby="author-role-title"
            class="rounded-xl border border-default bg-default"
          >
            <div class="border-b border-default px-4 py-3 sm:px-5">
              <h2
                id="author-role-title"
                class="text-sm font-semibold text-highlighted"
              >
                作者能力
              </h2>
              <p class="mt-0.5 text-xs text-muted">
                作者始终受所有者约束，只能管理自己的文档。管理员能力不可在此修改。
              </p>
            </div>
            <div class="grid gap-3 p-4 sm:grid-cols-2 sm:p-5">
              <label
                v-for="option in capabilityOptions"
                :key="option.key"
                class="flex items-start gap-3 rounded-lg border border-default p-3"
              >
                <UCheckbox
                  :model-value="selectedAuthorCapabilities.includes(option.key)"
                  :aria-label="option.label"
                  @update:model-value="
                    selectedAuthorCapabilities = toggleCapability(
                      selectedAuthorCapabilities,
                      option.key,
                      Boolean($event),
                    )
                  "
                />
                <span class="min-w-0">
                  <span class="block text-sm font-medium text-highlighted">
                    {{ option.label }}
                  </span>
                  <span class="mt-0.5 block text-xs leading-5 text-muted">
                    {{ option.description }}
                  </span>
                </span>
              </label>
            </div>
            <div
              class="flex items-center justify-between gap-3 border-t border-default px-4 py-3 sm:px-5"
            >
              <p class="text-xs text-muted">
                修改后先复核差异，再创建、校验并发布新策略。
              </p>
              <UButton
                label="复核并发布"
                icon="i-tabler-shield-check"
                :disabled="!authorDirty"
                @click="prepareAuthorCapabilities"
              />
            </div>
          </section>

          <section
            aria-labelledby="custom-roles-title"
            class="overflow-hidden rounded-xl border border-default bg-default"
          >
            <div
              class="flex items-center justify-between gap-3 border-b border-default px-4 py-3 sm:px-5"
            >
              <div>
                <h2
                  id="custom-roles-title"
                  class="text-sm font-semibold text-highlighted"
                >
                  自定义角色
                </h2>
                <p class="mt-0.5 text-xs text-muted">
                  自定义角色只能组合普通能力，不能获得管理员专属能力。
                </p>
              </div>
              <UButton
                label="新建角色"
                icon="i-tabler-plus"
                color="neutral"
                variant="outline"
                size="sm"
                @click="openCreateRole"
              />
            </div>
            <div v-if="customRoles.length" class="divide-y divide-default">
              <div
                v-for="role in customRoles"
                :key="role.key"
                class="grid gap-3 px-4 py-3 sm:grid-cols-[minmax(0,1fr)_auto] sm:items-center sm:px-5"
              >
                <div class="min-w-0">
                  <div class="flex flex-wrap items-center gap-2">
                    <p class="text-sm font-medium text-highlighted">
                      {{ role.displayName }}
                    </p>
                    <span class="font-mono text-xs text-dimmed">
                      {{ role.key }}
                    </span>
                  </div>
                  <p class="mt-1 text-xs text-muted">
                    {{ role.capabilities.length }} 项能力 ·
                    {{
                      role.assignmentSources
                        .map(
                          (source) =>
                            sourceOptions.find((item) => item.key === source)
                              ?.label || source,
                        )
                        .join("、")
                    }}
                  </p>
                </div>
                <div class="flex gap-2">
                  <UButton
                    label="编辑"
                    icon="i-tabler-pencil"
                    size="sm"
                    color="neutral"
                    variant="outline"
                    @click="openEditRole(role)"
                  />
                  <UButton
                    label="停用"
                    icon="i-tabler-user-off"
                    size="sm"
                    color="error"
                    variant="soft"
                    @click="prepareRoleRetirement(role)"
                  />
                </div>
              </div>
            </div>
            <ManageEmpty
              v-else
              icon="i-tabler-user-shield"
              text="还没有自定义角色"
            />
          </section>
        </div>

        <aside class="min-w-0 space-y-4">
          <section
            aria-labelledby="automatic-role-title"
            class="rounded-xl border border-default bg-default p-4"
          >
            <h2
              id="automatic-role-title"
              class="text-sm font-semibold text-highlighted"
            >
              注册自动授权
            </h2>
            <p class="mt-1 text-xs leading-5 text-muted">
              启用后在用户首次进入本站时幂等补齐作者授权；关闭不会撤销已有授权。
            </p>
            <div
              class="mt-4 flex items-center justify-between gap-3 rounded-lg bg-elevated/40 p-3"
            >
              <span class="text-sm text-highlighted">
                注册用户自动成为作者
              </span>
              <USwitch
                :model-value="autoAuthorEnabled"
                aria-label="注册用户自动成为作者"
                @update:model-value="prepareAutomaticAuthor(Boolean($event))"
              />
            </div>
          </section>

          <section
            aria-labelledby="policy-status-title"
            class="rounded-xl border border-default bg-default p-4"
          >
            <h2
              id="policy-status-title"
              class="text-sm font-semibold text-highlighted"
            >
              当前策略
            </h2>
            <dl class="mt-4 grid gap-3 text-sm">
              <div class="flex items-center justify-between gap-3">
                <dt class="text-muted">活动版本</dt>
                <dd class="font-mono text-highlighted">
                  {{ activePolicy?.number ?? "未初始化" }}
                </dd>
              </div>
              <div class="flex items-center justify-between gap-3">
                <dt class="text-muted">内置角色</dt>
                <dd class="tabular-nums text-highlighted">管理员、作者</dd>
              </div>
              <div class="flex items-center justify-between gap-3">
                <dt class="text-muted">自定义角色</dt>
                <dd class="tabular-nums text-highlighted">
                  {{ customRoles.length }}
                </dd>
              </div>
            </dl>
            <UAlert
              class="mt-4"
              color="neutral"
              variant="soft"
              icon="i-tabler-info-circle"
              title="没有超级管理员"
              description="Identity 只负责登录身份，本站管理员由实例自己的授权数据决定。"
            />
          </section>
        </aside>
      </div>
    </template>

    <UModal
      v-model:open="policyConfirmOpen"
      :title="pendingPolicyChange?.title || '复核权限策略'"
      :description="pendingPolicyChange?.description"
      :ui="{ content: 'sm:max-w-xl', footer: 'justify-end' }"
    >
      <template #body>
        <div v-if="pendingPolicyChange" class="space-y-4">
          <div
            v-if="pendingPolicyChange.added.length"
            class="rounded-lg border border-success/25 bg-success/5 p-3"
          >
            <p class="text-sm font-medium text-success">将新增</p>
            <ul class="mt-2 grid gap-1 text-sm text-default">
              <li
                v-for="item in pendingPolicyChange.added"
                :key="`add-${item}`"
                class="flex items-center gap-2"
              >
                <UIcon name="i-tabler-plus" class="size-4 text-success" />
                {{ item }}
              </li>
            </ul>
          </div>
          <div
            v-if="pendingPolicyChange.removed.length"
            class="rounded-lg border border-warning/25 bg-warning/5 p-3"
          >
            <p class="text-sm font-medium text-warning">将移除</p>
            <ul class="mt-2 grid gap-1 text-sm text-default">
              <li
                v-for="item in pendingPolicyChange.removed"
                :key="`remove-${item}`"
                class="flex items-center gap-2"
              >
                <UIcon name="i-tabler-minus" class="size-4 text-warning" />
                {{ item }}
              </li>
            </ul>
          </div>
          <UAlert
            v-if="operationError"
            color="error"
            variant="soft"
            icon="i-tabler-alert-circle"
            title="无法发布策略"
            :description="operationError"
          />
          <p class="text-xs leading-5 text-muted">
            确认后才会创建服务端 draft，并依次执行校验、影响预览和原子激活。
          </p>
        </div>
      </template>
      <template #footer>
        <UButton
          label="取消"
          color="neutral"
          variant="outline"
          :disabled="publishing"
          @click="
            () => {
              policyConfirmOpen = false;
            }
          "
        />
        <UButton
          label="确认发布"
          icon="i-tabler-shield-check"
          :loading="publishing"
          @click="publishPolicyChange"
        />
      </template>
    </UModal>

    <UModal
      v-model:open="roleEditorOpen"
      :title="roleEditorMode === 'create' ? '新建自定义角色' : '编辑自定义角色'"
      description="自定义角色只组合 Docs 普通能力；管理员专属能力不可下放。"
      scrollable
      :ui="{ content: 'sm:max-w-2xl', footer: 'justify-end' }"
    >
      <template #body>
        <div class="space-y-5">
          <UAlert
            v-if="roleFormError"
            color="error"
            variant="soft"
            icon="i-tabler-alert-circle"
            title="请完善角色信息"
            :description="roleFormError"
          />
          <div class="grid gap-4 sm:grid-cols-2">
            <UFormField
              label="角色标识"
              description="创建后不可修改，用于 API 与审计。"
              required
            >
              <UInput
                v-model="roleForm.key"
                class="w-full font-mono"
                placeholder="reviewer"
                :disabled="roleEditorMode === 'edit'"
              />
            </UFormField>
            <UFormField label="显示名称" required>
              <UInput
                v-model="roleForm.displayName"
                class="w-full"
                placeholder="审校者"
                :disabled="roleEditorMode === 'edit'"
              />
            </UFormField>
          </div>
          <div>
            <h3 class="text-sm font-medium text-highlighted">角色能力</h3>
            <div class="mt-3 grid gap-3 sm:grid-cols-2">
              <label
                v-for="option in capabilityOptions"
                :key="option.key"
                class="flex items-start gap-3 rounded-lg border border-default p-3"
              >
                <UCheckbox
                  :model-value="roleForm.capabilities.includes(option.key)"
                  :aria-label="option.label"
                  @update:model-value="
                    toggleRoleCapability(option.key, Boolean($event))
                  "
                />
                <span class="min-w-0">
                  <span class="block text-sm font-medium text-highlighted">
                    {{ option.label }}
                  </span>
                  <span class="mt-0.5 block text-xs leading-5 text-muted">
                    {{ option.description }}
                  </span>
                </span>
              </label>
            </div>
          </div>
          <div>
            <h3 class="text-sm font-medium text-highlighted">授予来源</h3>
            <div class="mt-3 flex flex-wrap gap-4">
              <UCheckbox
                v-for="source in sourceOptions"
                :key="source.key"
                :model-value="roleForm.sources.includes(source.key)"
                :label="source.label"
                :disabled="roleEditorMode === 'edit'"
                @update:model-value="
                  toggleRoleSource(source.key, Boolean($event))
                "
              />
            </div>
          </div>
        </div>
      </template>
      <template #footer>
        <UButton
          label="取消"
          color="neutral"
          variant="outline"
          @click="
            () => {
              roleEditorOpen = false;
            }
          "
        />
        <UButton
          label="复核变更"
          icon="i-tabler-arrow-right"
          @click="prepareRoleSave"
        />
      </template>
    </UModal>

    <UModal
      v-model:open="reviewOpen"
      :title="reviewDecision === 'approve' ? '批准角色申请' : '拒绝角色申请'"
      :description="
        reviewDecision === 'approve'
          ? '批准后会立即创建对应角色授权。'
          : '拒绝不会创建授权，并会保留处理原因。'
      "
      :ui="{ footer: 'justify-end' }"
    >
      <template #body>
        <div class="space-y-4">
          <UAlert
            v-if="reviewError"
            color="error"
            variant="soft"
            icon="i-tabler-alert-circle"
            title="无法处理申请"
            :description="reviewError"
          />
          <div
            v-if="reviewTarget"
            class="rounded-lg border border-default bg-elevated/35 p-3"
          >
            <p class="font-mono text-xs text-highlighted">
              {{ reviewTarget.subject }}
            </p>
            <p class="mt-1 text-sm text-muted">
              {{ reviewTarget.reason || "未填写申请理由" }}
            </p>
          </div>
          <UFormField
            :label="reviewDecision === 'reject' ? '拒绝原因' : '处理说明'"
            :required="reviewDecision === 'reject'"
          >
            <UTextarea
              v-model="reviewReason"
              class="w-full"
              :rows="3"
              :placeholder="
                reviewDecision === 'reject'
                  ? '请说明拒绝原因'
                  : '可选，记录本次批准依据'
              "
            />
          </UFormField>
        </div>
      </template>
      <template #footer>
        <UButton
          label="取消"
          color="neutral"
          variant="outline"
          :disabled="reviewing"
          @click="
            () => {
              reviewOpen = false;
            }
          "
        />
        <UButton
          :label="reviewDecision === 'approve' ? '确认批准' : '确认拒绝'"
          :color="reviewDecision === 'approve' ? 'primary' : 'error'"
          :loading="reviewing"
          @click="submitReview"
        />
      </template>
    </UModal>
  </YAdminPage>
</template>
