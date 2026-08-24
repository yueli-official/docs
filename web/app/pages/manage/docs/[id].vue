<script setup lang="ts">
import DocEditorWorkbench from "~/components/manage/DocEditorWorkbench.vue";

definePageMeta({ layout: "manage" });
</script>

<template>
  <ManagePage
    id="document-editor"
    title="文档编辑器"
    icon="i-tabler-file-pencil"
    main-id="manage-main"
    body-class="w-full"
  >
    <DocEditorWorkbench />
  </ManagePage>
</template>
