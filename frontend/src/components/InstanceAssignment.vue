<template>
  <Drawer :show="show" @close="$emit('dismiss')">
    <DrawerContent
      :title="$t('subscription.instance-assignment.manage-license')"
      class="max-w-[100vw]"
    >
      <div class="divide-block-border space-y-5 w-[40rem] h-full">
        <div>
          <div class="flex space-x-2">
            <div class="text-gray-400">
              {{
                $t("subscription.instance-assignment.used-and-total-license")
              }}
            </div>
            <LearnMoreLink
              url="https://docs.bytebase.com/administration/license?source=console"
              class="ml-1 text-sm"
            />
          </div>
          <div class="mt-1 text-4xl flex items-center gap-x-2">
            <span>
              {{ actuatorStore.activatedInstanceCount }}
            </span>
            <span class="text-xl">/</span>
            <span>{{ totalLicenseCount }}</span>
          </div>
        </div>

        <AdvancedSearch
          v-model:params="state.params"
          :autofocus="false"
          :placeholder="$t('instance.filter-instance-name')"
          :scope-options="scopeOptions"
        />
        <PagedInstanceTable
          ref="pagedInstanceTableRef"
          session-key="bb.instance-table"
          :bordered="true"
          :filter="filter"
          :show-address="false"
          :show-external-link="false"
          :show-selection="canManageSubscription"
          :selected-instance-names="Array.from(state.selectedInstance)"
          :disabled-selected-instance-names="disabledSelectedInstanceNames"
          @update:selected-instance-names="
            (list) => (state.selectedInstance = new Set(list))
          "
        />
      </div>

      <template #footer>
        <div class="w-full flex justify-between items-center">
          <div class="w-full flex justify-end items-center gap-x-3">
            <NButton @click.prevent="cancel">
              {{ $t("common.cancel") }}
            </NButton>
            <NButton
              :disabled="
                !canManageSubscription ||
                !assignmentChanged ||
                state.processing ||
                state.selectedInstance.size > instanceLicenseCount
              "
              type="primary"
              @click.prevent="updateAssignment"
            >
              {{ $t("common.confirm") }}
            </NButton>
          </div>
        </div>
      </template>
    </DrawerContent>
  </Drawer>
</template>

<script lang="ts" setup>
import { NButton } from "naive-ui";
import { storeToRefs } from "pinia";
import { reactive, computed, ref, watch } from "vue";
import { useI18n } from "vue-i18n";
import AdvancedSearch from "@/components/AdvancedSearch";
import { useCommonSearchScopeOptions } from "@/components/AdvancedSearch/useCommonSearchScopeOptions";
import { PagedInstanceTable } from "@/components/v2";
import { Drawer, DrawerContent } from "@/components/v2";
import {
  pushNotification,
  useInstanceV1Store,
  useSubscriptionV1Store,
  useDatabaseV1Store,
  useActuatorV1Store,
} from "@/store";
import { environmentNamePrefix } from "@/store/modules/v1/common";
import type { Instance } from "@/types/proto-es/v1/instance_service_pb";
import { PlanType } from "@/types/proto-es/v1/subscription_service_pb";
import { type SearchParams, hasWorkspacePermissionV2 } from "@/utils";
import { convertScopeValueToEngine } from "@/utils/v1/common-conversions";
import LearnMoreLink from "./LearnMoreLink.vue";

const props = withDefaults(
  defineProps<{
    show: boolean;
    selectedInstanceList?: string[];
  }>(),
  {
    show: false,
    selectedInstanceList: () => [],
  }
);

interface LocalState {
  params: SearchParams;
  selectedInstance: Set<string>;
  processing: boolean;
}

const emit = defineEmits(["dismiss"]);

const state = reactive<LocalState>({
  params: {
    query: "",
    scopes: [],
  },
  selectedInstance: new Set(),
  processing: false,
});

const instanceV1Store = useInstanceV1Store();
const databaseV1Store = useDatabaseV1Store();
const subscriptionStore = useSubscriptionV1Store();
const actuatorStore = useActuatorV1Store();
const { t } = useI18n();
const pagedInstanceTableRef = ref<InstanceType<typeof PagedInstanceTable>>();
const { instanceLicenseCount } = storeToRefs(subscriptionStore);

const scopeOptions = computed(
  () => useCommonSearchScopeOptions(["environment", "engine"]).value
);

const selectedEnvironment = computed(() => {
  const environmentId = state.params.scopes.find(
    (scope) => scope.id === "environment"
  )?.value;
  if (!environmentId) {
    return;
  }
  return `${environmentNamePrefix}${environmentId}`;
});

const selectedEngines = computed(() => {
  return state.params.scopes
    .filter((scope) => scope.id === "engine")
    .map((scope) => convertScopeValueToEngine(scope.value));
});

const filter = computed(() => ({
  environment: selectedEnvironment.value,
  query: state.params.query,
  engines: selectedEngines.value,
}));

const canManageSubscription = computed((): boolean => {
  return (
    hasWorkspacePermissionV2("bb.instances.update") &&
    subscriptionStore.currentPlan !== PlanType.FREE
  );
});

const disabledSelectedInstanceNames = computed(() => {
  return new Set(
    pagedInstanceTableRef.value?.dataList
      ?.filter((ins) => ins.activation)
      ?.map((ins) => ins.name) ?? []
  );
});

const selectActivationInstances = (instances: Instance[]) => {
  for (const instance of instances) {
    if (instance.activation) {
      state.selectedInstance.add(instance.name);
    }
  }
};

watch(
  [() => props.show, () => props.selectedInstanceList],
  ([show, selectedInstanceList]) => {
    if (!show) {
      state.selectedInstance = new Set();
      return;
    }

    for (const instanceName of selectedInstanceList) {
      state.selectedInstance.add(instanceName);
    }
    selectActivationInstances(pagedInstanceTableRef.value?.dataList ?? []);
  },
  { immediate: true }
);

watch(
  () => pagedInstanceTableRef.value?.dataList ?? [],
  (dataList) => {
    selectActivationInstances(dataList);
    for (const instance of dataList) {
      if (instance.activation) {
        state.selectedInstance.add(instance.name);
      }
    }
  },
  { deep: true }
);

const totalLicenseCount = computed((): string => {
  if (instanceLicenseCount.value === Number.MAX_VALUE) {
    return t("subscription.unlimited");
  }
  return `${instanceLicenseCount.value}`;
});

const assignmentChanged = computed(() => {
  for (const instanceName of state.selectedInstance) {
    const instance = instanceV1Store.getInstanceByName(instanceName);
    if (!instance.activation) {
      return true;
    }
  }
  return false;
});

const cancel = () => {
  emit("dismiss");
};

const updateAssignment = async () => {
  if (state.processing) {
    return;
  }
  state.processing = true;

  for (const instanceName of state.selectedInstance) {
    const instance = instanceV1Store.getInstanceByName(instanceName);
    if (instance.activation) {
      continue;
    }
    // activate instance
    instance.activation = true;
    const composedInstance = await instanceV1Store.updateInstance(instance, [
      "activation",
    ]);
    databaseV1Store.updateDatabaseInstance(composedInstance);
  }

  // refresh activatedInstanceCount
  await actuatorStore.fetchServerInfo();

  pushNotification({
    module: "bytebase",
    style: "SUCCESS",
    title: t("subscription.instance-assignment.success-notification"),
  });
  state.processing = false;
  emit("dismiss");
};
</script>
