<template>
  <NScrollbar ref="scrollbarRef" class="space-y-5 m-2">
    <div
      v-for="(row, rowIndex) of rows"
      :key="`rows-${rowIndex + offset}`"
      class="font-mono"
    >
      <p
        class="font-medium text-gray-500 dark:text-gray-300 overflow-hidden whitespace-nowrap"
      >
        ******************************** {{ rowIndex + offset + 1 }}. row
        ********************************
      </p>
      <div class="py-2 px-3 bg-gray-50 dark:bg-gray-700 rounded-sm">
        <div
          v-for="header in table.getFlatHeaders()"
          :key="header.index"
          class="flex items-center text-gray-500 dark:text-gray-300 text-sm"
        >
          <div class="min-w-[7rem] text-left flex items-center font-medium">
            {{ header.column.columnDef.header }}
            <MaskingReasonPopover
              v-if="getMaskingReason && getMaskingReason(header.index)"
              :reason="getMaskingReason(header.index)"
              class="ml-0.5 shrink-0"
            />
            <SensitiveDataIcon
              v-else-if="isSensitiveColumn(header.index)"
              class="ml-0.5 shrink-0"
            />
            :
          </div>
          <div class="flex-1">
            <TableCell
              :table="table"
              :value="
                row.getVisibleCells()[header.index].getValue() as RowValue
              "
              :keyword="keyword"
              :set-index="setIndex"
              :row-index="offset + rowIndex"
              :col-index="header.index"
              :column-type="getColumnType(header)"
            />
          </div>
        </div>
      </div>
    </div>
  </NScrollbar>
  <div v-if="rows.length === 0" class="text-center w-full my-12 textinfolabel">
    {{ $t("sql-editor.no-data-available") }}
  </div>
</template>

<script setup lang="ts">
import type { Table } from "@tanstack/vue-table";
import { NScrollbar } from "naive-ui";
import { computed, watch, ref } from "vue";
import type { QueryRow, RowValue } from "@/types/proto-es/v1/sql_service_pb";
import TableCell from "./DataTable/TableCell.vue";
import MaskingReasonPopover from "./DataTable/common/MaskingReasonPopover.vue";
import SensitiveDataIcon from "./DataTable/common/SensitiveDataIcon.vue";
import { getColumnType } from "./DataTable/common/utils";
import { useSQLResultViewContext } from "./context";

const props = defineProps<{
  table: Table<QueryRow>;
  setIndex: number;
  offset: number;
  isSensitiveColumn: (index: number) => boolean;
  getMaskingReason?: (index: number) => any;
}>();

const { keyword } = useSQLResultViewContext();
const scrollbarRef = ref<InstanceType<typeof NScrollbar>>();

const rows = computed(() => props.table.getRowModel().rows);

watch(
  () => props.offset,
  () => scrollbarRef.value?.scrollTo({ top: 0 })
);
</script>
