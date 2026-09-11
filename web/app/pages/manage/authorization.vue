<script setup lang="ts">
import { assetMediaUrl } from "@yueli/asset-nuxt/media";
import { createDocsNotifier } from "~/utils/feedback";
import { ManageEmpty, SkeletonList } from "~/utils/manageComponents";
import { useMinimumLoading } from "@yueli/ui/feedback";
import { TabbedSurface, AuthorizationGrantBadge } from "@yueli/ui/admin";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "权限与申请 · 控制台" });

interface Role {
  key: string;
  displayName: string;
  kind: string;
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
  grants: Grant[];
  capabilities: Capability[];
  snapshot: PolicySnapshot | null;
}
interface Capability {
  key: string;
  displayName: string;
}
interface Grant {
  validFrom: string;
  id: string;
  subject: string;
  role: string;
  source: string;
}
interface PublicUser {
  avatar?: { mediaKey: string };
  userKey: string;
  handle: string;
  displayName: string;
}
interface AuthorizationUserRow {
  subject: string;
  grants: Grant[];
}
interface AuthorizationConsole {
  activeRevision: number;
  policy: Policy;
  roles: Role[];
  automaticRules: Array<{ key: string; enabled: boolean }>;
  applications: Application[];
  grants: Grant[];
  capabilities: Capability[];
}
type AuthorizationTab = "applications" | "permissions" | "users";
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
const activeTab = ref<AuthorizationTab>("applications");
const expandedRoleKey = ref("");
const publishing = ref(false);
const operationError = ref("");
const lastImpact = ref<Impact | null>(null);
const identityUsers = ref<Record<string, PublicUser>>({});
const userQuery = ref("");
const userRole = ref("all");
const selectedUsers = ref<string[]>([]);
const grantOpen = ref(false);
const grantBusy = ref(false);
const grantForm = reactive({ subject: "", role: "author" });
const revokeOpen = ref(false);
const revokeTargets = ref<Grant[]>([]);
const revokeError = ref("");
onMounted(() => {
  mounted.value = true;
});

const { data, pending, error, refresh } = await useAsyncData(
  "docs-authorization-management",
  async (): Promise<AuthorizationData> => {
    if (!isAdministrator.value) {
      return {
        roles: [],
        policies: [],
        applications: [],
        grants: [],
        capabilities: [],
        snapshot: null,
      };
    }
    const console = await call<AuthorizationConsole>(
      "/api/v1/authorization/manage/console",
    );
    return {
      roles: console.roles,
      policies: [console.policy],
      applications: console.applications,
      grants: console.grants,
      capabilities: console.capabilities,
      snapshot: {
        roles: console.roles,
        automaticRules: console.automaticRules,
      },
    };
  },
  {
    server: false,
    default: () => ({
      roles: [],
      policies: [],
      applications: [],
      grants: [],
      capabilities: [],
      snapshot: null,
    }),
  },
);

const showSkeleton = useMinimumLoading(
  computed(() => !mounted.value || pending.value),
);
const roles = computed(() => data.value?.roles ?? []);
const applications = computed(() => data.value?.applications ?? []);
const grants = computed(() => data.value?.grants ?? []);
const capabilities = computed(() => data.value?.capabilities ?? []);
const activePolicy = computed(() =>
  data.value?.policies.find((policy) => policy.state === "active"),
);
const author = computed(() =>
  roles.value.find((role) => role.key === "author"),
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
const tabItems = computed(() => [
  {
    label: "申请",
    value: "applications",
    icon: "i-tabler-inbox",
    badge: applications.value.length || undefined,
  },
  { label: "权限", value: "permissions", icon: "i-tabler-shield-lock" },
  {
    label: "用户管理",
    value: "users",
    icon: "i-tabler-users",
    badge: userRows.value.length || undefined,
  },
]);
const roleOptions = computed(() => [
  { label: "全部角色", value: "all" },
  ...roles.value.map((role) => ({ label: role.displayName, value: role.key })),
]);
const grantRoleOptions = computed(() =>
  roles.value
    .filter((role) => role.kind !== "custom" || role.capabilities.length)
    .map((role) => ({ label: role.displayName, value: role.key })),
);
const userRows = computed<AuthorizationUserRow[]>(() => {
  const rows = new Map<string, Grant[]>();
  for (const grant of grants.value) {
    rows.set(grant.subject, [...(rows.get(grant.subject) ?? []), grant]);
  }
  return [...rows.entries()]
    .map(([subject, subjectGrants]) => ({ subject, grants: subjectGrants }))
    .sort((left, right) =>
      userName(left.subject).localeCompare(userName(right.subject), "zh-CN"),
    );
});
const filteredUsers = computed(() => {
  const keyword = userQuery.value.trim().toLocaleLowerCase();
  return userRows.value.filter((user) => {
    if (
      userRole.value !== "all" &&
      !user.grants.some((grant) => grant.role === userRole.value)
    ) {
      return false;
    }
    if (!keyword) return true;
    return [
      user.subject,
      userName(user.subject),
      identityUsers.value[user.subject]?.handle,
    ].some((value) => value?.toLocaleLowerCase().includes(keyword));
  });
});
const allUsersSelected = computed(
  () =>
    filteredUsers.value.length > 0 &&
    filteredUsers.value.every((user) =>
      selectedUsers.value.includes(user.subject),
    ),
);
const subjectKey = computed(() =>
  [
    ...new Set(
      [...applications.value, ...grants.value]
        .map((item) => item.subject)
        .filter(Boolean),
    ),
  ]
    .sort()
    .join(","),
);

watch(
  subjectKey,
  async (value) => {
    if (!import.meta.client || !value) return;
    const subjects = value.split(",");
    const users: Record<string, PublicUser> = { ...identityUsers.value };
    try {
      for (let index = 0; index < subjects.length; index += 100) {
        const result = await $fetch<{ items: PublicUser[] }>(
          "/identity-api/api/v1/users",
          { query: { ids: subjects.slice(index, index + 100).join(",") } },
        );
        for (const user of result.items) users[user.userKey] = user;
      }
      identityUsers.value = users;
    } catch {
      toast.add({ title: "用户资料加载失败", description: "刷新页面后重试", color: "error" });
    }
  },
  { immediate: true },
);
watch([userQuery, userRole], () => {
  selectedUsers.value = [];
});

function normalizedCapabilities(values: readonly string[]) {
  return [...new Set(values)].sort();
}

function capabilityLabel(key: string) {
  return (
    capabilities.value.find((item) => item.key === key)?.displayName ??
    capabilityOptions.find((item) => item.key === key)?.label ??
    key
  );
}

function roleCapabilityOptions(role: Role) {
  if (role.protected) return capabilities.value;
  const editable = new Set<string>(capabilityOptions.map((item) => item.key));
  return capabilities.value.filter((item) => editable.has(item.key));
}

function roleHasCapability(role: Role, capability: string) {
  if (role.key === "author") {
    return selectedAuthorCapabilities.value.includes(capability);
  }
  return role.capabilities.includes(capability);
}

function toggleRoleRowCapability(
  role: Role,
  capability: string,
  enabled: boolean,
) {
  if (role.key !== "author" || role.protected) return;
  selectedAuthorCapabilities.value = toggleCapability(
    selectedAuthorCapabilities.value,
    capability,
    enabled,
  );
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
  } catch (caught) {
    reviewError.value =
      caught instanceof Error ? caught.message : "申请处理失败，请重试。";
  } finally {
    reviewing.value = false;
  }
}

function roleLabel(key: string) {
  return roles.value.find((role) => role.key === key)?.displayName ?? key;
}


const accountOrigin = useRuntimeConfig().public.accountUrl;
function userProfile(subject: string) {
  return `${String(accountOrigin).replace(/\/$/, "")}/u/${encodeURIComponent(subject)}`;
}
function userAvatar(subject: string) {
  const avatar = identityUsers.value[subject]?.avatar;
  return avatar ? { src: assetMediaUrl(avatar, "thumbnail"), alt: userName(subject) } : { icon: "i-tabler-user" };
}
function userSince(grants: { validFrom: string }[]) {
  return grants.map((grant) => grant.validFrom).filter((value) => value && !value.startsWith("0001") && Number.isFinite(Date.parse(value))).sort((a,b) => Date.parse(a)-Date.parse(b))[0] || "";
}

function userName(subject: string) {
  const user = identityUsers.value[subject];
  return user?.displayName || user?.handle || subject;
}

function userMeta(subject: string) {
  const user = identityUsers.value[subject];
  if (!user) return subject;
  return user.handle ? `@${user.handle} · ${subject}` : subject;
}

function toggleUser(subject: string, selected: boolean) {
  selectedUsers.value = selected
    ? [...new Set([...selectedUsers.value, subject])]
    : selectedUsers.value.filter((item) => item !== subject);
}

function toggleAllUsers(selected: boolean) {
  const visibleSubjects = filteredUsers.value.map((user) => user.subject);
  selectedUsers.value = selected
    ? [...new Set([...selectedUsers.value, ...visibleSubjects])]
    : selectedUsers.value.filter(
        (subject) => !visibleSubjects.includes(subject),
      );
}

function openGrant() {
  grantOpen.value = true;
}

function closeGrant() {
  grantOpen.value = false;
}

async function grantRole() {
  if (!grantForm.subject.trim() || !grantForm.role || grantBusy.value) return;
  grantBusy.value = true;
  try {
    await call("/api/v1/authorization/manage/grants", {
      method: "POST",
      body: { subject: grantForm.subject.trim(), role: grantForm.role },
    });
    grantForm.subject = "";
    grantOpen.value = false;
    await refresh();
  } catch (caught) {
    toast.add({
      title: "添加用户失败",
      description:
        caught instanceof Error ? caught.message : "请检查用户标识后重试。",
      color: "error",
    });
  } finally {
    grantBusy.value = false;
  }
}

async function revokeGrants(targets: Grant[]) {
  if (!targets.length || grantBusy.value) return;
  grantBusy.value = true;
  try {
    await Promise.all(
      targets.map((grant) =>
        call(`/api/v1/authorization/manage/grants/${grant.id}`, {
          method: "DELETE",
        }),
      ),
    );
    selectedUsers.value = [];
    await refresh();
    revokeOpen.value = false;
    revokeTargets.value = [];
    revokeError.value = "";
  } catch (caught) {
    revokeError.value = authorizationErrorMessage(caught, targets);
    await refresh();
  } finally {
    grantBusy.value = false;
  }
}

function authorizationErrorCode(caught: unknown) {
  const data = (caught as {
    data?: {
      code?: string;
      type?: string;
      failure?: { code?: string };
    };
  } | null)?.data;
  if (data?.failure?.code) return data.failure.code;
  if (data?.code) return data.code;
  return data?.type?.split("/").at(-1) ?? "";
}

function authorizationErrorMessage(caught: unknown, targets: Grant[] = []) {
  const code = authorizationErrorCode(caught);
  if (
    code === "docs.administrator_grant_protected" ||
    ((code === "docs.authorization_unavailable" ||
      code === "blog.authorization_unavailable") &&
      targets.some((grant) => grant.role === "administrator"))
  ) {
    return "不能撤销当前管理员角色。当前账号正在依靠该角色管理本站，或本站只剩这一位管理员；请先添加另一位管理员后再操作。";
  }
  if (code === "docs.forbidden" || code === "blog.forbidden") {
    return "当前账号没有撤销该角色的权限。";
  }
  if (
    code === "docs.authorization_unavailable" ||
    code === "blog.authorization_unavailable"
  ) {
    return "权限服务暂时不可用，请稍后重试。";
  }
  return "角色未撤销，请刷新权限数据后重试。";
}

function openRevoke(targets: Grant[]) {
  if (!targets.length) return;
  revokeTargets.value = [...targets];
  revokeError.value = "";
  revokeOpen.value = true;
}

function revokeGrant(grant: Grant) {
  openRevoke([grant]);
}

function revokeSelectedUsers() {
  const targets = grants.value.filter((grant) =>
    selectedUsers.value.includes(grant.subject),
  );
  openRevoke(targets);
}

function closeRevoke() {
  if (grantBusy.value) return;
  revokeOpen.value = false;
  revokeTargets.value = [];
  revokeError.value = "";
}

function submitRevoke() {
  return revokeGrants(revokeTargets.value);
}

function formatDate(value: string) {
  if (!value) return "未记录";
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return "未记录";
  return new Intl.DateTimeFormat("zh-CN", {timeZone: "Asia/Shanghai",
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
}
</script>

<template>
  <ManagePage
    id="authorization"
    title="权限与申请"
    icon="i-tabler-shield-lock"
    main-id="manage-main"
    body-class="w-full space-y-5"
  >
    <template #actions>
      <UButton
        v-if="isAdministrator && activeTab === 'permissions'"
        label="新建自定义角色"
        icon="i-tabler-user-shield"
        color="neutral"
        variant="outline"
        @click="openCreateRole"
      />
      <UButton
        v-else-if="isAdministrator && activeTab === 'users'"
        label="添加用户"
        icon="i-tabler-user-plus"
        @click="openGrant"
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

      <TabbedSurface
        v-model="activeTab"
        :items="tabItems"
        navigation-label="权限管理"
        data-manage-surface="authorization"
      >
          <section
            v-if="activeTab === 'applications'"
            aria-labelledby="applications-title"
            class="overflow-hidden"
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
                    <UUser :name="userName(application.subject)" :description="userMeta(application.subject)" :avatar="userAvatar(application.subject)" :to="userProfile(application.subject)" target="_blank" rel="noopener noreferrer" title="查看用户主页" />
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
                    申请时间：{{ formatDate(application.createdAt) }}
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

        <div v-else-if="activeTab === 'permissions'" class="min-w-0">
          <section aria-labelledby="roles-title">
            <div class="border-b border-default px-4 py-4 sm:px-5">
              <div>
                <h2
                  id="roles-title"
                  class="text-sm font-semibold text-highlighted"
                >
                  角色与能力
                </h2>
                <p class="mt-1 text-xs text-muted">
                  管理员能力受保护；作者只能管理自己拥有的文档。
                </p>
              </div>
            </div>

            <div class="divide-y divide-default">
              <article
                v-for="role in roles"
                :key="role.key"
                class="min-w-0 px-4 py-4 sm:px-5"
                data-authorization-role-row
              >
                <button
                  type="button"
                  class="flex w-full items-center gap-3 text-left"
                  :aria-expanded="expandedRoleKey === role.key"
                  @click="
                    expandedRoleKey =
                      expandedRoleKey === role.key ? '' : role.key
                  "
                >
                  <span
                    class="grid size-9 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary"
                  >
                    <UIcon name="i-tabler-shield-lock" class="size-5" />
                  </span>
                  <span class="min-w-0 flex-1">
                    <span class="flex flex-wrap items-center gap-2">
                      <span class="font-medium text-highlighted">
                        {{ role.displayName }}
                      </span>
                      <UBadge
                        :label="
                          role.protected
                            ? '受保护'
                            : role.kind === 'custom'
                              ? '自定义'
                              : '内置'
                        "
                        color="neutral"
                        variant="soft"
                      />
                    </span>
                    <span class="mt-1 block text-xs text-muted">
                      {{ role.capabilities.length }} 项能力
                    </span>
                  </span>
                  <UIcon
                    name="i-tabler-chevron-down"
                    class="size-5 shrink-0 text-muted transition-transform"
                    :class="expandedRoleKey === role.key ? 'rotate-180' : ''"
                  />
                </button>

                <div
                  v-if="expandedRoleKey === role.key"
                  class="mt-4 border-t border-default pt-4 sm:ml-12"
                >
                  <div class="grid gap-3 sm:grid-cols-2 xl:grid-cols-3">
                    <UCheckbox
                      v-for="capability in roleCapabilityOptions(role)"
                      :key="capability.key"
                      :model-value="roleHasCapability(role, capability.key)"
                      :label="capability.displayName"
                      :disabled="role.protected || role.key !== 'author' || publishing"
                      @update:model-value="
                        toggleRoleRowCapability(
                          role,
                          capability.key,
                          Boolean($event),
                        )
                      "
                    />
                  </div>

                  <div
                    v-if="role.key === 'author' || role.kind === 'custom'"
                    class="mt-4 flex flex-wrap items-center justify-end gap-2 border-t border-default pt-3"
                  >
                    <template v-if="role.key === 'author'">
                      <span v-if="!authorDirty" class="mr-auto text-xs text-muted">
                        作者能力没有未发布变更
                      </span>
                      <UButton
                        label="复核并发布"
                        icon="i-tabler-shield-check"
                        size="sm"
                        :disabled="!authorDirty"
                        @click="prepareAuthorCapabilities"
                      />
                    </template>
                    <template v-else>
                      <UButton
                        label="编辑角色"
                        icon="i-tabler-pencil"
                        color="neutral"
                        variant="outline"
                        size="sm"
                        @click="openEditRole(role)"
                      />
                      <UButton
                        label="停用角色"
                        icon="i-tabler-user-off"
                        color="error"
                        variant="soft"
                        size="sm"
                        @click="prepareRoleRetirement(role)"
                      />
                    </template>
                  </div>
                </div>
              </article>
            </div>
          </section>

          <section class="border-t border-default px-4 py-5 sm:px-5">
            <div class="flex items-center justify-between gap-4">
              <div>
                <h2 class="text-sm font-semibold text-highlighted">
                  注册用户自动成为作者
                </h2>
                <p class="mt-1 text-xs text-muted">
                  开启后，新用户首次进入本站时获得作者角色。
                </p>
              </div>
              <USwitch
                :model-value="autoAuthorEnabled"
                aria-label="注册用户自动成为作者"
                :disabled="publishing"
                @update:model-value="prepareAutomaticAuthor(Boolean($event))"
              />
            </div>
          </section>
        </div>

        <div v-else class="min-w-0">
          <div
            class="grid gap-3 border-b border-default p-4 sm:grid-cols-[minmax(0,1fr)_14rem] sm:px-5"
          >
            <UInput
              v-model="userQuery"
              icon="i-tabler-search"
              placeholder="搜索用户"
              class="w-full"
            />
            <USelect
              v-model="userRole"
              :items="roleOptions"
              value-key="value"
              class="w-full"
              aria-label="筛选用户角色"
            />
          </div>

          <div
            v-if="selectedUsers.length"
            class="flex flex-wrap items-center justify-between gap-3 border-b border-default bg-primary/5 px-4 py-3 sm:px-5"
            data-authorization-user-bulk
          >
            <span class="text-sm font-medium text-highlighted">
              已选择 {{ selectedUsers.length }} 位用户
            </span>
            <UButton
              label="批量撤销角色"
              color="error"
              variant="soft"
              size="sm"
              :disabled="grantBusy"
              @click="revokeSelectedUsers"
            />
          </div>

          <div v-if="filteredUsers.length" class="min-w-0">
            <div
              class="flex items-center border-b border-default px-4 py-3 sm:px-5"
            >
              <UCheckbox
                :model-value="allUsersSelected"
                label="选择全部结果"
                @update:model-value="toggleAllUsers(Boolean($event))"
              />
            </div>
            <div class="divide-y divide-default">
              <article
                v-for="user in filteredUsers"
                :key="user.subject"
                class="grid min-w-0 gap-3 px-4 py-4 sm:grid-cols-[auto_minmax(12rem,0.8fr)_minmax(16rem,1.2fr)] sm:items-center sm:px-5"
                data-authorization-user-row
              >
                <UCheckbox
                  :model-value="selectedUsers.includes(user.subject)"
                  :aria-label="`选择 ${userName(user.subject)}`"
                  @update:model-value="toggleUser(user.subject, Boolean($event))"
                />
                <div class="min-w-0">
                  <UUser :name="userName(user.subject)" :description="userMeta(user.subject)" :avatar="userAvatar(user.subject)" :to="userProfile(user.subject)" target="_blank" rel="noopener noreferrer" title="查看用户主页" />
                  <time class="mt-2 block text-xs text-muted" :datetime="userSince(user.grants) || undefined">授权生效：{{ formatDate(userSince(user.grants)) }}</time>
                </div>
                <div
                  class="ml-7 flex min-w-0 flex-wrap gap-2 sm:ml-0 sm:justify-end"
                >
                  <div
                    v-for="grant in user.grants"
                    :key="grant.id"
                    class="inline-flex min-w-0 items-center gap-2"
                  >
                    <AuthorizationGrantBadge :role="roleLabel(grant.role)" :source="grant.source" />
                    <UButton
                      label="撤销"
                      icon="i-tabler-user-minus"
                      color="neutral"
                      variant="outline"
                      size="sm"
                      :aria-label="`撤销 ${userName(user.subject)} 的${roleLabel(grant.role)}角色`"
                      :disabled="grantBusy"
                      @click="revokeGrant(grant)"
                    />
                  </div>
                </div>
              </article>
            </div>
          </div>
          <ManageEmpty
            v-else
            icon="i-tabler-users"
            text="当前没有匹配的用户"
            class="m-5"
          />
        </div>
      </TabbedSurface>
    </template>

    <UModal
      v-model:open="grantOpen"
      title="添加用户"
      description="为 Identity 用户授予本站角色。"
      :ui="{ footer: 'justify-end' }"
    >
      <template #body>
        <div class="space-y-4">
          <UFormField label="用户标识" required>
            <UInput
              v-model="grantForm.subject"
              placeholder="Identity 用户 ID"
              class="w-full"
            />
          </UFormField>
          <UFormField label="角色" required>
            <USelect
              v-model="grantForm.role"
              value-key="value"
              :items="grantRoleOptions"
              class="w-full"
            />
          </UFormField>
        </div>
      </template>
      <template #footer>
        <UButton
          label="取消"
          color="neutral"
          variant="outline"
          :disabled="grantBusy"
          @click="closeGrant"
        />
        <UButton
          label="添加用户"
          icon="i-tabler-user-plus"
          :disabled="!grantForm.subject.trim() || grantBusy"
          :loading="grantBusy"
          @click="grantRole"
        />
      </template>
    </UModal>

    <UModal
      v-model:open="revokeOpen"
      title="撤销角色"
      :description="
        revokeTargets.length === 1
          ? `撤销 ${userName(revokeTargets[0]?.subject || '')} 的${roleLabel(revokeTargets[0]?.role || '')}角色。`
          : `撤销所选 ${selectedUsers.length} 位用户的全部本站角色。`
      "
      :dismissible="!grantBusy"
      :ui="{ footer: 'justify-end' }"
    >
      <template #body>
        <div class="space-y-4">
          <UAlert
            v-if="revokeError"
            color="error"
            variant="soft"
            icon="i-tabler-alert-circle"
            title="角色未撤销"
            :description="revokeError"
          />
          <p class="text-sm leading-6 text-muted">
            撤销后，对应用户会立即失去该角色授予的本站权限。管理员角色必须至少保留一位，当前账号也不能撤销正在提供本次管理权限的角色。
          </p>
        </div>
      </template>
      <template #footer>
        <UButton
          label="取消"
          color="neutral"
          variant="outline"
          :disabled="grantBusy"
          @click="closeRevoke"
        />
        <UButton
          label="确认撤销"
          icon="i-tabler-user-minus"
          color="error"
          :loading="grantBusy"
          @click="submitRevoke"
        />
      </template>
    </UModal>

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
  </ManagePage>
</template>
