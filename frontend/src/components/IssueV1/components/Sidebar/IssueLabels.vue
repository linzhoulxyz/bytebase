<template>
  <div class="flex flex-col gap-y-1">
    <div class="flex items-center gap-x-1 textlabel">
      <span>{{ $t("common.labels") }}</span>
    </div>
    <IssueLabelSelector
      :disabled="disabled"
      :selected="value"
      :project="project"
      :size="'medium'"
      @update:selected="onLablesUpdate"
    />
  </div>
</template>

<script setup lang="ts">
import IssueLabelSelector from "@/components/IssueV1/components/IssueLabelSelector.vue";
import type { ComposedProject } from "@/types";

withDefaults(
  defineProps<{
    value: string[];
    project: ComposedProject;
    disabled?: boolean;
  }>(),
  {
    disabled: false,
  }
);

const emit = defineEmits<{
  (e: "update:value", labels: string[]): void;
}>();

const onLablesUpdate = async (labels: string[]) => {
  emit("update:value", labels);
};
</script>
