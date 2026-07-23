<script setup lang="ts">
import { createPlatformNotifier } from "@platform/ui/feedback";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "权限与申请 · 控制台" });

interface Role {
  key: string;
  displayName: string;
  protected: boolean;
  capabilities: string[];
}
interface Policy {
  number: number;
  state: string;
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
  automaticRules: Array<{ key: string; enabled: boolean }>;
}

const { call } = useApi();
const toast = createPlatformNotifier(useToast());
const busy = ref(false);
const editableCapabilities = [
  ["docs.document.create", "新建文档"],
  ["docs.document.read", "查看管理文档"],
  ["docs.document.update", "编辑自己的文档"],
  ["docs.document.publish", "发布自己的文档"],
  ["docs.document.archive", "归档自己的文档"],
] as const;

const { data: rolesData, refresh: refreshRoles } = await useAsyncData(
  "docs-authorization-roles",
  () => call<{ items: Role[] }>("/api/v1/authorization/manage/roles"),
  { server: false, default: () => ({ items: [] }) },
);
const { data: policiesData, refresh: refreshPolicies } = await useAsyncData(
  "docs-authorization-policies",
  () => call<{ items: Policy[] }>("/api/v1/authorization/manage/policies"),
  { server: false, default: () => ({ items: [] }) },
);
const { data: applicationsData, refresh: refreshApplications } =
  await useAsyncData(
    "docs-authorization-applications",
    () =>
      call<{ items: Application[] }>(
        "/api/v1/authorization/manage/applications?state=pending",
      ),
    { server: false, default: () => ({ items: [] }) },
  );

const author = computed(() =>
  rolesData.value?.items.find((role) => role.key === "author"),
);
const selected = ref<string[]>([]);
const autoAuthorEnabled = ref(false);
watch(
  author,
  (role) => {
    selected.value = [...(role?.capabilities ?? [])];
  },
  { immediate: true },
);
watch(
  policiesData,
  async () => {
    const active = policiesData.value?.items.find(
      (policy) => policy.state === "active",
    );
    if (!active) return;
    const snapshot = await call<PolicySnapshot>(
      `/api/v1/authorization/manage/policies/${active.number}`,
    );
    autoAuthorEnabled.value =
      snapshot.automaticRules.find(
        (rule) => rule.key === "docs.registration_author",
      )?.enabled ?? false;
  },
  { immediate: true },
);

function toggleCapability(capability: string, enabled: boolean) {
  selected.value = enabled
    ? [...new Set([...selected.value, capability])]
    : selected.value.filter((item) => item !== capability);
}

async function saveAuthorCapabilities() {
  const active = policiesData.value?.items.find(
    (policy) => policy.state === "active",
  );
  if (!active) return;
  busy.value = true;
  try {
    const draft = await call<{ policy: Policy }>(
      "/api/v1/authorization/manage/policies/drafts",
      { method: "POST", body: { expectedActiveRevision: active.number } },
    );
    await call(
      `/api/v1/authorization/manage/policies/${draft.policy.number}/roles/author/capabilities`,
      { method: "PUT", body: { capabilities: selected.value } },
    );
    const validation = await call<{ valid: boolean; violations: string[] }>(
      `/api/v1/authorization/manage/policies/${draft.policy.number}/validate`,
      { method: "POST" },
    );
    if (!validation.valid)
      throw new Error(validation.violations.join("；") || "策略校验失败");
    await call(
      `/api/v1/authorization/manage/policies/${draft.policy.number}/activate`,
      { method: "POST", body: { expectedActiveRevision: active.number } },
    );
    await Promise.all([refreshRoles(), refreshPolicies()]);
    // feedback-contract: policy activation changes authorization across later requests
    toast.add({ title: "作者能力已更新", color: "success" });
  } catch (error) {
    toast.add({
      title: "保存失败",
      description: error instanceof Error ? error.message : "请稍后重试",
      color: "error",
    });
  } finally {
    busy.value = false;
  }
}

async function review(
  application: Application,
  decision: "approve" | "reject",
) {
  busy.value = true;
  try {
    await call(
      `/api/v1/authorization/manage/applications/${application.id}/review`,
      { method: "POST", body: { decision } },
    );
    await refreshApplications();
    // feedback-contract: review changes another user's authorization outside this list
    toast.add({
      title: decision === "approve" ? "已批准" : "已拒绝",
      color: "success",
    });
  } finally {
    busy.value = false;
  }
}

async function setAutomaticAuthor(enabled: boolean) {
  const active = policiesData.value?.items.find(
    (policy) => policy.state === "active",
  );
  if (!active) return;
  busy.value = true;
  try {
    const draft = await call<{ policy: Policy }>(
      "/api/v1/authorization/manage/policies/drafts",
      { method: "POST", body: { expectedActiveRevision: active.number } },
    );
    await call(
      `/api/v1/authorization/manage/policies/${draft.policy.number}/automatic/docs.registration_author`,
      { method: "PUT", body: { enabled } },
    );
    await call(
      `/api/v1/authorization/manage/policies/${draft.policy.number}/activate`,
      { method: "POST", body: { expectedActiveRevision: active.number } },
    );
    autoAuthorEnabled.value = enabled;
    await refreshPolicies();
    // feedback-contract: automatic grants apply to future registration events
    toast.add({
      title: enabled ? "已开启注册自动作者" : "已关闭注册自动作者",
      color: "success",
    });
  } catch (error) {
    toast.add({
      title: "自动授权设置失败",
      description: error instanceof Error ? error.message : "请稍后重试",
      color: "error",
    });
  } finally {
    busy.value = false;
  }
}
</script>

<template>
  <main class="mx-auto grid w-full max-w-5xl gap-6 p-6">
    <header>
      <h1 class="text-2xl font-semibold text-highlighted">权限与申请</h1>
      <p class="mt-1 text-sm text-muted">
        角色能力通过版本化策略发布；Identity 角色不会自动获得本站权限。
      </p>
    </header>

    <UCard>
      <template #header>
        <div>
          <h2 class="font-semibold">作者能力</h2>
          <p class="text-sm text-muted">
            所有者约束不可关闭，作者始终只能操作自己的文档。
          </p>
        </div>
      </template>
      <div class="grid gap-3 sm:grid-cols-2">
        <UCheckbox
          v-for="[capability, label] in editableCapabilities"
          :key="capability"
          :model-value="selected.includes(capability)"
          :label="label"
          @update:model-value="toggleCapability(capability, Boolean($event))"
        />
      </div>
      <template #footer>
        <UButton
          label="发布新策略"
          :loading="busy"
          @click="saveAuthorCapabilities"
        />
      </template>
    </UCard>

    <UCard>
      <template #header>
        <div>
          <h2 class="font-semibold">注册自动授权</h2>
          <p class="text-sm text-muted">
            仅影响以后收到的注册事件；关闭不会撤销已有作者授权。
          </p>
        </div>
      </template>
      <USwitch
        :model-value="autoAuthorEnabled"
        label="新注册用户自动成为作者"
        :disabled="busy"
        @update:model-value="setAutomaticAuthor(Boolean($event))"
      />
    </UCard>

    <UCard>
      <template #header>
        <h2 class="font-semibold">待审批申请</h2>
      </template>
      <div
        v-if="applicationsData?.items.length"
        class="divide-y divide-default"
      >
        <div
          v-for="application in applicationsData.items"
          :key="application.id"
          class="flex flex-wrap items-center justify-between gap-4 py-4 first:pt-0 last:pb-0"
        >
          <div>
            <p class="font-medium">{{ application.subject }}</p>
            <p class="text-sm text-muted">
              {{ application.reason || "未填写申请理由" }}
            </p>
          </div>
          <div class="flex gap-2">
            <UButton
              size="sm"
              label="批准"
              :loading="busy"
              @click="review(application, 'approve')"
            />
            <UButton
              size="sm"
              label="拒绝"
              color="neutral"
              variant="soft"
              :loading="busy"
              @click="review(application, 'reject')"
            />
          </div>
        </div>
      </div>
      <p v-else class="text-sm text-muted">暂无待审批申请。</p>
    </UCard>
  </main>
</template>
