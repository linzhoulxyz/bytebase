<template>
  <NDataTable
    key="sensitive-column-table"
    :columns="dataTableColumns"
    :data="columnList"
    :row-key="itemKey"
    :checked-row-keys="checkedItemKeys"
    size="small"
    @update:checked-row-keys="handleUpdateCheckedRowKeys($event as string[])"
  />
</template>

<script lang="tsx" setup>
import { TrashIcon } from "lucide-vue-next";
import { NDataTable, NPopconfirm, type DataTableColumn } from "naive-ui";
import { computed, h, ref, watch, withModifiers } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter, RouterLink } from "vue-router";
import ClassificationCell from "@/components/ColumnDataTable/ClassificationCell.vue";
import SemanticTypeCell from "@/components/ColumnDataTable/SemanticTypeCell.vue";
import { updateColumnCatalog } from "@/components/ColumnDataTable/utils";
import type { MaskData } from "@/components/SensitiveData/types";
import { MiniActionButton } from "@/components/v2";
import {
  useSettingV1Store,
  useDatabaseCatalog,
  getColumnCatalog,
} from "@/store";
import type { ComposedDatabase } from "@/types";
import { DataClassificationSetting_DataClassificationConfig as DataClassificationConfig } from "@/types/proto/v1/setting_service";
import { autoDatabaseRoute } from "@/utils";

const props = defineProps<{
  database: ComposedDatabase;
  showOperation: boolean;
  rowClickable: boolean;
  rowSelectable: boolean;
  columnList: MaskData[];
  checkedColumnIndexList: number[];
}>();

const emit = defineEmits<{
  (event: "delete", item: MaskData): void;
  (event: "checked:update", list: number[]): void;
}>();

const { t } = useI18n();
const router = useRouter();
const checkedColumnIndex = ref<Set<number>>(
  new Set(props.checkedColumnIndexList)
);
const settingStore = useSettingV1Store();

watch(
  () => props.columnList,
  () => (checkedColumnIndex.value = new Set()),
  { deep: true }
);
watch(
  () => props.checkedColumnIndexList,
  (val) => (checkedColumnIndex.value = new Set(val)),
  { deep: true }
);

const itemKey = (item: MaskData) => {
  const parts = [];
  const { schema, table, column } = item;
  if (schema) {
    parts.push(schema);
  }
  parts.push(table, column);
  return parts.join("::");
};

const databaseCatalog = useDatabaseCatalog(props.database.name, false);

const classificationConfig = computed(() => {
  return (
    settingStore.getProjectClassification(
      props.database.projectEntity.dataClassificationConfigId
    ) ?? DataClassificationConfig.fromPartial({})
  );
});

const checkedItemKeys = computed(() => {
  const keys: string[] = [];
  props.checkedColumnIndexList.forEach((index) => {
    const item = props.columnList[index];
    if (item) {
      keys.push(itemKey(item));
    }
  });
  return keys;
});

const dataTableColumns = computed(() => {
  const columns: DataTableColumn<MaskData>[] = [
    {
      key: "table",
      title: t("common.table"),
      resizable: true,
      width: "minmax(min-content, auto)",
      render(item) {
        return (
          <div>
            <RouterLink
              to={{
                ...autoDatabaseRoute(router, props.database),
                query: {
                  schema: item.schema,
                  table: item.table,
                },
                hash: "overview",
              }}
              class="normal-link"
              exactActiveClass=""
            >
              {item.schema ? `${item.schema}.${item.table}` : item.table}
            </RouterLink>
          </div>
        );
      },
    },
    {
      key: "column",
      title: t("database.column"),
      resizable: true,
      width: "minmax(min-content, auto)",
      render(item) {
        return item.column;
      },
    },
    {
      key: "semantic-type",
      title: t("settings.sensitive-data.semantic-types.table.semantic-type"),
      width: "minmax(min-content, auto)",
      resizable: true,
      render(item) {
        return (
          <SemanticTypeCell
            database={props.database}
            schema={item.schema}
            table={item.table}
            column={item.column}
            readonly={!props.showOperation}
          />
        );
      },
    },
    {
      key: "classification",
      title: t("database.classification.self"),
      width: "minmax(min-content, auto)",
      resizable: true,
      render(item) {
        const columnCatalog = getColumnCatalog(
          databaseCatalog.value,
          item.schema,
          item.table,
          item.column
        );

        return (
          <ClassificationCell
            classification={columnCatalog.classification}
            classificationConfig={classificationConfig.value}
            readonly={!props.showOperation}
            onApply={(id: string) => onClassificationIdApply(item, id)}
          />
        );
      },
    },
  ];
  if (props.showOperation) {
    columns.push({
      key: "operation",
      title: t("common.operation"),
      width: "6rem",
      render(item) {
        return h(
          NPopconfirm,
          {
            onPositiveClick: () => {
              emit("delete", item);
            },
          },
          {
            trigger: () => {
              return h(
                MiniActionButton,
                {
                  tag: "div",
                  onClick: withModifiers(() => {
                    // noop
                  }, ["stop", "prevent"]),
                },
                {
                  default: () => h(TrashIcon, { class: "w-4 h-4" }),
                }
              );
            },
            default: () =>
              h(
                "div",
                { class: "whitespace-nowrap" },
                t("settings.sensitive-data.remove-sensitive-column-tips")
              ),
          }
        );
      },
    });
  }
  if (props.rowSelectable) {
    columns.unshift({
      type: "selection",
      multiple: true,
      cellProps: () => {
        return {
          onClick: (e: MouseEvent) => {
            e.stopPropagation();
          },
        };
      },
    });
  }
  return columns;
});

const onClassificationIdApply = async (
  item: MaskData,
  classification: string
) => {
  await updateColumnCatalog({
    database: props.database.name,
    schema: item.schema,
    table: item.table,
    column: item.column,
    columnCatalog: { classification },
    notification: !classification ? "common.removed" : undefined,
  });
};

const handleUpdateCheckedRowKeys = (keys: string[]) => {
  const keysSet = new Set(keys);
  const checkedIndexList: number[] = [];
  props.columnList.forEach((item, index) => {
    const key = itemKey(item);
    if (keysSet.has(key)) {
      checkedIndexList.push(index);
    }
  });
  emit("checked:update", checkedIndexList);
};
</script>
