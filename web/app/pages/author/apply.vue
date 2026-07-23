<script setup lang="ts">
definePageMeta({ middleware: "auth" });
useSeoMeta({ title: "申请成为作者" });

interface Application {
  id: string;
  state: string;
  reason: string;
}

const { call } = useApi();
const toast = useToast();
const reason = ref("");
const busy = ref(false);
const { data, refresh } = await useAsyncData(
  "docs-my-role-applications",
  () =>
    call<{ items: Application[] }>(
      "/api/v1/authorization/applications/mine",
    ),
  { server: false, default: () => ({ items: [] }) },
);
const pending = computed(() =>
  data.value?.items.find((application) => application.state === "pending"),
);

async function apply() {
  busy.value = true;
  try {
    await call("/api/v1/authorization/applications", {
      method: "POST",
      body: { role: "author", reason: reason.value },
    });
    await refresh();
    toast.add({ title: "申请已提交", color: "success" });
  } finally {
    busy.value = false;
  }
}

async function withdraw() {
  if (!pending.value) return;
  busy.value = true;
  try {
    await call(
      `/api/v1/authorization/applications/${pending.value.id}/withdraw`,
      { method: "POST" },
    );
    await refresh();
    toast.add({ title: "申请已撤回", color: "success" });
  } finally {
    busy.value = false;
  }
}
</script>

<template>
  <main class="mx-auto w-full max-w-2xl px-4 py-16">
    <UCard>
      <template #header>
        <div>
          <h1 class="text-xl font-semibold">申请成为作者</h1>
          <p class="mt-1 text-sm text-muted">作者可以创建文档，并编辑、发布和归档自己创建的内容。</p>
        </div>
      </template>
      <div v-if="pending" class="space-y-4">
        <UAlert title="申请正在等待管理员审批" color="info" variant="soft" />
        <UButton label="撤回申请" color="neutral" variant="soft" :loading="busy" @click="withdraw" />
      </div>
      <form v-else class="space-y-4" @submit.prevent="apply">
        <UFormField label="申请理由">
          <UTextarea v-model="reason" :rows="5" maxlength="2000" class="w-full" />
        </UFormField>
        <UButton type="submit" label="提交申请" :loading="busy" />
      </form>
    </UCard>
  </main>
</template>
