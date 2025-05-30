<template>
  <BBGrid
    class="border"
    :column-list="COLUMNS"
    :data-source="filteredTemplateList"
    :row-clickable="false"
    row-key="key"
  >
    <template #item="{ item: tpl }: { item: RuleTemplate }">
      <div class="bb-grid-cell !pl-2 !pr-1 break-normal">
        {{ titleOfTemplate(tpl) }}
      </div>
      <div class="bb-grid-cell gap-x-1 !pl-1 !pr-2">
        <NButton size="tiny" @click="applyTemplate(tpl)">
          {{ $t("custom-approval.risk-rule.template.load") }}
        </NButton>
        <NButton size="tiny" @click="viewTemplate = tpl">
          {{ $t("common.view") }}
        </NButton>
      </div>
    </template>
  </BBGrid>

  <BBModal
    v-if="viewTemplate"
    :title="$t('custom-approval.risk-rule.template.view')"
    class="!w-auto lg:!max-w-[36rem]"
    @close="viewTemplate = undefined"
  >
    <template #header>
      <div>{{ $t("custom-approval.risk-rule.template.view") }}</div>
    </template>
    <template #subtitle>
      <div
        class="text-xs text-control-light mt-1 whitespace-pre-wrap overflow-hidden"
      >
        {{ titleOfTemplate(viewTemplate) }}
      </div>
    </template>
    <ViewTemplate :template="viewTemplate" />
  </BBModal>
</template>

<script lang="ts" setup>
import { NButton, useDialog } from "naive-ui";
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";
import { BBGrid, BBModal, type BBGridColumn } from "@/bbkit";
import type { ConditionGroupExpr } from "@/plugins/cel";
import { buildCELExpr } from "@/plugins/cel";
import { Expr } from "@/types/proto/google/type/expr";
import type { Risk } from "@/types/proto/v1/risk_service";
import { Risk_Source } from "@/types/proto/v1/risk_service";
import { batchConvertParsedExprToCELString, defer } from "@/utils";
import { useRiskCenterContext } from "../context";
import ViewTemplate from "./ViewTemplate.vue";
import {
  type RuleTemplate,
  useRuleTemplates,
  titleOfTemplate,
} from "./template";

const props = defineProps<{
  dirty?: boolean;
  source: Risk_Source;
}>();

const emit = defineEmits<{
  (
    event: "apply-template",
    overrides: Partial<Risk>,
    expr: ConditionGroupExpr
  ): void;
}>();

const { t } = useI18n();
const { dialog } = useRiskCenterContext();
const templateList = useRuleTemplates();
const viewTemplate = ref<RuleTemplate>();
const nDialog = useDialog();

const COLUMNS = computed(() => {
  const columns: BBGridColumn[] = [
    {
      title: t("custom-approval.risk-rule.rule-name"),
      width: "1fr",
      class: "!pl-2 !pr-1 !py-0.5",
    },
    {
      title: t("common.operations"),
      width: "auto",
      class: "!px-1",
    },
  ];
  return columns;
});

const filteredTemplateList = computed(() => {
  return templateList.value.filter((tpl) => {
    return (
      tpl.source === Risk_Source.SOURCE_UNSPECIFIED ||
      tpl.source === props.source
    );
  });
});

const confirmApplyTemplate = async () => {
  if (dialog.value?.mode === "CREATE" && !props.dirty) {
    return true;
  }
  const d = defer<boolean>();
  nDialog.warning({
    title: t("custom-approval.risk-rule.template.load-template"),
    content: t("common.will-override-current-data"),
    maskClosable: false,
    closeOnEsc: false,
    positiveText: t("common.confirm"),
    negativeText: t("common.cancel"),
    onPositiveClick: () => d.resolve(true),
    onNegativeClick: () => d.resolve(false),
  });
  return d.promise;
};

const applyTemplate = async (template: RuleTemplate) => {
  if (!(await confirmApplyTemplate())) {
    return;
  }

  const { expr, source, level } = template;
  const title = titleOfTemplate(template);
  const { mode } = dialog.value!;
  const celexpr = await buildCELExpr(expr);
  if (!celexpr) {
    return;
  }
  const expressions = await batchConvertParsedExprToCELString([celexpr]);
  const overrides: Partial<Risk> = {
    condition: Expr.fromPartial({ expression: expressions[0] }),
    level,
    title,
  };
  if (mode === "CREATE" && source !== Risk_Source.SOURCE_UNSPECIFIED) {
    overrides.source = source;
  }
  emit("apply-template", overrides, expr);
};
</script>
