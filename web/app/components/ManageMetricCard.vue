<script setup lang="ts">
import { AdminMetricCard } from "@yueli/ui/admin";
defineProps<{
  label: string;
  value: string | number;
  detail: string;
  icon: string;
  tone?: "primary" | "success" | "warning" | "neutral";
  to?: string;
}>();
</script>

<template>
  <AdminMetricCard :label="label" :value="value" :detail="detail" :icon="icon" :tone="tone" :to="to" />
</template>
