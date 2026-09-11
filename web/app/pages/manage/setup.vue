<script setup lang="ts">
import { getApiFailure } from "@yueli/http-runtime";

definePageMeta({ layout: false });
useSeoMeta({ title: "初始化站点" });

const { brand: siteTitle } = useSiteRuntime();
const { user } = useAuth();
const { call } = useApi();
const { status, refresh, markClaimed } = useAdministratorClaim();
const { refresh: refreshMe } = useMe();
const phase = ref<"idle" | "claiming" | "complete">("idle");
const conflict = ref(false);
const failure = ref(false);

const identity = computed(() => user.value?.name || user.value?.email || "当前账号");
const initial = computed(() => identity.value.trim().charAt(0).toUpperCase() || "用");
const actionLabel = computed(() => {
  if (phase.value === "claiming") return "初始化中";
  if (phase.value === "complete") return "已完成";
  return "初始化站点并成为管理员";
});

async function claimAdministrator() {
  if (phase.value !== "idle") return;
  phase.value = "claiming";
  conflict.value = false;
  failure.value = false;
  try {
    await call<{ claimed: boolean; created: boolean }>(
      "/api/v1/authorization/setup/claim",
      { method: "POST" },
    );
    phase.value = "complete";
    markClaimed();
    await refreshMe();
    await navigateTo("/manage");
  } catch (error: unknown) {
    const code = getApiFailure(error)?.code;
    const latest = await refresh().catch(() => null);
    if (code === "docs.initial_administrator_already_claimed" || latest?.claimed) {
      conflict.value = true;
    } else {
      failure.value = true;
    }
    phase.value = "idle";
  }
}
</script>

<template>
  <div class="min-h-svh bg-muted text-default">
    <header class="flex min-h-16 items-center justify-between gap-4 border-b border-default bg-default px-4 sm:px-6">
      <NuxtLink to="/" class="truncate font-semibold text-highlighted">{{ siteTitle }}</NuxtLink>
      <ConsumerManageAccountControl home-to="/" show-appearance trigger-mode="inline" />
    </header>
    <main class="mx-auto max-w-lg px-4 py-12 sm:py-20">
      <UCard>
        <div class="space-y-6">
          <div class="space-y-2">
            <h1 class="font-display text-2xl font-semibold text-highlighted">初始化站点</h1>
            <p class="text-sm leading-6 text-muted">本站尚未设置管理员。确认后，当前账号将获得本站全部管理权限；此操作仅可成功一次。</p>
          </div>
          <div class="flex items-center gap-3">
            <UAvatar :text="initial" />
            <div class="min-w-0">
              <p class="truncate text-sm font-semibold">{{ identity }}</p>
              <p v-if="user?.email && user.email !== identity" class="truncate text-xs text-muted">{{ user.email }}</p>
            </div>
          </div>
          <UAlert v-if="conflict" color="warning" icon="i-tabler-user-shield" title="本站已由另一位用户完成初始化" description="当前账号没有获得管理员权限。" />
          <UAlert v-else-if="failure" color="error" icon="i-tabler-alert-circle" title="初始化失败" description="请稍后重试。" />
          <div class="flex flex-wrap justify-end gap-2">
            <UButton v-if="conflict" to="/" color="neutral" variant="outline" label="返回首页" />
            <UButton v-else :label="actionLabel" :loading="phase === 'claiming'" :disabled="phase === 'complete' || status?.claimed" @click="claimAdministrator" />
          </div>
        </div>
      </UCard>
    </main>
  </div>
</template>
