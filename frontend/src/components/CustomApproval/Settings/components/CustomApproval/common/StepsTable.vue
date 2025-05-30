<template>
  <BBGrid
    :column-list="COLUMNS"
    :data-source="steps"
    :row-clickable="false"
    class="border"
  >
    <template #item="{ item: step, row }: { item: ApprovalStep; row: number }">
      <div class="bb-grid-cell justify-center text-center !pl-2">
        {{ row + 1 }}
      </div>
      <div class="bb-grid-cell">
        <RoleSelect
          v-if="editable"
          :suffix="''"
          :value="step.nodes[0]?.role"
          style="width: 80%"
          @update:value="
            (role) => {
              step.nodes[0] = {
                type: ApprovalNode_Type.ANY_IN_GROUP,
                role: role as string,
              };
              $emit('update');
            }
          "
        />
        <template v-else>
          {{ approvalNodeText(step.nodes[0]) }}
        </template>
      </div>
      <div v-if="editable" class="bb-grid-cell gap-x-1">
        <NButton
          :disabled="row === 0 || !allowAdmin"
          size="tiny"
          @click="reorder(step, row, -1)"
        >
          <heroicons:arrow-up />
        </NButton>
        <NButton
          :disabled="row === steps.length - 1 || !allowAdmin"
          size="tiny"
          @click="reorder(step, row, 1)"
        >
          <heroicons:arrow-down />
        </NButton>
        <SpinnerButton
          size="tiny"
          :tooltip="$t('custom-approval.approval-flow.node.delete')"
          :disabled="!allowAdmin"
          :on-confirm="() => removeStep(step, row)"
        >
          <heroicons:trash />
        </SpinnerButton>
      </div>
    </template>
    <template v-if="editable && allowAdmin" #footer>
      <NButton @click="addStep">
        <template #icon><heroicons:plus /></template>
        <span>
          {{ $t("custom-approval.approval-flow.node.add") }}
        </span>
      </NButton>
    </template>
  </BBGrid>
</template>

<script lang="ts" setup>
import { NButton } from "naive-ui";
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import { BBGrid, type BBGridColumn } from "@/bbkit";
import { RoleSelect } from "@/components/v2";
import { PresetRoleType } from "@/types";
import type { ApprovalFlow } from "@/types/proto/v1/issue_service";
import {
  ApprovalNode_Type,
  ApprovalStep,
  ApprovalStep_Type,
} from "@/types/proto/v1/issue_service";
import { approvalNodeText } from "@/utils";
import { SpinnerButton } from "../../common";
import { useCustomApprovalContext } from "../context";

const props = defineProps<{
  flow: ApprovalFlow;
  editable: boolean;
}>();

const emit = defineEmits<{
  (event: "update"): void;
}>();

const { t } = useI18n();

const context = useCustomApprovalContext();
const { allowAdmin } = context;

const COLUMNS = computed(() => {
  const columns: BBGridColumn[] = [
    {
      title: t("custom-approval.approval-flow.node.order"),
      width: "4rem",
      class: "justify-center !pl-2",
    },
    { title: t("custom-approval.approval-flow.node.approver"), width: "1fr" },
  ];
  if (props.editable) {
    columns.push({ title: t("common.operations"), width: "auto" });
  }
  return columns;
});

const steps = computed(() => {
  return props.flow.steps;
});

const reorder = (step: ApprovalStep, index: number, offset: -1 | 1) => {
  const target = index + offset;
  if (target < 0 || target >= steps.value.length) return;
  const tmp = steps.value[index];
  steps.value[index] = steps.value[target];
  steps.value[target] = tmp;

  emit("update");
};
const addStep = () => {
  steps.value.push(
    ApprovalStep.fromPartial({
      type: ApprovalStep_Type.ANY,
      nodes: [
        {
          type: ApprovalNode_Type.ANY_IN_GROUP,
          role: PresetRoleType.WORKSPACE_ADMIN,
        },
      ],
    })
  );
  emit("update");
};

const removeStep = async (step: ApprovalStep, index: number) => {
  steps.value.splice(index, 1);
  emit("update");
};
</script>
